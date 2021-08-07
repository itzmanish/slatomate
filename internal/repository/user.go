package repository

import (
	"github.com/itzmanish/slatomate/internal/db"
	"github.com/itzmanish/slatomate/internal/entity"
)

type UserRepository interface {
	CreateUser(*entity.User) (*entity.User, error)
	GetUser(*entity.User) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUser(*entity.User) error
	GetAllUser() ([]*entity.User, error)
}

type userDB struct {
	db *db.PostgresDB
}

func NewUserRepository(db *db.PostgresDB) UserRepository {
	return &userDB{db}
}

func (u *userDB) CreateUser(user *entity.User) (*entity.User, error) {
	return nil, nil
}
func (u *userDB) GetUser(user *entity.User) (*entity.User, error) {
	return nil, nil
}
func (u *userDB) UpdateUser(user *entity.User) (*entity.User, error) {
	return nil, nil
}
func (u *userDB) DeleteUser(user *entity.User) error {
	return nil
}
func (u *userDB) GetAllUser() ([]*entity.User, error) {
	return nil, nil
}
