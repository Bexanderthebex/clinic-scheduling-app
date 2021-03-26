package physician

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
)

type Physician struct {
	Id              string
	FirstName       string
	LastName        string
	MiddleName      string
	Specializations []*models.Specialization `gorm:"many2many:PhysicianSpecialization;"`
	Hospitals       []*models.Hospital       `gorm:"many2many:PhysicianHospitalAffiliation;"`
}
