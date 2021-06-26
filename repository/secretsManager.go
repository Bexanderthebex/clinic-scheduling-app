package repository

import (
	"github.com/Bexanderthebex/clinic-scheduling-app/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

func GetSecrets() (*secretcache.Cache, error) {
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
