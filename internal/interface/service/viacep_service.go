package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ViaCEPService struct {
	BaseURL string
}

type Location struct {
	City string `json:"localidade"`
}

func NewViaCEPService() *ViaCEPService {
	return &ViaCEPService{
		BaseURL: "https://viacep.com.br/ws",
	}
}

func (s *ViaCEPService) GetLocationByCEP(cep string) (string, error) {
	url := fmt.Sprintf("%s/%s/json/", s.BaseURL, cep)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("location not found")
	}
	defer resp.Body.Close()

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil || location.City == "" {
		return "", errors.New("location not found")
	}
	return location.City, nil
}
