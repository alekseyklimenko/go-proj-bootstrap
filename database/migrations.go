package database

import (
	"github.com/alekseyklimenko/go-proj-bootstrap/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Some{})
}
