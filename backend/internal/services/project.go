package service

import (
	"context"
	"github.com/Max425/manager/internal/model/core"
	"log/slog"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *core.Project) (*core.Project, error)
	FindProjectByID(ctx context.Context, id int) (*core.Project, error)
	FindProjectsByCompanyID(ctx context.Context, companyID int) ([]*core.Project, error)
	UpdateProject(ctx context.Context, project *core.Project) (*core.Project, error)
	DeleteProject(ctx context.Context, id int) error
	GetProjectEmployees(ctx context.Context, id int) ([]*core.Employee, error)
	AddEmployeeToProject(ctx context.Context, companyId int, employees []int) error
}

type ProjectService struct {
	log         *slog.Logger
	projectRepo ProjectRepository
}

func NewProjectService(projectRepo ProjectRepository, log *slog.Logger) *ProjectService {
	return &ProjectService{projectRepo: projectRepo, log: log}
}

func (s *ProjectService) CreateProject(ctx context.Context, project *core.Project) (*core.Project, error) {
	return s.projectRepo.CreateProject(ctx, project)
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id int) (*core.Project, error) {
	return s.projectRepo.FindProjectByID(ctx, id)
}

func (s *ProjectService) GetProjectEmployees(ctx context.Context, id int) ([]*core.Employee, error) {
	return s.projectRepo.GetProjectEmployees(ctx, id)
}

func (s *ProjectService) GetProjectsByCompanyID(ctx context.Context, companyID int) ([]*core.Project, error) {
	return s.projectRepo.FindProjectsByCompanyID(ctx, companyID)
}

func (s *ProjectService) UpdateProject(ctx context.Context, project *core.Project) (*core.Project, error) {
	_, err := s.projectRepo.FindProjectByID(ctx, project.ID)
	if err != nil {
		return nil, err
	}

	return s.projectRepo.UpdateProject(ctx, project)
}

func (s *ProjectService) DeleteProject(ctx context.Context, id int) error {
	_, err := s.projectRepo.FindProjectByID(ctx, id)
	if err != nil {
		return err
	}

	return s.projectRepo.DeleteProject(ctx, id)
}

func (s *ProjectService) AddEmployeeToProject(ctx context.Context, companyId int, employees []int) error {
	return s.projectRepo.AddEmployeeToProject(ctx, companyId, employees)
}
