package service

type Repository interface {
	CompanyRepository
}

type Service struct {
	CompanyService
}

func NewService(repo Repository) *Service {
	return &Service{
		*NewCompanyService(repo),
	}
}
