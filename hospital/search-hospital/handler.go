package search_hospital

import (
	"encoding/json"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"gorm.io/gorm"
)

type Request struct {
	HospitalID   string
	HospitalName string
}

type Response struct {
	Data  *models.Hospital
	Error error
}

func SearchHospital(db *gorm.DB, searchCache repository.DocumentCache, request *Request) *Response {
	if request.HospitalID != "" {
		findHospitalResult := hospital.FindById(db, request.HospitalID)
		return &Response{
			Data: findHospitalResult,
		}
	} else {
		matchHospitalQuery := searchCache.CreateQueryStatement("name", request.HospitalName)

		searchRes, searcErr := searchCache.Find(matchHospitalQuery, config.GetString("ES_HOSPITAL_INDEX_NAME"))
		if searcErr != nil {
			return &Response{
				Error: searcErr,
			}
		}

		nearestHospitalSearchMatch := &models.Hospital{}
		total := int(searchRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
		if total == 0 {
			return &Response{
				Data:  nil,
				Error: nil,
			}
		}
		// TODO: Clean this one up
		jsonString, jsonStringConversionErr := json.Marshal(searchRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"])
		if jsonStringConversionErr != nil {
			return &Response{
				Error: jsonStringConversionErr,
			}
		}

		json.Unmarshal(jsonString, nearestHospitalSearchMatch)
		return &Response{
			Data:  nearestHospitalSearchMatch,
			Error: nil,
		}
	}

	return nil
}
