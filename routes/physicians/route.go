package physicians

import (
	"encoding/json"
	"github.com/Bexanderthebex/clinic-scheduling-app/physician/signup"
	"github.com/Bexanderthebex/clinic-scheduling-app/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

func Initialize(route *gin.Engine, db *gorm.DB) {
	physiciansRoute := route.Group("/physicians")

	createPhysician(physiciansRoute, db)
}

type PhysicianRequestBody struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name"`
}

func createPhysician(group *gin.RouterGroup, db *gorm.DB) {
	group.POST("", func(c *gin.Context) {

		jsonData, _ := c.GetRawData()
		var physicianReqBody PhysicianRequestBody
		json.Unmarshal(jsonData, &physicianReqBody)

		v := routes.NewValidator()
		validationErrors := v.Struct(physicianReqBody)

		if validationErrors != nil {
			fieldErrors := validationErrors.(validator.ValidationErrors)
			if len(validationErrors.(validator.ValidationErrors)) > 0 {
				validationError := &routes.RouteValidationError{
					ValidationError: fieldErrors[0],
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError.BuildResponseError()})
				return
			}
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
}
