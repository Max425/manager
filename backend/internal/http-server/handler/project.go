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
	UpdateProject(ctx context.Context, id int, project *core.Project) (*core.Project, error)
	DeleteProject(ctx context.Context, id int) error
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
// @Router /api/projects/{id} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	var project dto.Project
	if err = c.BindJSON(&project); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreProject, err := convert.ProjectDtoToCore(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.projectService.UpdateProject(c.Request.Context(), id, coreProject)
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
