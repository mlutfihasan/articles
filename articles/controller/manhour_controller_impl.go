package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/service"
)

type ManhourControllerImpl struct {
	ManhourService service.ManhourService
}

func NewManhourController(leaveAllowanceService service.ManhourService) ManhourController {
	return &ManhourControllerImpl{
		ManhourService: leaveAllowanceService,
	}
}

func (controller *ManhourControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c)
	manhourResponses := controller.ManhourService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(manhourResponses),
		Data:    manhourResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}
