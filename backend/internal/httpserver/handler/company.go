package handler

import (
	"context"
	"github.com/Max425/manager/internal/httpserver/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, company model.Company) (model.Company, error)
	GetCompanyByID(ctx context.Context, id int) (model.Company, error)
	UpdateCompany(ctx context.Context, id int, company model.Company) (model.Company, error)
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

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var company model.Company
	if err := c.BindJSON(&company); err != nil {
		h.log.Error("Error binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.companyService.CreateCompany(c.Request.Context(), company)
	if err != nil {
		h.log.Error("Error creating company", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	company, err := h.companyService.GetCompanyByID(c.Request.Context(), id)
	if err != nil {
		h.log.Error("Error getting company", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var company model.Company
	if err = c.BindJSON(&company); err != nil {
		h.log.Error("Error binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.companyService.UpdateCompany(c.Request.Context(), id, company)
	if err != nil {
		h.log.Error("Error updating company", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Error converting id", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = h.companyService.DeleteCompany(c.Request.Context(), id)
	if err != nil {
		h.log.Error("Error deleting company", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
