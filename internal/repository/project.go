package repository

import (
	"github.com/itzmanish/slatomate/internal/db"
	"github.com/itzmanish/slatomate/internal/entity"
)

type ProjectRepository interface {
	CreateProject(*entity.Project) (*entity.Project, error)
	GetProject(*entity.Project) (*entity.Project, error)
	GetAllProjct() ([]*entity.Project, error)
	DeleteProject(*entity.Project) error
}

type projectDB struct {
	db *db.PostgresDB
}

func NewProjectRepository(db *db.PostgresDB) ProjectRepository {
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
