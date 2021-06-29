package main

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/crons"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"github.com/Bexanderthebex/clinic-scheduling-app/routes/hospitals"
	gin "github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"time"

	"github.com/Bexanderthebex/clinic-scheduling-app/routes/physicians"
)

func main() {
	errorFindingConfig := config.InitiateConfig()
	if errorFindingConfig != nil {
		log.Fatal(errorFindingConfig)
	}

	secretsCache, createSecretsCacheError := repository.GetSecrets()
	if createSecretsCacheError != nil {
		log.Fatal(createSecretsCacheError.Error())
	}
	db, _ := repository.NewConnection(secretsCache)
	SearchCache, createSeachCacheError := repository.NewElasticSearchClient(secretsCache)
	if createSecretsCacheError != nil {
		log.Fatal(createSeachCacheError)
	}

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
	hospitals.AddDocumentCache(SearchCache)

	cronList := crons.Spawn(db, SearchCache)

	c := cron.New()
	for _, cronObj := range cronList {
		c.AddFunc("@every 10s", cronObj.Run)
	}

	// Run CRON jobs on a different go routine
	go func() {
		c.Start()
	}()

	route.Run(":5000")
}
