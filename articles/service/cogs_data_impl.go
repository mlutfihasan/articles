package service

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/repository"
	"gorm.io/gorm"
)

type CogsDataServiceImpl struct {
	CogsDataRepository repository.CogsDataRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCogsDataService(
	cogsData repository.CogsDataRepository,
	db *gorm.DB,
	validate *validator.Validate,
) CogsDataService {
	return &CogsDataServiceImpl{
		CogsDataRepository: cogsData,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CogsDataServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.CogsDataResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	cogsDatas := service.CogsDataRepository.FindAll(tx, filters)
	return cogsDatas.ToCogsDataResponses()
}
