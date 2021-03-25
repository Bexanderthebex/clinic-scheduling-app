package main

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	gin "github.com/gin-gonic/gin"
	"log"
)

func main() {
	errorFindingConfig := config.InitiateConfig()
	if errorFindingConfig != nil {
		log.Fatal(errorFindingConfig)
	}

	repository.NewConnection()

	route := gin.Default()

	route.Group("/appointments")

	route.Run(":5000")
}
