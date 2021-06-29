package crons

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"gorm.io/gorm"
)

func Spawn(db *gorm.DB, documentCache repository.DocumentCache) []CRON {
	h := HospitalSearchIndexer{}
	h.Initialize(db, models.Hospital{})
	h.AddDocumentCache(documentCache)

	return []CRON{h}
}
