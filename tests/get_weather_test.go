package tests

import (
	"errors"
	"testing"

	"github.com/pietronirod/lab1/internal/entity"
	"github.com/pietronirod/lab1/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetWeatherByCEP_Success(t *testing.T) {
	mockViaCEP := new(MockViaCEPService)
	mockWeatherAPI := new(MockWeatherAPIService)

	mockViaCEP.On("GetLocationByCEP", "01001000").Return("São Paulo", nil)
	mockWeatherAPI.On("GetTemperatureByLocation", "São Paulo").Return(25.0, nil)

	useCase := usecase.NewWeatherUseCase(mockViaCEP, mockWeatherAPI)

	result, err := useCase.GetWeatherByCEP("01001000")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &entity.Temperature{
		Celsius:    25.0,
		Fahrenheit: 77.0,
		Kelvin:     298.0,
	}, result)
}

func TestGetWeatherByCEP_LocationNotFound(t *testing.T) {
	mockViaCEP := new(MockViaCEPService)
	mockWeatherAPI := new(MockWeatherAPIService)

	mockViaCEP.On("GetLocationByCEP", "99999999").Return("", errors.New("location not found"))

	useCase := usecase.NewWeatherUseCase(mockViaCEP, mockWeatherAPI)

	result, err := useCase.GetWeatherByCEP("99999999")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "location not found", err.Error())
}
