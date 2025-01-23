package usecase

import (
	"errors"

	"github.com/pietronirod/lab1/internal/entity"
	"github.com/pietronirod/lab1/internal/interface/service"
)

type WeatherUseCase struct {
	ViaCEPService     service.ViaCEPServiceInterface
	WeatherAPIService service.WeatherAPIServiceInterface
}

func NewWeatherUseCase(viaCEP service.ViaCEPServiceInterface, weatherAPI service.WeatherAPIServiceInterface) *WeatherUseCase {
	return &WeatherUseCase{
		ViaCEPService:     viaCEP,
		WeatherAPIService: weatherAPI,
	}
}

func (uc *WeatherUseCase) GetWeatherByCEP(cep string) (*entity.Temperature, error) {
	location, err := uc.ViaCEPService.GetLocationByCEP(cep)
	if err != nil {
		return nil, errors.New("location not found")
	}

	celsius, err := uc.WeatherAPIService.GetTemperatureByLocation(location)
	if err != nil {
		return nil, errors.New("failed to fetch temperature")
	}

	return &entity.Temperature{
		Celsius:    celsius,
		Fahrenheit: celsius*1.8 + 32,
		Kelvin:     celsius * 273,
	}, nil
}
