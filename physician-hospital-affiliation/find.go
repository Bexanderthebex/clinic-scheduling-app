package physician_hospital_affiliation

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, physicianID string, hospitalID string) *models.PhysicianHospitalAffiliation {
	var physicianHospitalAffiliation models.PhysicianHospitalAffiliation
	db.First(&physicianHospitalAffiliation, "physician_id = ? and hospital_id = ?", physicianID, hospitalID)
	return &physicianHospitalAffiliation
}
