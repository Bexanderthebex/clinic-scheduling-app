package qualification

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/physician"
	"gorm.io/gorm"
	"time"
)

type Qualification struct {
	gorm.Model
	Id                uint
	PhysicianId       string
	QualificationName string
	InstituteName     string
	ProcurementYear   time.Time
	Physician         physician.Physician `gorm:"foreignKey:PhysicianId"`
}
