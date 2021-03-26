package physician

import "gorm.io/gorm"

func Create(db *gorm.DB, physician *Physician) {
	db.Create(physician)
}
