package database

import (
	"github.com/alekseyklimenko/go-proj-bootstrap/config"
	"github.com/alekseyklimenko/go-proj-bootstrap/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func New(conf *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.Database.Dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		logger.NewEntry().WithError(err).Fatal("Failed to connect db")
	}
	return db
}
