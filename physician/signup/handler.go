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
	Data  *physician.Physician
	Error error
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
		Data:  newPhysician,
		Error: nil,
	}
}
