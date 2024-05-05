package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tun43p/tun43p.com/internal/api/urls"
)

func CreateDatabase() *gorm.DB {
	dbFile := os.Getenv("API_DATABASE")

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&urls.URLResponse{}); err != nil {
		panic("Failed to migrate database")
	}

	return db
}
