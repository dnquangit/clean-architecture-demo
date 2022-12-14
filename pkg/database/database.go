package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PostgresDB(connectStr string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connectStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
