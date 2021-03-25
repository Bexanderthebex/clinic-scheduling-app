package repository

import (
	"encoding/json"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"log"
)

func createSecretCache() (*secretcache.Cache, error) {
	return secretcache.New(
		func(cache *secretcache.Cache) {
			cache.CacheConfig = secretcache.CacheConfig{
				MaxCacheSize: secretcache.DefaultMaxCacheSize + 10,
				VersionStage: secretcache.DefaultVersionStage,
				CacheItemTTL: secretcache.DefaultCacheItemTTL,
			}
		},
		func(cache *secretcache.Cache) {
			clientSession, _ := session.NewSession(
				&aws.Config{
					Region:      aws.String(config.GetString("AWS_REGION")),
					Credentials: credentials.NewSharedCredentials("", config.GetString("AWS_PROFILE")),
				})
			client := secretsmanager.New(clientSession)
			cache.Client = client
		},
	)
}

// TODO: make this accomodate read and write replicas
func NewConnection() (*gorm.DB, error) {
	secretCache, errorCreatingCache := createSecretCache()

	if errorCreatingCache != nil {
		log.Fatalln(errorCreatingCache)
	}

	secretKey, errorGettingSSMSecret := secretCache.GetSecretString(config.GetString("SSM_DATABASE_CONFIG_SECRET_STRING"))
	if errorGettingSSMSecret != nil {
		log.Fatalln(errorGettingSSMSecret)
	}

	log.Println(secretKey)

	var databaseConfig DatabaseConfig
	err := json.Unmarshal([]byte(secretKey), &databaseConfig)

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(databaseConfig.GetGormDsnString())

	return gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  databaseConfig.GetGormDsnString(),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			}),
		&gorm.Config{},
	)
}
