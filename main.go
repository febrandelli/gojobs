package main

import (
	"github.com/febrandelli/gojobs/config"
	"github.com/febrandelli/gojobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Initialize()
	if err != nil {
		logger.Errorf("Error initializing config, error: %v", err)
		return
	}
	router.Initialize()
}
