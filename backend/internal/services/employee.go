package service

import (
	"context"
	"github.com/Max425/manager/internal/model/core"
	"log/slog"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *core.Employee) (*core.Employee, error)
	FindEmployeeByID(ctx context.Context, id int) (*core.Employee, error)
	UpdateEmployee(ctx context.Context, id int, employee *core.Employee) (*core.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type EmployeeService struct {
	log          *slog.Logger
	employeeRepo EmployeeRepository
}

func NewEmployeeService(employeeRepo EmployeeRepository, log *slog.Logger) *EmployeeService {
	return &EmployeeService{employeeRepo: employeeRepo, log: log}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee *core.Employee) (*core.Employee, error) {
	return s.employeeRepo.CreateEmployee(ctx, employee)
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (*core.Employee, error) {
	return s.employeeRepo.FindEmployeeByID(ctx, id)
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, id int, employee *core.Employee) (*core.Employee, error) {
	_, err := s.employeeRepo.FindEmployeeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.employeeRepo.UpdateEmployee(ctx, id, employee)
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	_, err := s.employeeRepo.FindEmployeeByID(ctx, id)
	if err != nil {
		return err
	}

	return s.employeeRepo.DeleteEmployee(ctx, id)
}