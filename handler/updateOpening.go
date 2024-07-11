package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT updating opening",
	})
}
