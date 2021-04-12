package physician

import "gorm.io/gorm"

func FindById(db *gorm.DB, physicianId string) *Physician {
	var physician Physician
	db.First(&physician, physicianId)
	return &physician
}

func FindByLastName(db *gorm.DB, lastName string) *Physician {
	var physician Physician
	db.Find(&physician, &Physician{LastName: lastName})
	return &physician
}
