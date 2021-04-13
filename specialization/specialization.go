package specialization

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

func CreateIfNotExists(db *gorm.DB, specializationName string) (*models.Specialization, error) {
	var specialization models.Specialization
	result := db.FirstOrCreate(&specialization, models.Specialization{SpecializationName: specializationName})
	if result.Error != nil {
		return nil, result.Error
	}
	return &specialization, nil
}
