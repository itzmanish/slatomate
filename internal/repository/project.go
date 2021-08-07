package repository

import "github.com/itzmanish/slatomate/internal/entity"

type ProjectRepository interface {
	CreateProject(*entity.Project) (*entity.Project, error)
	GetProject(*entity.Project) (*entity.Project, error)
	GetAllProjct() ([]*entity.Project, error)
}

type projectDB struct{}

func (db *projectDB) CreateProject(project *entity.Project) (*entity.Project, error) {
	return nil, nil
}
func (db *projectDB) GetProject(project *entity.Project) (*entity.Project, error) {
	return nil, nil
}
func (db *projectDB) GetAllProjct() (*entity.Project, error) {
	return nil, nil
}
