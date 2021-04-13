package physician

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, physicianId string) *models.Physician {
	var physician models.Physician
	db.First(&physician, "id = ?", physicianId)
	return &physician
}

func FindByLastName(db *gorm.DB, lastName string) *models.Physician {
	var physician models.Physician
	db.Find(&physician, &models.Physician{LastName: lastName})
	return &physician
}
