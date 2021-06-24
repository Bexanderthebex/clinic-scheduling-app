package physician

import (
	"fmt"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	physician_hospital_affiliation "github.com/Bexanderthebex/clinic-scheduling-app/physician-hospital-affiliation"
	"gorm.io/gorm"
)

type CreateHospitalAffiliationError struct {
	PhysicianID string
	HospitalID  string
}

func (h *CreateHospitalAffiliationError) Error() string {
	return fmt.Sprintf("Adding physician hospital affiliation to Physician with ID #{h.PhyisicianID} failed")
}

type PhysicianHospitalAffiliationAlreadyExistsError struct {
	PhysicianID string
	HospitalID  string
}

func (a *PhysicianHospitalAffiliationAlreadyExistsError) Error() string {
	return fmt.Sprintf("Physician Hospital Affiliation already exists")
}

func AddHospitalAffiliation(db *gorm.DB, physician *models.Physician, hospital *models.Hospital) error {
	physicianHospitalAffiliation := physician_hospital_affiliation.FindById(db, physician.Id, hospital.Id)

	if physicianHospitalAffiliation.PhysicianId != "" {
		return &PhysicianHospitalAffiliationAlreadyExistsError{
			PhysicianID: physician.Id,
			HospitalID:  hospital.Id,
		}
	}

	error := db.Model(physician).Association("Hospitals").Append(hospital)
	if error != nil {
		return &CreateHospitalAffiliationError{
			PhysicianID: physician.Id,
			HospitalID:  hospital.Id,
		}
	}
	return nil
}
