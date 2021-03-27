package physician

type Physician struct {
	Id         string `gorm:coumn:"id"`
	FirstName  string `gorm:column:"first_name'`
	LastName   string `gorm:column:"last_name"`
	MiddleName string `gorm:column:"middle_name"`
	// TODO: specify this at run time
	//Specializations []*specialization.Specialization `gorm:"many2many:PhysicianSpecialization;"`
	//Hospitals       []*hospital.Hospital             `gorm:"many2many:PhysicianHospitalAffiliation;"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Physician) TableName() string {
	return "physician"
}
