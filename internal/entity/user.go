package entity

import (
	"time"

	"github.com/google/uuid"
)

// User represent a user object
type User struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Password    string
	SlackAPIKey string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
