package http_server

import (
	"github.com/Max425/manager/internal/config"
	"github.com/Max425/manager/internal/http-server/handler"
	"github.com/Max425/manager/internal/repository"
	"github.com/Max425/manager/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"

	_ "github.com/Max425/manager/docs"
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
