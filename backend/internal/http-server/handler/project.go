package handler

import (
	"context"
	"errors"
	"github.com/Max425/manager/internal/model/common"
	"github.com/Max425/manager/internal/model/convert"
	"github.com/Max425/manager/internal/model/core"
	"github.com/Max425/manager/internal/model/dto"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type ProjectService interface {
	CreateProject(ctx context.Context, project *core.Project) (*core.Project, error)
	GetProjectByID(ctx context.Context, id int) (*core.Project, error)
	GetProjectsByCompanyID(ctx context.Context, companyID int) ([]*core.Project, error)
	UpdateProject(ctx context.Context, project *core.Project) (*core.Project, error)
	DeleteProject(ctx context.Context, id int) error
	GetProjectEmployees(ctx context.Context, id int) ([]*core.Employee, error)
	AddEmployeeToProject(ctx context.Context, companyId int, employees []int) error
}

type ProjectHandler struct {
	log            *slog.Logger
	projectService ProjectService
}

func NewProjectHandler(log *slog.Logger, projectService ProjectService) *ProjectHandler {
	return &ProjectHandler{
		log:            log,
		projectService: projectService,
	}
}

// CreateProject создает новый проект.
// @Summary Создает новый проект
// @Description Создает новый проект с заданными данными.
// @Tags Project
// @Accept json
// @Produce json
// @Param project body dto.Project true "Данные проекта"
// @Success 200 {object} dto.Project "Успешно создан проект"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project dto.Project
	if err := c.BindJSON(&project); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreProject, err := convert.ProjectDtoToCore(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coreProject.CompanyID = 1 //TODO: fix
	result, err := h.projectService.CreateProject(c.Request.Context(), coreProject)
	if err != nil {
		h.log.Error("Error creating project", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.ProjectCoreToDto(result))
}

// GetProject возвращает информацию о проекте по его ID.
// @Summary Возвращает информацию о проекте по ID
// @Description Возвращает информацию о проекте по указанному ID.
// @Tags Project
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Success 200 {object} dto.Project "Успешно получен проект"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/projects/{id} [get]
func (h *ProjectHandler) GetProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	project, err := h.projectService.GetProjectByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error getting project", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.ProjectCoreToDto(project))
}

// GetProjectEmployees возвращает сотрудников проекта по ID.
// @Summary Возвращает сотрудников проекта по ID
// @Description Возвращает сотрудников проекта по ID.
// @Tags Project
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Success 200 {object} []dto.Employee "Успешно получены сотрудники"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/projects/{id}/employees [get]
func (h *ProjectHandler) GetProjectEmployees(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	employees, err := h.projectService.GetProjectEmployees(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error getting project", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	var dtoEmployees []*dto.Employee
	for _, employee := range employees {
		dtoEmployees = append(dtoEmployees, convert.EmployeeCoreToDto(employee))
	}

	c.JSON(http.StatusOK, dtoEmployees)
}

// GetProjectsByCompanyID возвращает все проекты компании по ее ID из контекста.
// @Summary Возвращает все проекты компании по ID компании из контекста
// @Description Возвращает все проекты компании по ID компании из контекста.
// @Tags Company
// @Accept json
// @Produce json
// @Success 200 {array} dto.Project "Успешно получены проекты компании"
// @Failure 404 {object} string "Компания не найдена"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/companies/projects [get]
func (h *ProjectHandler) GetProjectsByCompanyID(c *gin.Context) {
	companyID := 1 //c.Value("company_id").(int) //TODO: fix

	projects, err := h.projectService.GetProjectsByCompanyID(c.Request.Context(), companyID)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error getting projects by company ID", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	var dtoProjects []*dto.Project
	for _, project := range projects {
		dtoProjects = append(dtoProjects, convert.ProjectCoreToDto(project))
	}

	c.JSON(http.StatusOK, dtoProjects)
}

// UpdateProject обновляет информацию о проекте.
// @Summary Обновляет информацию о проекте
// @Description Обновляет информацию о проекте с указанным ID новыми данными.
// @Tags Project
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Param project body dto.Project true "Новые данные проекта"
// @Success 200 {object} dto.Project "Успешно обновлен проект"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/projects [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var project dto.Project
	if err := c.BindJSON(&project); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreProject, err := convert.ProjectDtoToCore(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.projectService.UpdateProject(c.Request.Context(), coreProject)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error update project", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.ProjectCoreToDto(result))
}

// DeleteProject удаляет проект по его ID.
// @Summary Удаляет проект по ID
// @Description Удаляет проект с указанным ID.
// @Tags Project
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Success 200 {object} string "Успешное удаление проекта"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	err = h.projectService.DeleteProject(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error deleting project", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}

// AddEmployeeToProject добавляет сотрудника в проект.
// @Summary Добавляет сотрудника в проект
// @Description Добавляет сотрудника в проект.
// @Tags Project
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Param project body []int true "ID сотрудников"
// @Success 200 {object} string "Всё хорошо"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/projects/{id}/employees [post]
func (h *ProjectHandler) AddEmployeeToProject(c *gin.Context) {
	var employees []int
	if err := c.BindJSON(&employees); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	err = h.projectService.AddEmployeeToProject(c.Request.Context(), id, employees)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error getting project", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}
	c.JSON(http.StatusOK, "Employee add")
}
