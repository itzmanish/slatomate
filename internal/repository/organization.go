package repository

import (
	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/utils"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(*entity.Organization) (*entity.Organization, error)
	GetOrganization(*entity.Organization) (*entity.Organization, error)
	GetAllOrganization(uuid.UUID) ([]*entity.Organization, error)
	DeleteOrganization(*entity.Organization) error
}

type organizationDB struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationDB{db}
}

func (p *organizationDB) CreateOrganization(organization *entity.Organization) (*entity.Organization, error) {
	res := p.db.Create(organization)
	return organization, utils.TranslateErrors(res)
}

func (p *organizationDB) GetOrganization(query *entity.Organization) (*entity.Organization, error) {
	var organization entity.Organization
	req := p.db.Where(query).First(&organization)
	return &organization, utils.TranslateErrors(req)
}

func (p *organizationDB) DeleteOrganization(organization *entity.Organization) error {
	res := p.db.Table("organizations").Delete(organization)
	return utils.TranslateErrors(res)
}

func (p *organizationDB) GetAllOrganization(userid uuid.UUID) ([]*entity.Organization, error) {
	var prj []*entity.Organization
	req := p.db.Where(&entity.Organization{UserID: userid}).Find(&prj)
	return prj, utils.TranslateErrors(req)
}
