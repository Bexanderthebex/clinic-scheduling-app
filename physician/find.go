package physician

import "gorm.io/gorm"

func FindByLastName(db *gorm.DB, lastName string) *Physician {
	var physician Physician
	db.Find(&physician, &Physician{LastName: lastName})
	return &physician
}
