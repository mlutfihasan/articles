package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
)

type CogsDataController interface {
	FindAll(c *gin.Context, auth *auth.AccessDetails)
}
