package hospitals

import (
	"encoding/json"
	create_hospital "github.com/Bexanderthebex/clinic-scheduling-app/hospital/create-hospital"
	search_hospital "github.com/Bexanderthebex/clinic-scheduling-app/hospital/search-hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"github.com/Bexanderthebex/clinic-scheduling-app/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var documentCache repository.DocumentCache

func Initialize(route *gin.Engine, db *gorm.DB) {
	hospitalsRoute := route.Group("/hospitals")

	createHospital(hospitalsRoute, db)
	searchHospital(hospitalsRoute, db)
}

func AddDocumentCache(cache repository.DocumentCache) {
	documentCache = cache
}

func createHospital(group *gin.RouterGroup, db *gorm.DB) {
	group.POST("", func(c *gin.Context) {
		jsonData, _ := c.GetRawData()
		var createHospitalReq CreateHospitalReq
		json.Unmarshal(jsonData, &createHospitalReq)

		v := routes.NewValidator()
		validationError := routes.CheckForErrors(createHospitalReq, v)

		if validationError != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError.BuildResponseError()})
			return
		}

		createHospital := &create_hospital.Request{
			HospitalName: createHospitalReq.HospitalName,
			City:         createHospitalReq.City,
			Address:      createHospitalReq.City,
			Latitude:     createHospitalReq.Latitude,
			Longitude:    createHospitalReq.Longitude,
		}
		res := create_hospital.CreateHospital(db, createHospital)

		if res.Error == nil {
			c.JSON(http.StatusOK, gin.H{"data": res.Data})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		}
	})
}

func searchHospital(group *gin.RouterGroup, db *gorm.DB) {
	group.GET("", func(c *gin.Context) {
		jsonData, _ := c.GetRawData()
		var findHospitalReq FindHospitalReq
		json.Unmarshal(jsonData, &findHospitalReq)

		v := routes.NewValidator()
		validationError := routes.CheckForErrors(findHospitalReq, v)

		if validationError != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError.BuildResponseError()})
			return
		}

		findHospital := &search_hospital.Request{}
		if findHospitalReq.ID != "" {
			findHospital.HospitalID = findHospitalReq.ID
		} else {
			findHospital.HospitalName = findHospitalReq.HospitalName
		}

		res := search_hospital.SearchHospital(db, documentCache, findHospital)

		if res.Error == nil {
			c.JSON(http.StatusOK, gin.H{"data": res.Data})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		}
	})
}
