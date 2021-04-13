package physician

import (
	"fmt"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

type AddPhysicianSpecializationError struct {
	PhysicianId string
}

func (p *AddPhysicianSpecializationError) Error() string {
	return fmt.Sprintf("Adding specializations for physician with id %s failed", p.PhysicianId)
}

func AddSpecialization(db *gorm.DB, physician *models.Physician, specialization *models.Specialization) error {
	error := db.Model(physician).Association("Specializations").Append(specialization)
	if error != nil {
		return &AddPhysicianSpecializationError{
			PhysicianId: physician.Id,
		}
	}
	return nil
}
