package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type WeatherAPIService struct {
	BaseURL string
}

type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func NewWeatherAPIService() *WeatherAPIService {
	return &WeatherAPIService{
		BaseURL: "http://api.weatherapi.com/v1",
	}
}

func (s *WeatherAPIService) GetTemperatureByLocation(location string) (float64, error) {
	apiKey := viper.GetString("WEATHER_API_KEY")
	if apiKey == "" {
		return 0, errors.New("missing WeatherAPI key")
	}
	encodedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", s.BaseURL, apiKey, encodedLocation)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return 0, errors.New("failed to fetch temperature data")
	}
	defer resp.Body.Close()

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return 0, errors.New("failed to decode temperature data")
	}

	return weatherResp.Current.TempC, nil
}
