package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
