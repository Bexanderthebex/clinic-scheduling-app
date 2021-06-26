package repository

import (
	"encoding/json"
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"log"
)

// TODO: make this accomodate read and write replicas
func NewConnection(secretsCache *secretcache.Cache) (*gorm.DB, error) {
	secretKey, errorGettingSSMSecret := secretsCache.GetSecretString(config.GetString("SSM_DATABASE_CONFIG_SECRET_STRING"))
	if errorGettingSSMSecret != nil {
		log.Fatalln(errorGettingSSMSecret)
	}

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
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
}

func ConfigureConnectionPool() {}
