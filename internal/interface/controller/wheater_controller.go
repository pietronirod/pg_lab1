package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pietronirod/lab1/internal/usecase"
)

func GetWeatherWithUseCase(c *gin.Context, weatherUseCase *usecase.WeatherUseCase) {
	cep := c.Param("cep")

	if len(cep) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	temp, err := weatherUseCase.GetWeatherByCEP(cep)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, temp)
}
