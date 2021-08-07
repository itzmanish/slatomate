package repository

import (
	"github.com/itzmanish/slatomate/internal/entity"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	CreateProject(*entity.Project) (*entity.Project, error)
	GetProject(*entity.Project) (*entity.Project, error)
	GetAllProjct() ([]*entity.Project, error)
	DeleteProject(*entity.Project) error
}

type projectDB struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectDB{db}
}

func (db *projectDB) CreateProject(project *entity.Project) (*entity.Project, error) {
	return nil, nil
}

func (db *projectDB) GetProject(project *entity.Project) (*entity.Project, error) {
	return nil, nil
}

func (db *projectDB) DeleteProject(project *entity.Project) error {
	return nil
}

func (db *projectDB) GetAllProjct() ([]*entity.Project, error) {
	return nil, nil
}
