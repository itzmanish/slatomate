package repository

import (
	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/utils"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(uuid.UUID, *entity.Organization) error
	GetOrganization(*entity.Organization) (*entity.Organization, error)
	GetAllOrganization(uuid.UUID) ([]*entity.Organization, error)
	DeleteOrganization(*entity.Organization) error
	UpdateOrganization(*entity.Organization) (*entity.Organization, error)
}

type organizationDB struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationDB{db}
}

func (p *organizationDB) CreateOrganization(userID uuid.UUID, organization *entity.Organization) error {
	return p.db.Model(&entity.User{ID: userID}).Association("Organizations").Append(organization)
}

func (p *organizationDB) GetOrganization(query *entity.Organization) (*entity.Organization, error) {
	var organization entity.Organization
	req := p.db.Preload("Jobs").Where(query).First(&organization)
	return &organization, utils.TranslateErrors(req)
}

func (p *organizationDB) DeleteOrganization(organization *entity.Organization) error {
	res := p.db.Table("organizations").Where(organization).Delete(organization)
	return utils.TranslateErrors(res)
}

func (p *organizationDB) UpdateOrganization(params *entity.Organization) (*entity.Organization, error) {
	org, err := p.GetOrganization(&entity.Organization{ID: params.ID})
	if err != nil {
		return params, err
	}
	req := p.db.Model(&org).Updates(params)
	return org, utils.TranslateErrors(req)
}

func (p *organizationDB) GetAllOrganization(userid uuid.UUID) ([]*entity.Organization, error) {
	var prj []*entity.Organization
	req := p.db.Preload("Jobs").Where(&entity.Organization{UserID: userid}).Find(&prj)
	return prj, utils.TranslateErrors(req)
}
