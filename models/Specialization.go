package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Specialization struct {
	Id                 string `gorm:column:"id"`
	SpecializationName string `gorm:column:"specialization_name"`
}

// TableName overrides the table name used by User to `profiles`
func (Specialization) TableName() string {
	return "specializations"
}

func (s *Specialization) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Id == "" {
		s.Id = uuid.NewString()
	}

	return
}
