package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pietronirod/lab1/internal/interface/service"
	"github.com/stretchr/testify/assert"
)

func TestGetLocationByCEP_Success(t *testing.T) {
	mockResponse := `{"localidade": "São Paulo"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	viaCEPService := &service.ViaCEPService{BaseURL: server.URL}
	location, err := viaCEPService.GetLocationByCEP("01001000")

	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", location)
}

func TestGetLocationByCEP_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	viaCEPService := &service.ViaCEPService{BaseURL: server.URL}

	location, err := viaCEPService.GetLocationByCEP("99999999")

	assert.Error(t, err)
	assert.Equal(t, "", location)
	assert.Equal(t, "location not found", err.Error())
}
