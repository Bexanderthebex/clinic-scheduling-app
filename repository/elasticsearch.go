package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"strings"
)

type ElasticSearchConfig struct {
	Addresses string `json:"addresses"`
}

type CreatingNewElasticSearchClientError struct {
	ErrorMessage string
}

func (e *CreatingNewElasticSearchClientError) Error() string {
	return fmt.Sprintf("Error initializing ElasticSearch client %s", e.ErrorMessage)
}

type ElasticSearchActionItem struct {
	Action     string
	Item       map[string]interface{}
	DocumentId string
	IndexName  string
}

type ElasticSearchCache struct {
	esClient *elasticsearch.Client
	actions  []ElasticSearchActionItem
	index    string
}

func (es ElasticSearchCache) Find(map[string]interface{}) map[string]interface{} {
	return make(map[string]interface{})
}

func (es ElasticSearchCache) IndexExists(indexName string) (bool, error) {
	response, checkIndexExistsError := es.esClient.Indices.Exists([]string{indexName})
	if checkIndexExistsError != nil {
		return false, checkIndexExistsError
	}

	if response.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}

func (es ElasticSearchCache) CreateIndex(indexName string) (map[string]interface{}, error) {
	response, createIndexError := es.esClient.Indices.Create(indexName)
	if createIndexError != nil {
		return nil, createIndexError
	}

	mapResponse := make(map[string]interface{})
	json.NewDecoder(response.Body).Decode(&mapResponse)

	return mapResponse, createIndexError
}

func (es ElasticSearchCache) ExecuteBulkActions() {
	bulkApi, createBulkIndexerError := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{})
	if createBulkIndexerError != nil {
		panic(createBulkIndexerError)
	}

	defaultContext := context.Background()
	for _, a := range es.actions {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(a.Item)
		item := esutil.BulkIndexerItem{
			Index:           a.IndexName,
			Action:          a.Action,
			DocumentID:      a.DocumentId,
			Body:            b,
			RetryOnConflict: nil,
			OnSuccess:       nil,
			OnFailure:       nil,
		}
		addToBulkApiError := bulkApi.Add(defaultContext, item)
		if addToBulkApiError != nil {
			log.Println(addToBulkApiError)
			panic(addToBulkApiError)
		}
	}

	bulkApi.Close(defaultContext)
}

func (es ElasticSearchCache) AddBulkIndexAction(item map[string]interface{}, indexName string) DocumentCache {
	esItem := &ElasticSearchActionItem{
		Action:    "index",
		Item:      item,
		IndexName: indexName,
	}
	if item["id"] != nil || item["id"] != "" {
		esItem.DocumentId = fmt.Sprintf("%s", item["id"])
	}
	es.actions = append(es.actions, *esItem)
	return es
}

func NewElasticSearchClient(secretsCache *secretcache.Cache) (ElasticSearchCache, error) {
	secrets, errorGettingSSMSecret := secretsCache.GetSecretString(config.GetString("ES_CONFIG_STRING"))
	if errorGettingSSMSecret != nil {
		return ElasticSearchCache{}, &CreatingNewElasticSearchClientError{
			ErrorMessage: errorGettingSSMSecret.Error(),
		}
	}

	appEsConfig := &ElasticSearchConfig{}
	encodingError := json.Unmarshal([]byte(secrets), appEsConfig)
	if encodingError != nil {
		return ElasticSearchCache{}, &CreatingNewElasticSearchClientError{
			ErrorMessage: encodingError.Error(),
		}
	}
	esConfig := elasticsearch.Config{
		Addresses: strings.Split(appEsConfig.Addresses, ","),
	}

	esClient, createSearchCacheError := elasticsearch.NewClient(esConfig)
	if createSearchCacheError != nil {
		return ElasticSearchCache{}, createSearchCacheError
	}

	return ElasticSearchCache{esClient: esClient}, nil
}
