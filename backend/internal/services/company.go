package service

import (
	"context"
	"github.com/Max425/manager/internal/model/core"
	"log/slog"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *core.Company) (*core.Company, error)
	FindCompanyByID(ctx context.Context, id int) (*core.Company, error)
	UpdateCompany(ctx context.Context, id int, company *core.Company) (*core.Company, error)
	DeleteCompany(ctx context.Context, id int) error
}

type CompanyService struct {
	log         *slog.Logger
	companyRepo CompanyRepository
}

func NewCompanyService(companyRepo CompanyRepository, log *slog.Logger) *CompanyService {
	return &CompanyService{companyRepo: companyRepo, log: log}
}

func (s *CompanyService) CreateCompany(ctx context.Context, company *core.Company) (*core.Company, error) {
	return s.companyRepo.CreateCompany(ctx, company)
}

func (s *CompanyService) GetCompanyByID(ctx context.Context, id int) (*core.Company, error) {
	return s.companyRepo.FindCompanyByID(ctx, id)
}

func (s *CompanyService) UpdateCompany(ctx context.Context, id int, company *core.Company) (*core.Company, error) {
	_, err := s.companyRepo.FindCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.companyRepo.UpdateCompany(ctx, id, company)
}

func (s *CompanyService) DeleteCompany(ctx context.Context, id int) error {
	_, err := s.companyRepo.FindCompanyByID(ctx, id)
	if err != nil {
		return err
	}

	return s.companyRepo.DeleteCompany(ctx, id)
}
