package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pietronirod/lab1/internal/interface/controller"
	"github.com/pietronirod/lab1/internal/interface/service"
	"github.com/pietronirod/lab1/internal/usecase"
)

func SetupRouter() *gin.Engine {
	viaCEPService := service.NewViaCEPService()
	weatherAPIService := service.NewWeatherAPIService()
	weatherUseCase := usecase.NewWeatherUseCase(viaCEPService, weatherAPIService)

	r := gin.Default()
	r.GET("/weather/:cep", func(c *gin.Context) {
		controller.GetWeatherWithUseCase(c, weatherUseCase)
	})
	return r
}
