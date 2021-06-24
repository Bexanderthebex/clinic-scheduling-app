package models

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type PhysicianHospitalAffiliation struct {
	Id          string    `gorm:column:"id"`
	PhysicianId string    `gorm:column:"physician_id"`
	HospitalId  string    `gorm:column:"hospital_id"`
	StartDate   time.Time `gorm:column:"start_date"`
	EndDate     time.Time `gorm:column:"end_date"`
}

// TableName overrides the table name used by User to `profiles`
func (PhysicianHospitalAffiliation) TableName() string {
	return "physician_hospitals"
}

func (h *PhysicianHospitalAffiliation) BeforeCreate(tx *gorm.DB) (err error) {
	if h.Id == "" {
		h.Id = uuid.NewString()
	}

	return
}
