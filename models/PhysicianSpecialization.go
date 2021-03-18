package models

type PhysicianSpecialization struct {
	Id               uint `gorm:"primaryKey"`
	PhysicianId      string
	SpecializationId uint
}
