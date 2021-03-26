package models

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/physician"
	"gorm.io/gorm"
)

type Specialization struct {
	gorm.Model
	Id                 int
	UniqueCode         string
	SpecializationName string
	Physicians         []*physician.Physician `gorm:"many2many:PhysicianSpecialization"`
}
