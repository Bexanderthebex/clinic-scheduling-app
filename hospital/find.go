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

func FindOne(db *gorm.DB, conditions map[string]interface{}) *models.Hospital {
	hospitals := &models.Hospital{}
	db.Where(conditions).First(hospitals)
	return hospitals
}
