package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pietronirod/lab1/internal/interface/service"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureByLocation_Success(t *testing.T) {
	mockReponse := `{"current": {"temp_c": 25.0}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockReponse))
	}))
	defer server.Close()

	viper.Set("WEATHER_API_KEY", "mock_key")
	weatherAPIService := &service.WeatherAPIService{BaseURL: server.URL}
	temp, err := weatherAPIService.GetTemperatureByLocation("São Paulo")

	assert.NoError(t, err)
	assert.Equal(t, 25.0, temp)
}

func TestGetTemperatureByLocation_Failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	viper.Set("WEATHER_API_KEY", "mock_key")
	weatherAPIService := &service.WeatherAPIService{BaseURL: server.URL}
	temp, err := weatherAPIService.GetTemperatureByLocation("São Paulo")

	assert.Error(t, err)
	assert.Equal(t, 0.0, temp)
	assert.Equal(t, "failed to fetch temperature data", err.Error())
}
