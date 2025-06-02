package service

import (
	"context"
	"github.com/Max425/manager/common/slices"
	"github.com/Max425/manager/internal/model/common"
	"github.com/Max425/manager/internal/model/convert"
	"github.com/Max425/manager/internal/model/core"
	"github.com/Max425/manager/internal/model/dto"
	"sort"
)

func (s *EmployeeService) GetAutoEmployees(
	ctx context.Context,
	autoEmployees dto.AutoEmployees,
	companyID int) ([]dto.AutoEmployee, error) {

	allEmployees, err := s.employeeRepo.FindEmployeesByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.AutoEmployee, 0, len(autoEmployees.AutoEmployee))
	for _, autoEmployee := range autoEmployees.AutoEmployee {
		if autoEmployee.Pin {
			result = append(result, autoEmployee)
		} else {
			result = append(result, chooseEmployee(allEmployees, autoEmployee, autoEmployees.Project))
		}
	}

	return result, nil
}

func chooseEmployee(allEmployees []*core.Employee, autoEmployee dto.AutoEmployee, project dto.Project) dto.AutoEmployee {
	employees := slices.FilterAll(allEmployees,
		func(_ int, v *core.Employee) bool { return v.Position == autoEmployee.Position },
	)
	if len(employees) == 0 || len(employees) == 1 && autoEmployee.Employee.ID != 0 {
		return autoEmployee
	}

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Rating[len(employees[i].Rating)-1] < employees[j].Rating[len(employees[j].Rating)-1]
	})

	if project.Complexity >= common.MinComplexity {
		employees = employees[len(employees)/2:]
	}

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].ActiveProjectsCount < employees[j].ActiveProjectsCount
	})

	index := 0
	if employees[index].ID == autoEmployee.Employee.ID {
		index++
	}

	autoEmployee.Employee = *convert.EmployeeCoreToDto(employees[index])
	return autoEmployee
}
