package service

import (
	"context"
	"errors"
	"github.com/Max425/manager/internal/httpserver/model"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company model.Company) (model.Company, error)
	FindByID(ctx context.Context, id int) (model.Company, error)
	Update(ctx context.Context, id int, company model.Company) (model.Company, error)
	Delete(ctx context.Context, id int) error
}

type CompanyService struct {
	companyRepo CompanyRepository
}

func NewCompanyService(companyRepo CompanyRepository) *CompanyService {
	return &CompanyService{companyRepo: companyRepo}
}

func (s *CompanyService) CreateCompany(ctx context.Context, company model.Company) (model.Company, error) {
	// Здесь может быть логика валидации
	return s.companyRepo.CreateCompany(ctx, company)
}

func (s *CompanyService) GetCompanyByID(ctx context.Context, id int) (model.Company, error) {
	// Дополнительная логика, если требуется
	return s.companyRepo.FindByID(ctx, id)
}

func (s *CompanyService) UpdateCompany(ctx context.Context, id int, company model.Company) (model.Company, error) {
	// Проверка на существование компании
	_, err := s.companyRepo.FindByID(ctx, id)
	if err != nil {
		// Возвращаем ошибку, если компания не найдена
		return model.Company{}, errors.New("company not found")
	}

	return s.companyRepo.Update(ctx, id, company)
}

func (s *CompanyService) DeleteCompany(ctx context.Context, id int) error {
	// Проверка на существование компании перед удалением
	_, err := s.companyRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("company not found")
	}

	return s.companyRepo.Delete(ctx, id)
}
