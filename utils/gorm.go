package utils

import (
	"github.com/itzmanish/go-micro/v2/errors"
	"gorm.io/gorm"
)

var (
	// ErrNotFound is a 404 Error
	ErrNotFound = errors.NotFound("NOT_FOUND", "Resource not found")
	// ErrDatabase is a 500 error
	ErrDatabase = errors.InternalServerError("DATABASE_ERROR", "A database error occured")
)

// TranslateErrors takes a pointer to a gorm database, gets the errors and transforms
// them into a single micro error which can be returned safely to a handler.
func TranslateErrors(db *gorm.DB) error {
	err := db.Error
	if err == nil {
		return err
	}
	switch err {
	case gorm.ErrRecordNotFound:
		return ErrNotFound
	case gorm.ErrInvalidTransaction:
	case gorm.ErrDryRunModeUnsupported:
	case gorm.ErrUnsupportedDriver:
	case gorm.ErrNotImplemented:
		return ErrDatabase
	}

	return errors.BadRequest("VALIDATION_FAILED", err.Error())
}
