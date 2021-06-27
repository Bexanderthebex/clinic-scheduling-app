package crons

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/Bexanderthebex/clinic-scheduling-app/models"
	"github.com/Bexanderthebex/clinic-scheduling-app/repository"
	"log"
)

type HospitalSearchIndexer struct {
	store         interface{}
	hospitalModel models.Hospital
	documentCache repository.DocumentCache
}

func (h *HospitalSearchIndexer) Run() {
	// Check if the index already exists

	exists, headIndexError := h.documentCache.IndexExists(config.GetString("ES_HOSPITAL_INDEX_NAME"))
	if headIndexError != nil {
		panic(headIndexError)
	}

	if !exists {
		createIndexResult, createIndexError := h.documentCache.CreateIndex(config.GetString("ES_HOSPITAL_INDEX_NAME"))
		if createIndexError != nil {
			panic(createIndexError)
		}
		log.Println(createIndexResult)
	}

}

func (h *HospitalSearchIndexer) Initialize(db interface{}, hospitalModel models.Hospital) {
	h.store = db
	h.hospitalModel = hospitalModel
}

func (h *HospitalSearchIndexer) AddDocumentCache(dc repository.DocumentCache) {
	h.documentCache = dc
}
