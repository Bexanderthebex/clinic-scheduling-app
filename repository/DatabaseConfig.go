package repository

import (
	"fmt"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	Port         uint
	Host         string
	DatabaseName string
}

// TODO: add support for ssl mode
// TODO: add support for timezone
func (dbConfig *DatabaseConfig) GetGormDsnString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DatabaseName, dbConfig.Port)
}
