package service

import "log/slog"

type Repository interface {
	CompanyRepository
}

type Service struct {
	CompanyService
}

func NewService(repo Repository, log *slog.Logger) *Service {
	return &Service{
		*NewCompanyService(repo, log),
	}
}
