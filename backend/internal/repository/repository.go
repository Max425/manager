package repository

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Repository struct {
	CompanyRepository
	ProjectRepository
	EmployeeRepository
	SessionRepository
}

func NewRepository(db *sqlx.DB, log *slog.Logger) *Repository {
	return &Repository{
		*NewCompanyRepository(db, log),
		*NewProjectRepository(db, log),
		*NewEmployeeRepository(db, log),
		*NewSessionRepository(db, log),
	}
}
