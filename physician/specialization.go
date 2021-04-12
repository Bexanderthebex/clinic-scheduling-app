package physician

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/specialization"
	"gorm.io/gorm"
)

func AddSpecialization(db *gorm.DB, physician *Physician, specialization *specialization.Specialization) error {
	res := db.Model(physician).Association("Specializations").Append(specialization)
	return res
}
