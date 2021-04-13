package physician

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, physician *models.Physician) {
	db.Create(physician)
}
