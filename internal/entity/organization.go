package entity

import (
	"time"

	"github.com/google/uuid"
	slatomate "github.com/itzmanish/slatomate/proto/gen/slatomate/v1"
	"gorm.io/gorm"
)

// Organization represent an organization
type Organization struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key; unique; type:uuid;"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Name        string    `json:"name" gorm:"type:varchar(100)"`
	SlackAPIKey string
	Jobs        []Job
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (org *Organization) BeforeCreate(tx *gorm.DB) error {
	u := uuid.New()
	org.ID = u
	return nil
}

//SerializeOrganization converts proto Organization to Organization struct
func SerializeOrganization(in *slatomate.Organization) Organization {
	if in == nil {
		return Organization{}
	}
	Organization := Organization{
		Name:        in.Name,
		SlackAPIKey: in.SlackApikey,
	}

	return Organization
}

//DeserializeOrganization converts Organization to proto Organization
func DeserializeOrganization(in *Organization) slatomate.Organization {
	return slatomate.Organization{
		Id:          in.ID.String(),
		Name:        in.Name,
		SlackApikey: in.SlackAPIKey,
		CreatedAt:   in.CreatedAt.Format(time.RFC3339),
	}
}
