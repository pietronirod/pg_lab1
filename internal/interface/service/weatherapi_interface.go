package service

type WeatherAPIServiceInterface interface {
	GetTemperatureByLocation(location string) (float64, error)
}
