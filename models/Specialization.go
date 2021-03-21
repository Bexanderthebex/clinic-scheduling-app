package models

import "gorm.io/gorm"

type Specialization struct {
	gorm.Model
	Id                 int
	UniqueCode         string
	SpecializationName string
	Physicians         []*Physician `gorm:"many2many:PhysicianSpecialization"`
}
