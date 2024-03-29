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
	var employee core.Employee
	err := er.db.GetContext(ctx, &employee, "SELECT * FROM employee WHERE id=$1", id)
	if err != nil {
		er.log.Error("Error finding employee", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if employee.ID == 0 {
		return nil, core.ErrNotFound
	}
	return &employee, nil
}

func (er *EmployeeRepository) FindEmployeesByCompanyID(ctx context.Context, companyID int) ([]*core.Employee, error) {
	var employees []*core.Employee
	err := er.db.SelectContext(ctx, &employees, "SELECT * FROM employee WHERE company_id=$1", companyID)
	if err != nil {
		er.log.Error("Error finding employees by company ID", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if len(employees) == 0 {
		return nil, core.ErrNotFound
	}
	return employees, nil
}

func (er *EmployeeRepository) GetEmployerProjects(ctx context.Context, id int) ([]*core.Project, error) {
	var projects []*core.Project
	err := er.db.SelectContext(ctx, &projects, `select * from project 
         							where id in (select project_id from employee_project 
                                        where employee_id = $1)`, id)
	if err != nil {
		er.log.Error("Error finding employees by company ID", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if len(projects) == 0 {
		return nil, core.ErrNotFound
	}
	return projects, nil
}

func (er *EmployeeRepository) UpdateEmployee(ctx context.Context, employee *core.Employee) (*core.Employee, error) {
	_, err := er.db.ExecContext(ctx, "UPDATE employee SET company_id=$1, name=$2, position=$3, mail=$4, password=$5, salt=$6, image=$7, rating=$8 WHERE id=$9",
		employee.CompanyID, employee.Name, employee.Position, employee.Mail, employee.Password, employee.Salt, employee.Image, employee.Rating, employee.ID)
	if err != nil {
		er.log.Error("Error updating employee", slog.String("error", err.Error()))
		return nil, err
	}
	return er.FindEmployeeByID(ctx, employee.ID)
}

func (er *EmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	_, err := er.db.ExecContext(ctx, "DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		er.log.Error("Error deleting employee", slog.String("error", err.Error()))
		return core.ErrInternal
	}
	return nil
}
