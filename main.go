package main

import (
	"encoding/json"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/physician/signup"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"github.com/Bexanderthebex/clinic-scheduling-app/routes"
	gin "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

type PhysicianRequestBody struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name"`
}

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

	route.Group("/appointments")

	route.POST("/physicians", func(c *gin.Context) {

		jsonData, _ := c.GetRawData()
		var physicianReqBody PhysicianRequestBody
		json.Unmarshal(jsonData, &physicianReqBody)

		v := routes.NewValidator()
		validationErrors := v.Struct(physicianReqBody)

		fieldErrors := validationErrors.(validator.ValidationErrors)
		if len(validationErrors.(validator.ValidationErrors)) > 0 {
			validationError := &routes.RouteValidationError{
				ValidationError: fieldErrors[0],
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError.BuildResponseError()})
			return
		}

		req := &signup.Request{
			FirstName:  physicianReqBody.FirstName,
			LastName:   physicianReqBody.LastName,
			MiddleName: physicianReqBody.MiddleName,
		}

		res := signup.Signup(db, req)

		if res.Error == nil {
			c.JSON(http.StatusOK, gin.H{"data": res.Data})
		}

	})

	route.Run(":5000")
}
