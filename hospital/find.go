package hospital

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, hospitalID string) *models.Hospital {
	var hospital models.Hospital
	db.First(&hospital, "id = ?", hospitalID)
	return &hospital
}
