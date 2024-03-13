package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	CompanyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		*NewCompanyRepository(db),
	}
}
