package modify_specializations

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	physicianUseCases "github.com/Bexanderthebex/clinic-scheduling-app/physician"
	"github.com/Bexanderthebex/clinic-scheduling-app/specialization"
	"gorm.io/gorm"
)

type Request struct {
	PhysicianId     string
	Specializations []string
}

type Response struct {
	Data  []*models.Specialization
	Error error
}

func AddSpecializations(db *gorm.DB, req *Request) *Response {
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

	for _, specializationName := range req.Specializations {
		s, createSpecializationError := specialization.CreateIfNotExists(tx, specializationName)
		if createSpecializationError != nil {
			tx.Rollback()
			response.Error = createSpecializationError
			return response
		}

		addSpecializationToPhysicianError := physicianUseCases.AddSpecialization(tx, physician, s)
		if addSpecializationToPhysicianError != nil {
			tx.Rollback()
			response.Error = addSpecializationToPhysicianError
			return response
		}
		response.Data = append(response.Data, s)
	}
	txCommitResult := tx.Commit()
	if txCommitResult != nil {
		response.Error = txCommitResult.Error
		return response
	}

	return response
}
