package physician

import "gorm.io/gorm"

func Create(db *gorm.DB, physician *Physician) {
	transaction := db.Create(physician)
	transaction.Commit()
}
