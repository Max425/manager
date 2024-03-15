package service

import "log/slog"

type Repository interface {
	CompanyRepository
	EmployeeRepository
	ProjectRepository
}

type Service struct {
	CompanyService
	EmployeeService
	ProjectService
}

func NewService(repo Repository, log *slog.Logger) *Service {
	return &Service{
		*NewCompanyService(repo, log),
		*NewEmployeeService(repo, log),
		*NewProjectService(repo, log),
	}
}
