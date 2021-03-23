package main

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	gin "github.com/gin-gonic/gin"
	"log"
)

func main() {
	errorFindingConfig := config.InitiateConfig()
	if errorFindingConfig != nil {
		log.Fatal(errorFindingConfig)
	}

	route := gin.Default()

	route.Group("/appointments")

	route.Run(":5000")
}
