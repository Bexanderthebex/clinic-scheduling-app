package physician_specialization

type PhysicianSpecialization struct {
	Id               uint `gorm:"primaryKey"`
	PhysicianId      string
	SpecializationId uint
}
