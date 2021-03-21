package models

import "gorm.io/gorm"

type Physician struct {
	gorm.Model
	Id              string
	FirstName       string
	LastName        string
	MiddleName      string
	Specializations []*Specialization `gorm:"many2many:PhysicianSpecialization;"`
	Hospitals       []*Hospital       `gorm:"many2many:PhysicianHospitalAffiliation;"`
}
