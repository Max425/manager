package convert

import (
	"github.com/Max425/manager/internal/model/core"
	"github.com/Max425/manager/internal/model/dto"
)

func EmployeeDtoToCore(dtoEmployee *dto.Employee) *core.Employee {
	return &core.Employee{
		ID:        dtoEmployee.ID,
		CompanyID: dtoEmployee.CompanyID,
		Name:      dtoEmployee.Name,
		Position:  dtoEmployee.Position,
		Mail:      dtoEmployee.Mail,
		Password:  dtoEmployee.Password,
		Salt:      dtoEmployee.Salt,
		Image:     dtoEmployee.Image,
		Rating:    dtoEmployee.Rating,
	}
}

func EmployeeCoreToDto(coreEmployee *core.Employee) *dto.Employee {
	return &dto.Employee{
		ID:        coreEmployee.ID,
		CompanyID: coreEmployee.CompanyID,
		Name:      coreEmployee.Name,
		Position:  coreEmployee.Position,
		Mail:      coreEmployee.Mail,
		Password:  coreEmployee.Password,
		Salt:      coreEmployee.Salt,
		Image:     coreEmployee.Image,
		Rating:    coreEmployee.Rating,
	}
}
