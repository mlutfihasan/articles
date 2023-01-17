package service

import (
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
)

type ManhourService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.ManhourResponse
}
