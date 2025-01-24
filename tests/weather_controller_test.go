package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pietronirod/lab1/internal/interface/controller"
	"github.com/pietronirod/lab1/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetWeather_ValidCEP(t *testing.T) {
	mockViaCEP, mockWeatherAPI := CreateMockService()

	mockViaCEP.On("GetLocationByCEP", "01001000").Return("São Paulo", nil)
	mockWeatherAPI.On("GetTemperatureByLocation", "São Paulo").Return(25.0, nil)

	useCase := usecase.NewWeatherUseCase(mockViaCEP, mockWeatherAPI)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/weather/:cep", func(c *gin.Context) {
		controller.GetWeatherWithUseCase(c, useCase)
	})

	req, _ := http.NewRequest("GET", "/weather/01001000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "temp_C")
	assert.Contains(t, w.Body.String(), "temp_F")
	assert.Contains(t, w.Body.String(), "temp_K")
}

func TestGetWeather_InvalidCEP(t *testing.T) {
	mockViaCEP, mockWeatherAPI := CreateMockService()

	useCase := usecase.NewWeatherUseCase(mockViaCEP, mockWeatherAPI)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/weather/:cep", func(c *gin.Context) {
		controller.GetWeatherWithUseCase(c, useCase)
	})

	req, _ := http.NewRequest("GET", "/weather/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.JSONEq(t, `{"message": "invalid zipcode"}`, w.Body.String())
}

func TestGetWeather_CEPNotFound(t *testing.T) {
	mockViaCEP, mockWeatherAPI := CreateMockService()

	mockViaCEP.On("GetLocationByCEP", "99999999").Return("", errors.New("location not found"))

	useCase := usecase.NewWeatherUseCase(mockViaCEP, mockWeatherAPI)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/weather/:cep", func(c *gin.Context) {
		cep := c.Param("cep")
		temp, err := useCase.GetWeatherByCEP(cep)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "cannot find zipcode"})
			return
		}
		c.JSON(http.StatusOK, temp)
	})

	req, _ := http.NewRequest("GET", "/weather/99999999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, `{"message": "cannot find zipcode"}`, w.Body.String())
}

func TestGetWeather_TemperatureServiceError(t *testing.T) {
	mockViaCEP, mockWeatherAPI := CreateMockService()

	mockViaCEP.On("GetLocationByCEP", "01001000").Return("São Paulo", nil)
	mockWeatherAPI.On("GetTemperatureByLocation", "São Paulo").Return(0.0, errors.New("failed to fetch temperature"))

	useCase := usecase.NewWeatherUseCase(mockViaCEP, mockWeatherAPI)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/weather/:cep", func(c *gin.Context) {
		controller.GetWeatherWithUseCase(c, useCase)
	})

	req, _ := http.NewRequest("GET", "/weather/01001000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"message": "temperature service error"}`, w.Body.String())
}
