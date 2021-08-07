package entity

import (
	"time"

	"github.com/google/uuid"
)

// Project represent a project
type Project struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
