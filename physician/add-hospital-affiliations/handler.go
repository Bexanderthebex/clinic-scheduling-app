package add_hospital_affiliations

import (
	"fmt"
	"github.com/Bexanderthebex/clinic-scheduling-app/hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	physicianUseCases "github.com/Bexanderthebex/clinic-scheduling-app/physician"
	"gorm.io/gorm"
)

type Request struct {
	PhysicianId string
	HospitalIDs []string
}

type Response struct {
	Data  []*models.Hospital
	Error error
}

type HospitalDoesNotExistError struct {
	HospitalID string
}

func (e *HospitalDoesNotExistError) Error() string {
	return fmt.Sprintf("Hospital with ID #{e.HospitalID} does not exist")
}

func AddHospitalAffiliation(db *gorm.DB, req *Request) *Response {
	response := &Response{}
	tx := db.Begin()
	if tx.Error != nil {
		response.Error = tx.Error
		return response
	}

	physician := physicianUseCases.FindById(tx, req.PhysicianId)

	if physician == nil {
		response.Error = &physicianUseCases.PhysicianNotFoundError{PhysicianId: req.PhysicianId}
	}

	for _, hospitalId := range req.HospitalIDs {
		h := hospital.FindById(db, hospitalId)
		if h == nil {
			response.Error = &HospitalDoesNotExistError{
				HospitalID: hospitalId,
			}
		}

		addHospitalAffiliationError := physicianUseCases.AddHospitalAffiliation(db, physician, h)
		if addHospitalAffiliationError != nil {
			tx.Rollback()
			response.Error = addHospitalAffiliationError
			return response
		}
		response.Data = append(response.Data, h)
	}
	txCommitRes := tx.Commit()
	if txCommitRes.Error != nil {
		response.Data = []*models.Hospital{}
		response.Error = txCommitRes.Error
		return response
	}

	return response
}
