package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/service"
)

type CogsDataControllerImpl struct {
	CogsDataService service.CogsDataService
}

func NewCogsDataController(leaveAllowanceService service.CogsDataService) CogsDataController {
	return &CogsDataControllerImpl{
		CogsDataService: leaveAllowanceService,
	}
}

func (controller *CogsDataControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c)
	departmentResponses := controller.CogsDataService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(departmentResponses),
		Data:    departmentResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}
