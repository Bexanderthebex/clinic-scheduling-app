package get_hospital

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

type Request struct {
	HospitalId string
}

type Response struct {
	Data  *models.Hospital
	Error error
}

func GetHospital(db *gorm.DB, request *Request) *Response {
	h := hospital.FindById(db, request.HospitalId)

	if h == nil {
		return &Response{
			Data:  nil,
			Error: &hospital.HospitalDoesNotExistError{HospitalID: ""},
		}
	}

	return &Response{
		Data:  h,
		Error: nil,
	}
}
