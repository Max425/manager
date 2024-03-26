package http_server

import (
	_ "github.com/Max425/manager/docs"
	"github.com/Max425/manager/internal/config"
	"github.com/Max425/manager/internal/http-server/handler"
	"github.com/Max425/manager/internal/repository"
	"github.com/Max425/manager/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
)

type Service interface {
	handler.CompanyService
}

func NewHttpServer(log *slog.Logger, postgres config.PostgresConfig, listenAddr string) (*http.Server, error) {
	// connect to db
	dbConnect, err := repository.NewPostgresDB(postgres)
	if err != nil {
		return nil, err
	}

	// create all repositories
	managerRepo := repository.NewRepository(dbConnect, log)

	// create all services
	managerService := service.NewService(managerRepo, log)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	companyHandler := handler.NewCompanyHandler(log, managerService)
	employeeHandler := handler.NewEmployeeHandler(log, managerService)
	projectHandler := handler.NewProjectHandler(log, managerService)

	api := router.Group("/api")
	{
		company := api.Group("/companies")
		{
			company.POST("/", companyHandler.CreateCompany)
			company.GET("/:id", companyHandler.GetCompany)
			company.GET("/employees", employeeHandler.GetEmployeesByCompanyID)
			company.GET("/projects", projectHandler.GetProjectsByCompanyID)
			company.PUT("/", companyHandler.UpdateCompany)
			company.DELETE("/:id", companyHandler.DeleteCompany)
		}
		employees := api.Group("/employees")
		{
			employees.POST("/", employeeHandler.CreateEmployee)
			employees.GET("/:id", employeeHandler.GetEmployee)
			employees.PUT("/", employeeHandler.UpdateEmployee)
			employees.DELETE("/:id", employeeHandler.DeleteEmployee)
		}
		projects := api.Group("/projects")
		{
			projects.POST("/", projectHandler.CreateProject)
			projects.GET("/:id", projectHandler.GetProject)
			projects.PUT("/", projectHandler.UpdateProject)
			projects.DELETE("/:id", projectHandler.DeleteProject)
		}
	}

	return &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}, nil
}
