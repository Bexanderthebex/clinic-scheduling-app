package config

import "github.com/spf13/viper"

func InitiateConfig() error {
	viper.SetConfigFile("config/.env")

	return viper.ReadInConfig()
}

func GetString(key string) string {
	return viper.GetString(key)
}
