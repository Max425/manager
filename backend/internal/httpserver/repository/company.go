package repository

import (
	"context"
	"github.com/Max425/manager/internal/httpserver/model"
	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	db *sqlx.DB
}

func NewCompanyRepository(db *sqlx.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (cr *CompanyRepository) CreateCompany(ctx context.Context, company model.Company) (model.Company, error) {
	//positionsArray := "{" + strings.Join(company.Positions, ",") + "}"

	query := "INSERT INTO company (name, positions, image, description) VALUES ($1, $2, $3, $4) RETURNING id"
	row := cr.db.QueryRowContext(ctx, query, company.Name, company.Positions, company.Image, company.Description)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return model.Company{}, err
	}
	company.ID = id
	return company, nil
}

func (cr *CompanyRepository) FindByID(ctx context.Context, id int) (model.Company, error) {
	var company model.Company
	err := cr.db.GetContext(ctx, &company, "SELECT * FROM company WHERE id=$1", id)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}

func (cr *CompanyRepository) Update(ctx context.Context, id int, company model.Company) (model.Company, error) {
	_, err := cr.db.ExecContext(ctx, "UPDATE company SET name=$1, positions=$2, image=$3, description=$4 WHERE id=$5",
		company.Name, company.Positions, company.Image, company.Description, id)
	if err != nil {
		return model.Company{}, err
	}
	updatedCompany, err := cr.FindByID(ctx, id)
	if err != nil {
		return model.Company{}, err
	}
	return updatedCompany, nil
}

func (cr *CompanyRepository) Delete(ctx context.Context, id int) error {
	_, err := cr.db.ExecContext(ctx, "DELETE FROM company WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
