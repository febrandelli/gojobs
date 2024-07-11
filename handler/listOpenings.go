package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET listing all openings",
	})
}
