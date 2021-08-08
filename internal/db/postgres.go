package db

import (
	"os"

	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/slatomate/internal/entity"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New returns an instance of store.Service backed by a PG database
func New() (*gorm.DB, error) {

	// Construct the DB Source
	dsn := os.Getenv("POSTGRES_DSN")

	if len(dsn) == 0 {
		return nil, errors.InternalServerError("DB_INIT", "DSN not provided")
	}

	// Connect to the DB

	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entity.User{}, &entity.Organization{})
	return db, nil
}
