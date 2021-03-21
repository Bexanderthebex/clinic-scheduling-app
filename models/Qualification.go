package models

import (
	"gorm.io/gorm"
	"time"
)

type Qualification struct {
	gorm.Model
	Id                uint
	PhysicianId       string
	QualificationName string
	InstituteName     string
	ProcurementYear   time.Time
	Physician         Physician `gorm:"foreignKey:PhysicianId"`
}
