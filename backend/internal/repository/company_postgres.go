package repository

import (
	"context"
	"github.com/Max425/manager/internal/model/core"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type CompanyRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewCompanyRepository(db *sqlx.DB, log *slog.Logger) *CompanyRepository {
	return &CompanyRepository{db: db, log: log}
}

func (cr *CompanyRepository) CreateCompany(ctx context.Context, company *core.Company) (*core.Company, error) {
	query := "INSERT INTO company (name, positions, image, description) VALUES ($1, $2, $3, $4) RETURNING id"
	row := cr.db.QueryRowContext(ctx, query, company.Name, company.Positions, company.Image, company.Description)
	var id int
	err := row.Scan(&id)
	if err != nil {
		cr.log.Error("Error create company", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	company.ID = id
	return company, nil
}

func (cr *CompanyRepository) FindCompanyByID(ctx context.Context, id int) (*core.Company, error) {
	var company core.Company
	err := cr.db.GetContext(ctx, &company, "SELECT * FROM company WHERE id=$1", id)
	if err != nil {
		cr.log.Error("Error find company", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if company.ID == 0 {
		return nil, core.ErrNotFound
	}

	return &company, nil
}

func (cr *CompanyRepository) UpdateCompany(ctx context.Context, company *core.Company) (*core.Company, error) {
	_, err := cr.db.ExecContext(ctx, "UPDATE company SET name=$1, positions=$2, image=$3, description=$4 WHERE id=$5",
		company.Name, company.Positions, company.Image, company.Description, company.ID)
	if err != nil {
		cr.log.Error("Error updating company", slog.String("error", err.Error()))
		return nil, err
	}
	return cr.FindCompanyByID(ctx, company.ID)
}

func (cr *CompanyRepository) DeleteCompany(ctx context.Context, id int) error {
	_, err := cr.db.ExecContext(ctx, "DELETE FROM company WHERE id=$1", id)
	if err != nil {
		cr.log.Error("Error delete company", slog.String("error", err.Error()))
		return core.ErrInternal
	}
	return nil
}
