package models

type Physician struct {
	Id              string            `gorm:column:"id"`
	FirstName       string            `gorm:column:"first_name"`
	LastName        string            `gorm:column:"last_name"`
	MiddleName      string            `gorm:column:"middle_name"`
	Specializations []*Specialization `gorm:"many2many:physician_specializations;"`
	Hospitals       []*Hospital       `gorm:"many2many:physician_hospitals;"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Physician) TableName() string {
	return "physicians"
}
