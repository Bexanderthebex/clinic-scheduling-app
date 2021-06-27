package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"github.com/elastic/go-elasticsearch/v7"
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

type ElasticSearchCache struct {
	esClient *elasticsearch.Client
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
