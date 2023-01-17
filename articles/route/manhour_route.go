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

func ManhourRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	manhourService := service.NewManhourService(
		repository.NewManhourRepository(),
		db,
		validate,
	)
	manhourController := controller.NewManhourController(manhourService)

	router.GET("/manhour", auth.Auth(manhourController.FindAll, []string{}))
}
