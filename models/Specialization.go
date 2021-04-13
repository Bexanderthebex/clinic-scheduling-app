package models

type Specialization struct {
	Id                 string `gorm:column:"id"`
	SpecializationName string `gorm:column:"specialization_name"`
}

// TableName overrides the table name used by User to `profiles`
func (Specialization) TableName() string {
	return "specializations"
}
