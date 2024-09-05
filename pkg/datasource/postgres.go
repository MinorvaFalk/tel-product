package datasource

import (
	"log"
	"tel/product/config"
	"tel/product/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.ReadConfig().Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.NewGormLogger(),
	})
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql db instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping sql db: %v", err)
	}

	return db
}
