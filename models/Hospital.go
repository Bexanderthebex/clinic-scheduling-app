package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hospital struct {
	Id      string  `gorm:column:"id"`
	Name    string  `gorm:column:"name"`
	City    string  `gorm:column:"city"'`
	Address string  `gorm:column:"address"`
	Lat     float64 `gorm:column:"lat"`
	Long    float64 `gorm:column:"long"`
	Logo    string  `gorm:column:"logo"`
}

// TableName overrides the table name used by User to `profiles`
func (Hospital) TableName() string {
	return "hospitals"
}

func (h *Hospital) BeforeCreate(tx *gorm.DB) (err error) {
	if h.Id == "" {
		h.Id = uuid.NewString()
	}

	return
}
