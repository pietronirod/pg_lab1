package tests

func CreateMockService() (*MockViaCEPService, *MockWeatherAPIService) {
	mockViaCEP := new(MockViaCEPService)
	mockWeatherAPI := new(MockWeatherAPIService)
	return mockViaCEP, mockWeatherAPI
}
