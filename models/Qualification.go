package models

import "time"

type Qualification struct {
	Id                uint
	PhysicianId       string
	QualificationName string
	InstituteName     string
	ProcurementYear   time.Time
}
