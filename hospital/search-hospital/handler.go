package search_hospital

import (
	"encoding/json"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"github.com/Bexanderthebex/clinic-scheduling-app/models/elasticsearch"
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

		esr := elasticsearch.ElasticSearchResponse{Response: searchRes}
		nearestHospitalSearchMatch := &models.Hospital{}
		if esr.TotalHits() == 0 {
			return &Response{
				Data:  nil,
				Error: nil,
			}
		}

		jsonString, jsonStringConversionErr := json.Marshal(esr.First())
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
