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
	hospital := &models.Hospital{}
	db.Where(conditions).First(hospital)
	return hospital
}

func FindAll(db *gorm.DB) ([]*models.Hospital, error) {
	hospitals := make([]*models.Hospital, 0)
	result := db.Find(&hospitals)
	if result.Error != nil {
		return nil, result.Error
	}
	return hospitals, nil
}
