package repository

import (
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/domain"
	"gorm.io/gorm"
)

type CogsDataRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.CogsDatas
}
