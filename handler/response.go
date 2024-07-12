package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"errorCode": code,
		"message":   msg,
	})

	c.Header("Content-Type", "application/json; charset=utf-8")
}

func sendSuccessResponse(c *gin.Context, op string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %v successfull", op),
		"data":    data,
	})
}
