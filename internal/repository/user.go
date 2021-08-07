package repository

import (
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*entity.User) (*entity.User, error)
	GetUser(*entity.User) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUser(*entity.User) error
	GetAllUser() ([]*entity.User, error)
}

type userDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userDB{db}
}

func (u *userDB) CreateUser(user *entity.User) (*entity.User, error) {
	res := u.db.Create(&user)
	return user, utils.TranslateErrors(res)
}
func (u *userDB) GetUser(query *entity.User) (*entity.User, error) {
	var user entity.User
	req := u.db.Where(query).First(&user)
	return &user, utils.TranslateErrors(req)
}
func (u *userDB) UpdateUser(params *entity.User) (*entity.User, error) {
	user, err := u.GetUser(&entity.User{ID: params.ID})
	if err != nil {
		return params, err
	}
	req := u.db.Model(&user).Updates(params)
	return user, utils.TranslateErrors(req)
}
func (u *userDB) DeleteUser(user *entity.User) error {
	res := u.db.Table("users").Delete(user)
	return utils.TranslateErrors(res)
}
func (u *userDB) GetAllUser() ([]*entity.User, error) {
	var users []*entity.User
	req := u.db.Model(&entity.User{}).Find(&users)
	return users, utils.TranslateErrors(req)
}
