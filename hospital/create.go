package hospital

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, hospital *models.Hospital) {
	db.Create(hospital)
}
