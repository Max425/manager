package convert

import (
	"fmt"
	"github.com/Max425/manager/internal/model/core"
	"github.com/Max425/manager/internal/model/dto"
	"strings"
)

func CompanyDtoToCore(dtoCompany *dto.Company) (*core.Company, error) {
	return core.NewCompany(
		dtoCompany.ID,
		dtoCompany.Name,
		fmt.Sprintf("{%s}", strings.Join(dtoCompany.Positions, ",")),
		dtoCompany.Image,
		dtoCompany.Description,
	)
}

func CompanyCoreToDto(coreCompany *core.Company) *dto.Company {
	return &dto.Company{
		ID:          coreCompany.ID,
		Name:        coreCompany.Name,
		Positions:   strings.Split(strings.Trim(coreCompany.Positions, "{}"), ","),
		Image:       coreCompany.Image,
		Description: coreCompany.Description,
	}
}
