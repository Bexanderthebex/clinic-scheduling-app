package modify_specializations

import (
	physicianUseCases "github.com/Bexanderthebex/clinic-scheduling-app/physician"
	"github.com/Bexanderthebex/clinic-scheduling-app/specialization"
	"gorm.io/gorm"
)

type Request struct {
	PhysicianId     string
	Specializations []*specialization.Specialization
}

type Response struct {
	Data  string
	Error error
}

func AddSpecializations(db *gorm.DB, req *Request) *Response {
	physician := physicianUseCases.FindById(db, req.PhysicianId)

	appendSpecializationError := db.Model(physician).Association("Specializations").Append(req.Specializations)

	return &Response{
		Data:  "sample",
		Error: appendSpecializationError,
	}
}
