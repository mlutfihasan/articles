package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/controller"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/repository"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/service"
	"gorm.io/gorm"
)

func CogsDataRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	cogsDataService := service.NewCogsDataService(
		repository.NewCogsDataRepository(),
		db,
		validate,
	)
	cogsDataController := controller.NewCogsDataController(cogsDataService)

	router.GET("/cogs/data", auth.Auth(cogsDataController.FindAll, []string{}))
}
