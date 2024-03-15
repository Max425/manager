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

type CompanyService interface {
	CreateCompany(ctx context.Context, company *core.Company) (*core.Company, error)
	GetCompanyByID(ctx context.Context, id int) (*core.Company, error)
	UpdateCompany(ctx context.Context, id int, company *core.Company) (*core.Company, error)
	DeleteCompany(ctx context.Context, id int) error
}

type CompanyHandler struct {
	log            *slog.Logger
	companyService CompanyService
}

func NewCompanyHandler(log *slog.Logger, companyService CompanyService) *CompanyHandler {
	return &CompanyHandler{
		log:            log,
		companyService: companyService,
	}
}

// CreateCompany создает новую компанию.
// @Summary Создает новую компанию
// @Description Создает новую компанию с заданными данными.
// @Tags Company
// @Accept json
// @Produce json
// @Param company body core.Company true "Данные компании"
// @Success 200 {object} core.Company "Успешно создана компания"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/companies [post]
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var company dto.Company
	if err := c.BindJSON(&company); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreCompany, err := convert.CompanyDtoToCore(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.companyService.CreateCompany(c.Request.Context(), coreCompany)
	if err != nil {
		h.log.Error("Error creating company", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.CompanyCoreToDto(result))
}

// GetCompany возвращает информацию о компании по ее ID.
// @Summary Возвращает информацию о компании по ID
// @Description Возвращает информацию о компании по указанному ID.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "ID компании"
// @Success 200 {object} core.Company "Успешно получена компания"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/companies/{id} [get]
func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	company, err := h.companyService.GetCompanyByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error getting company", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.CompanyCoreToDto(company))
}

// UpdateCompany обновляет информацию о компании.
// @Summary Обновляет информацию о компании
// @Description Обновляет информацию о компании с указанным ID новыми данными.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "ID компании"
// @Param company body core.Company true "Новые данные компании"
// @Success 200 {object} core.Company "Успешно обновлена компания"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/companies/{id} [put]
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	var company dto.Company
	if err = c.BindJSON(&company); err != nil {
		h.log.Error("Error binding JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	coreCompany, err := convert.CompanyDtoToCore(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.companyService.UpdateCompany(c.Request.Context(), id, coreCompany)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error update company", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, convert.CompanyCoreToDto(result))
}

// DeleteCompany удаляет компанию по ее ID.
// @Summary Удаляет компанию по ID
// @Description Удаляет компанию с указанным ID.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "ID компании"
// @Success 200 {object} string "Успешное удаление компании"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/companies/{id} [delete]
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrBadRequest.String()})
		return
	}

	err = h.companyService.DeleteCompany(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": common.ErrNotFound.String()})
			return
		}
		h.log.Error("Error deleting company", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
