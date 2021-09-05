package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/proto/slatomate"
	"github.com/itzmanish/slatomate/utils"
	"gorm.io/gorm"
)

// User represent a user object
type User struct {
	ID            uuid.UUID      `json:"id" gorm:"primary_key; unique; type:uuid;"`
	Name          string         `json:"name" gorm:"type:varchar(100)"`
	Email         string         `json:"email" gorm:"type:varchar(100)"`
	Password      string         `json:"password" gorm:"type:varchar(200)"`
	APIKey        string         `json:"api_key"`
	Organizations []Organization `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// HaveOrg checks if given organization id exists in user's organization
func (user *User) HaveOrg(id uuid.UUID) bool {
	var exist bool
	for _, org := range user.Organizations {
		if org.ID == id {
			exist = true
		}
	}
	return exist
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(tx *gorm.DB) error {
	u := uuid.New()
	user.ID = u
	return nil
}

// SetPassword set hashed password
func (user *User) SetPassword(password string) error {
	hash, err := utils.GeneratePassword(password)
	if err != nil {
		return err
	}
	user.Password = hash
	return nil
}

func (user *User) GenerateAPIKey() error {
	APIKey := utils.RandomString(28)
	if len(APIKey) != 28 {
		return errors.New("failed to generate api key")
	}
	user.APIKey = APIKey
	return nil
}

// ValidatePassword validate users password is same as given password
func (user *User) ValidatePassword(password string) (bool, error) {
	return utils.ComparePassword(password, user.Password)
}

// BeforeSave performs the validations
func (user *User) BeforeSave(tx *gorm.DB) error {
	if user.Email != "" {
		var userWithEmail User
		tx.Where(User{Email: user.Email}).First(&userWithEmail)
		if userWithEmail.ID.String() != "" && userWithEmail.ID != user.ID {
			return errors.New("email has already been taken")
		}
	}
	return nil
}

//SerializeUser converts proto user to user struct
func SerializeUser(in *slatomate.User) User {
	if in == nil {
		return User{}
	}
	user := User{
		Name:  in.Name,
		Email: in.Email,
	}

	return user
}

//DeserializeUser convertsuser to proto user
func DeserializeUser(in *User) slatomate.User {
	var orgs []*slatomate.Organization
	if len(in.Organizations) > 0 {
		for _, org := range in.Organizations {
			o := DeserializeOrganization(&org)
			orgs = append(orgs, &o)
		}
	}
	return slatomate.User{
		Id:        in.ID.String(),
		Name:      in.Name,
		Email:     in.Email,
		ApiKey:    in.APIKey,
		Orgs:      orgs,
		CreatedAt: in.CreatedAt.Format(time.RFC3339),
	}
}
