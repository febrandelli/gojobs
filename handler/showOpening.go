package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET showing a opening",
	})
}
