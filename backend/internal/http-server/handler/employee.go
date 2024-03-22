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

type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee *core.Employee) (*core.Employee, error)
	GetEmployeeByID(ctx context.Context, id int) (*core.Employee, error)
	UpdateEmployee(ctx context.Context, employee *core.Employee) (*core.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type EmployeeHandler struct {
	log             *slog.Logger
	employeeService EmployeeService
}

func NewEmployeeHandler(log *slog.Logger, employeeService EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		log:             log,
		employeeService: employeeService,
	}
}

// CreateEmployee создает нового сотрудника.
// @Summary Создает нового сотрудника
// @Description Создает нового сотрудника с заданными данными.
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee body dto.Employee true "Данные сотрудника"
// @Success 200 {object} dto.Employee "Успешно создан сотрудник"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/employees [post]
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee dto.Employee
	if err := c.BindJSON(&employee); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreEmployee, err := convert.EmployeeDtoToCore(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.employeeService.CreateEmployee(c.Request.Context(), coreEmployee)
	if err != nil {
		h.log.Error("Error creating employee", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.EmployeeCoreToDto(result))
}

// GetEmployee возвращает информацию о сотруднике по его ID.
// @Summary Возвращает информацию о сотруднике по ID
// @Description Возвращает информацию о сотруднике по указанному ID.
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} dto.Employee "Успешно получен сотрудник"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/employees/{id} [get]
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	employee, err := h.employeeService.GetEmployeeByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error getting employee", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.EmployeeCoreToDto(employee))
}

// UpdateEmployee обновляет информацию о сотруднике.
// @Summary Обновляет информацию о сотруднике
// @Description Обновляет информацию о сотруднике с указанным ID новыми данными.
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Param employee body dto.Employee true "Новые данные сотрудника"
// @Success 200 {object} dto.Employee "Успешно обновлен сотрудник"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/employees [put]
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	var employee dto.Employee
	if err := c.BindJSON(&employee); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreEmployee, err := convert.EmployeeDtoToCore(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.employeeService.UpdateEmployee(c.Request.Context(), coreEmployee)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error update employee", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.EmployeeCoreToDto(result))
}

// DeleteEmployee удаляет сотрудника по его ID.
// @Summary Удаляет сотрудника по ID
// @Description Удаляет сотрудника с указанным ID.
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} string "Успешное удаление сотрудника"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/employees/{id} [delete]
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	err = h.employeeService.DeleteEmployee(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error deleting employee", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
