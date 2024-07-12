package router

import (
	"github.com/febrandelli/gojobs/config"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func Initialize() {
	logger = config.NewLogger("router")
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		logger.Error(err)
	}
}
