package convert

import (
	"github.com/Max425/manager/internal/model/core"
	"github.com/Max425/manager/internal/model/dto"
)

func EmployeeDtoToCore(dtoEmployee *dto.Employee) (*core.Employee, error) {
	return &core.Employee{
		ID:                   dtoEmployee.ID,
		CompanyID:            dtoEmployee.CompanyID,
		Name:                 dtoEmployee.Name,
		Position:             dtoEmployee.Position,
		Mail:                 dtoEmployee.Mail,
		Password:             dtoEmployee.Password,
		Salt:                 dtoEmployee.Salt,
		Image:                dtoEmployee.Image,
		Rating:               dtoEmployee.Rating,
		ActiveProjectsCount:  dtoEmployee.ActiveProjectsCount,
		OverdueProjectsCount: dtoEmployee.OverdueProjectsCount,
		TotalProjectsCount:   dtoEmployee.TotalProjectsCount,
	}, nil
}

func EmployeeCoreToDto(coreEmployee *core.Employee) *dto.Employee {
	return &dto.Employee{
		ID:                   coreEmployee.ID,
		CompanyID:            coreEmployee.CompanyID,
		Name:                 coreEmployee.Name,
		Position:             coreEmployee.Position,
		Mail:                 coreEmployee.Mail,
		Password:             coreEmployee.Password,
		Salt:                 coreEmployee.Salt,
		Image:                coreEmployee.Image,
		Rating:               coreEmployee.Rating,
		ActiveProjectsCount:  coreEmployee.ActiveProjectsCount,
		OverdueProjectsCount: coreEmployee.OverdueProjectsCount,
		TotalProjectsCount:   coreEmployee.TotalProjectsCount,
	}
}
