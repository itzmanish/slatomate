package db

import (
	"os"

	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/slatomate/internal/entity"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct{ db *gorm.DB }

// New returns an instance of store.Service backed by a PG database
func New() (*PostgresDB, error) {

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
	db.AutoMigrate(&entity.User{}, &entity.Project{})
	return &PostgresDB{db}, nil
}

// Close will terminate the database connection
func (p PostgresDB) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// String will return database driver
func (p PostgresDB) String() string {
	return "postgres"
}
