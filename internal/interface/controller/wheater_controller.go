package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pietronirod/lab1/internal/interface/service"
	"github.com/pietronirod/lab1/internal/usecase"
)

var viaCEPService = service.NewViaCEPService()
var weatherAPIService = service.NewWeatherAPIService()
var weatherUseCase = usecase.NewWeatherUseCase(viaCEPService, weatherAPIService)

func GetWeather(c *gin.Context) {
	cep := c.Param("cep")
	if len(cep) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	temp, err := weatherUseCase.GetWeatherByCEP(cep)
	if err != nil {
		switch err.Error() {
		case "location not found":
			c.JSON(http.StatusNotFound, gin.H{"message": "cannot find zipcode"})
		case "failed to fetch temperature":
			c.JSON(http.StatusInternalServerError, gin.H{"message": "temperature service error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, temp)
}
