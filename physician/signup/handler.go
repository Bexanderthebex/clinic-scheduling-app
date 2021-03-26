package signup

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/physician"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Request struct {
	FirstName  string
	LastName   string
	MiddleName string
}

type Response struct {
	StatusCode uint
	Message    string
	Data       *physician.Physician
}

func Signup(db *gorm.DB, req *Request) *Response {
	newPhysician := &physician.Physician{
		Id:         uuid.NewString(),
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
	}

	physician.Create(db, newPhysician)

	return &Response{
		StatusCode: 200,
		Message:    "Physician successfully created",
		Data:       newPhysician,
	}
}
