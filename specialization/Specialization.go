package specialization

import (
	"gorm.io/gorm"
)

type Specialization struct {
	gorm.Model
	Id                 int
	UniqueCode         string
	SpecializationName string
}
