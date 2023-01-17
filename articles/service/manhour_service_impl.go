package service

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/auth"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/repository"
	"gorm.io/gorm"
)

type ManhourServiceImpl struct {
	ManhourRepository repository.ManhourRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewManhourService(
	manhour repository.ManhourRepository,
	db *gorm.DB,
	validate *validator.Validate,
) ManhourService {
	return &ManhourServiceImpl{
		ManhourRepository: manhour,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ManhourServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.ManhourResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	manhours := service.ManhourRepository.FindAll(tx, filters)
	return manhours.ToManhourResponses()
}
