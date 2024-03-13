package httpserver

import (
	"github.com/Max425/manager/internal/config"
	"github.com/Max425/manager/internal/httpserver/handler"
	"github.com/Max425/manager/internal/httpserver/repository"
	service "github.com/Max425/manager/internal/httpserver/services"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
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
	managerRepo := repository.NewRepository(dbConnect)

	// create all services
	managerService := service.NewService(managerRepo)

	router := gin.New()
	gin.Default()
	router.Use(sloggin.New(log), gin.Recovery())

	companyHandler := handler.NewCompanyHandler(log, managerService)

	api := router.Group("/api")
	{
		//	lists := api.Group("/lists")
		//	{
		//		lists.POST("/", h.createList)
		//		lists.GET("/", h.getAllLists)
		//		lists.GET("/:id", h.getListById)
		//		lists.PUT("/:id", h.updateList)
		//		lists.DELETE("/:id", h.deleteList)
		//
		//		items := lists.Group(":id/items")
		//		{
		//			items.POST("/", h.createItem)
		//			items.GET("/", h.getAllItems)
		//		}
		//	}
		//
		company := api.Group("/companies")
		{
			company.POST("/", companyHandler.CreateCompany)
			company.GET("/:id", companyHandler.GetCompany)
			company.PUT("/:id", companyHandler.UpdateCompany)
			company.DELETE("/:id", companyHandler.DeleteCompany)
		}
	}

	return &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}, nil
}
