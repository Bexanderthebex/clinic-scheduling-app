package create_hospital

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

type Request struct {
	HospitalName string
	City         string
	Address      string
	Latitude     float64
	Longitude    float64
}

type Response struct {
	Data  *models.Hospital
	Error error
}

func CreateHospital(db *gorm.DB, req *Request) *Response {
	newHospital := &models.Hospital{
		Name:    req.HospitalName,
		City:    req.City,
		Address: req.Address,
		Lat:     req.Latitude,
		Long:    req.Longitude,
	}

	hospitalQuery := newHospital.AsMap()
	hospitalExists := hospital.FindOne(db, hospitalQuery)

	if (models.Hospital{}) != *hospitalExists {
		return &Response{
			Error: &hospital.HospitalAlreadyExistsError{
				HospitalID: hospitalExists.Id,
			},
		}
	}

	hospital.Create(db, newHospital)

	return &Response{
		Data:  newHospital,
		Error: nil,
	}
}
