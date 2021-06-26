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

func NewElasticSearchClient(secretsCache *secretcache.Cache) (*elasticsearch.Client, error) {
	secrets, errorGettingSSMSecret := secretsCache.GetSecretString(config.GetString("ES_CONFIG_STRING"))
	if errorGettingSSMSecret != nil {
		return nil, &CreatingNewElasticSearchClientError{
			ErrorMessage: errorGettingSSMSecret.Error(),
		}
	}

	appEsConfig := &ElasticSearchConfig{}
	encodingError := json.Unmarshal([]byte(secrets), appEsConfig)
	if encodingError != nil {
		return nil, &CreatingNewElasticSearchClientError{
			ErrorMessage: encodingError.Error(),
		}
	}
	esConfig := elasticsearch.Config{
		Addresses: strings.Split(appEsConfig.Addresses, ","),
	}

	return elasticsearch.NewClient(esConfig)
}
