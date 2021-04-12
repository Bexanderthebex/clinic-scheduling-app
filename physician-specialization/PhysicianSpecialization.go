package physician_specialization

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PhysicianSpecialization struct {
	Id               string `gorm:column:"id";`
	PhysicianId      string
	SpecializationId string
}

func (ps PhysicianSpecialization) BeforeCreate(tx *gorm.DB) {
	ps.Id = uuid.NewString()

	return
}
