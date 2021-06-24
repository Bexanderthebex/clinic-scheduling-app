package physician

import (
	"fmt"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

type CreateHospitalAffiliationError struct {
	PhysicianID string
	HospitalID  string
}

func (h *CreateHospitalAffiliationError) Error() string {
	return fmt.Sprintf("Adding physician hospital affiliation to Physician with ID #{h.PhyisicianID} failed")
}

func AddHospitalAffiliation(db *gorm.DB, physician *models.Physician, hospital *models.Hospital) error {
	error := db.Model(physician).Association("Hospitals").Append(hospital)
	if error != nil {
		return &CreateHospitalAffiliationError{
			PhysicianID: physician.Id,
			HospitalID:  hospital.Id,
		}
	}
	return nil
}
