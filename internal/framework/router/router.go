package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pietronirod/lab1/internal/interface/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/weather/:cep", controller.GetWeather)
	return r
}
