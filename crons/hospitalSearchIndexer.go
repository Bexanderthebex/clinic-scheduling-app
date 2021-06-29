package crons

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/hospital"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"gorm.io/gorm"
	"log"
)

type HospitalSearchIndexer struct {
	store         *gorm.DB
	hospitalModel interface{}
	documentCache repository.DocumentCache
}

func (h HospitalSearchIndexer) Run() {
	elasticSearchHopistalIndexName := config.GetString("ES_HOSPITAL_INDEX_NAME")
	exists, headIndexError := h.documentCache.IndexExists(elasticSearchHopistalIndexName)
	if headIndexError != nil {
		log.Println(headIndexError)
		panic(headIndexError)
	}

	if !exists {
		createIndexResult, createIndexError := h.documentCache.CreateIndex(elasticSearchHopistalIndexName)
		if createIndexError != nil {
			panic(createIndexError)
		}
		log.Println(createIndexResult)
	}

	hospitals, findAllHospitalsError := hospital.FindAll(h.store)
	if findAllHospitalsError != nil {
		panic(findAllHospitalsError)
	}

	for _, hospital := range hospitals {
		h.documentCache = h.documentCache.AddBulkIndexAction(hospital.AsMap(), elasticSearchHopistalIndexName)
	}
	h.documentCache.ExecuteBulkActions()
}

func (h *HospitalSearchIndexer) Initialize(db *gorm.DB, hospitalModel interface{}) {
	h.store = db
	h.hospitalModel = hospitalModel
}

func (h *HospitalSearchIndexer) AddDocumentCache(dc repository.DocumentCache) {
	h.documentCache = dc
}
