package main

import (
	"github.com/davi-sant/go-opportunities/config"
	"github.com/davi-sant/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("Main")
	// Initialize configs
	err := config.InitDataBase()

	if err != nil {
		logger.Errorf("Error ao tentar iniciar Data Base: %v", err)
		return
	}

	// Initialize routes
	router.Routes()

}
