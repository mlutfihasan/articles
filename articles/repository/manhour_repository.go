package repository

import (
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/domain"
	"gorm.io/gorm"
)

type ManhourRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Manhours
}
