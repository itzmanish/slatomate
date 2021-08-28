package repository

import (
	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/utils"
	"gorm.io/gorm"
)

type JobRepository interface {
	CreateJob(uuid.UUID, *entity.Job) error
	GetJob(*entity.Job) (*entity.Job, error)
	GetAllJob(uuid.UUID) ([]*entity.Job, error)
	DeleteJob(*entity.Job) error
}

type jobDB struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobDB{db}
}

func (p *jobDB) CreateJob(orgID uuid.UUID, job *entity.Job) error {
	return p.db.Model(&entity.Organization{ID: orgID}).Association("Jobs").Append(job)
}

func (p *jobDB) GetJob(query *entity.Job) (*entity.Job, error) {
	var job entity.Job
	req := p.db.Where(query).First(&job)
	return &job, utils.TranslateErrors(req)
}

func (p *jobDB) DeleteJob(job *entity.Job) error {
	res := p.db.Table("Jobs").Delete(job)
	return utils.TranslateErrors(res)
}

func (p *jobDB) GetAllJob(orgid uuid.UUID) ([]*entity.Job, error) {
	var job []*entity.Job
	req := p.db.Where(&entity.Job{OrganizationID: orgid}).Find(&job)
	return job, utils.TranslateErrors(req)
}
