package tests

import "github.com/stretchr/testify/mock"

type MockViaCEPService struct {
	mock.Mock
}

func (m *MockViaCEPService) GetLocationByCEP(cep string) (string, error) {
	args := m.Called(cep)
	return args.String(0), args.Error(1)
}

type MockWeatherAPIService struct {
	mock.Mock
}

func (m *MockWeatherAPIService) GetTemperatureByLocation(location string) (float64, error) {
	args := m.Called(location)
	return args.Get(0).(float64), args.Error(1)
}
