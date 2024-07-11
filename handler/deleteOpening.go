package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE deleting opening",
	})
}
