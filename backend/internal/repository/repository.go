package repository

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Repository struct {
	CompanyRepository
}

func NewRepository(db *sqlx.DB, log *slog.Logger) *Repository {
	return &Repository{
		*NewCompanyRepository(db, log),
	}
}
