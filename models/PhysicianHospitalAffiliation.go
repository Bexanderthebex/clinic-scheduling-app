package models

import (
	"time"

	"gorm.io/gorm"
)

type PhysicianHospitalAffiliation struct {
	gorm.Model
	Id          uint
	PhysicianId string
	HospitalId  string
	StartDate   time.Time
	EndDate     time.Time
}
