package physician

import "github.com/Bexanderthebex/clinic-scheduling-app/specialization"

type Physician struct {
	Id              string                           `gorm:column:"id"`
	FirstName       string                           `gorm:column:"first_name"`
	LastName        string                           `gorm:column:"last_name"`
	MiddleName      string                           `gorm:column:"middle_name"`
	Specializations []*specialization.Specialization `gorm:"many2many:PhysicianSpecialization;"`
	//Hospitals       []*hospital.Hospital             `gorm:"many2many:PhysicianHospitalAffiliation;"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Physician) TableName() string {
	return "physicians"
}
