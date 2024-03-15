package convert

import (
	"fmt"
	"github.com/Max425/manager/internal/model/core"
	"github.com/Max425/manager/internal/model/dto"
	"strings"
)

func ProjectDtoToCore(dtoProject *dto.Project) (*core.Project, error) {
	return &core.Project{
		ID:           dtoProject.ID,
		CompanyID:    dtoProject.CompanyID,
		Name:         dtoProject.Name,
		Stages:       fmt.Sprintf("{%s}", strings.Join(dtoProject.Stages, ",")),
		Image:        dtoProject.Image,
		Description:  dtoProject.Description,
		CurrentStage: dtoProject.CurrentStage,
		Deadline:     dtoProject.Deadline,
		Status:       dtoProject.Status,
		Complexity:   dtoProject.Complexity,
	}, nil
}

func ProjectCoreToDto(coreProject *core.Project) *dto.Project {
	return &dto.Project{
		ID:           coreProject.ID,
		CompanyID:    coreProject.CompanyID,
		Name:         coreProject.Name,
		Stages:       strings.Split(strings.Trim(coreProject.Stages, "{}"), ","),
		Image:        coreProject.Image,
		Description:  coreProject.Description,
		CurrentStage: coreProject.CurrentStage,
		Deadline:     coreProject.Deadline,
		Status:       coreProject.Status,
		Complexity:   coreProject.Complexity,
	}
}
