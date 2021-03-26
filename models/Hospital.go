package models

import "github.com/Bexanderthebex/clinic-scheduling-app/physician"

type Hospital struct {
	Id         string `gorm:"primaryKey"`
	Name       string
	City       string
	Address    string
	Lat        float64
	Long       float64
	Logo       string
	Physicians []*physician.Physician `gorm:"many2many:PhysicianHospitalAffiliation;"`
}
