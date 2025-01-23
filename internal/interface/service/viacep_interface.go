package service

type ViaCEPServiceInterface interface {
	GetLocationByCEP(cep string) (string, error)
}
