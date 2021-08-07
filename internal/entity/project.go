package entity

import "github.com/google/uuid"

type Project struct {
	ID   uuid.UUID
	Name string
}
