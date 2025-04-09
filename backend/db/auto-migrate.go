package db

import (
	"github.com/gaelzamora/ticket-project/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}