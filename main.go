package main

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"github.com/Bexanderthebex/clinic-scheduling-app/routes/hospitals"
	gin "github.com/gin-gonic/gin"
	"log"
	"time"

	"github.com/Bexanderthebex/clinic-scheduling-app/routes/physicians"
)

func main() {
	errorFindingConfig := config.InitiateConfig()
	if errorFindingConfig != nil {
		log.Fatal(errorFindingConfig)
	}

	db, _ := repository.NewConnection()

	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	route := gin.Default()

	physicians.Initialize(route, db)
	hospitals.Initialize(route, db)

	route.Run(":5000")
}
