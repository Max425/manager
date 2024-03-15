package repository

import (
	"context"
	"github.com/Max425/manager/internal/model/core"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type EmployeeRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewEmployeeRepository(db *sqlx.DB, log *slog.Logger) *EmployeeRepository {
	return &EmployeeRepository{db: db, log: log}
}

func (er *EmployeeRepository) CreateEmployee(ctx context.Context, employee *core.Employee) (*core.Employee, error) {
	query := "INSERT INTO employee (company_id, name, position, mail, password, salt, image, rating) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	row := er.db.QueryRowContext(ctx, query, employee.CompanyID, employee.Name, employee.Position, employee.Mail, employee.Password, employee.Salt, employee.Image, employee.Rating)
	var id int
	err := row.Scan(&id)
	if err != nil {
		er.log.Error("Error creating employee", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	employee.ID = id
	return employee, nil
}

func (er *EmployeeRepository) FindEmployeeByID(ctx context.Context, id int) (*core.Employee, error) {
	var employee *core.Employee
	err := er.db.GetContext(ctx, employee, "SELECT * FROM employee WHERE id=$1", id)
	if employee == nil {
		return nil, core.ErrNotFound
	}
	if err != nil {
		er.log.Error("Error finding employee", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	return employee, nil
}

func (er *EmployeeRepository) UpdateEmployee(ctx context.Context, id int, employee *core.Employee) (*core.Employee, error) {
	_, err := er.db.ExecContext(ctx, "UPDATE employee SET company_id=$1, name=$2, position=$3, mail=$4, password=$5, salt=$6, image=$7, rating=$8 WHERE id=$9",
		employee.CompanyID, employee.Name, employee.Position, employee.Mail, employee.Password, employee.Salt, employee.Image, employee.Rating, id)
	if err != nil {
		er.log.Error("Error updating employee", slog.String("error", err.Error()))
		return nil, err
	}
	return er.FindEmployeeByID(ctx, id)
}

func (er *EmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	_, err := er.db.ExecContext(ctx, "DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		er.log.Error("Error deleting employee", slog.String("error", err.Error()))
		return core.ErrInternal
	}
	return nil
}
