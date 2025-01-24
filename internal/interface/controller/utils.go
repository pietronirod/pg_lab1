package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch err.Error() {
	case "location not found":
		c.JSON(http.StatusNotFound, gin.H{"message": "cannot find zipcode"})
	case "failed to fetch temperature":
		c.JSON(http.StatusInternalServerError, gin.H{"message": "temperature service error"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
	}
}
