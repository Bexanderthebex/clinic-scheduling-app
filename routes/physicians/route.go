package physicians

import (
	"encoding/json"
	add_hospital_affiliations "github.com/Bexanderthebex/clinic-scheduling-app/physician/add-hospital-affiliations"
	modify_specializations "github.com/Bexanderthebex/clinic-scheduling-app/physician/modify-specializations"
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
	addSpecializations(physiciansRoute, db)
	addHospitalAffiliations(physiciansRoute, db)
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

func addSpecializations(group *gin.RouterGroup, db *gorm.DB) {
	group.POST("/:physicianId/specializations", func(c *gin.Context) {
		jsonData, _ := c.GetRawData()
		physicianId := c.Param("physicianId")
		var physicianSpecializationsReqBody PhysicianSpecializationRequestBody
		json.Unmarshal(jsonData, &physicianSpecializationsReqBody)

		v := routes.NewValidator()
		validationErrors := v.Struct(physicianSpecializationsReqBody)

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

		req := &modify_specializations.Request{
			PhysicianId:     physicianId,
			Specializations: physicianSpecializationsReqBody.Specializations,
		}

		res := modify_specializations.AddSpecializations(db, req)

		if res.Error == nil {
			c.JSON(http.StatusOK, gin.H{"data": res.Data})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		}
	})
}

func addHospitalAffiliations(group *gin.RouterGroup, db *gorm.DB) {
	group.POST("/:physicianId/hospitals", func(c *gin.Context) {
		jsonData, _ := c.GetRawData()
		physicanId := c.Param("physicianId")
		var createHospitalAffiliations CreateHospitalAffiliations
		json.Unmarshal(jsonData, &createHospitalAffiliations)

		v := routes.NewValidator()
		validationErrors := v.Struct(createHospitalAffiliations)

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

		req := &add_hospital_affiliations.Request{
			PhysicianId: physicanId,
			HospitalIDs: createHospitalAffiliations.HospitalIDs,
		}

		res := add_hospital_affiliations.AddHospitalAffiliation(db, req)

		if res.Error == nil {
			c.JSON(http.StatusOK, gin.H{"data": res.Data})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		}
	})
}
