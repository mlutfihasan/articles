package repository

import (
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/domain"

	"gorm.io/gorm"
)

type ManhourRepositoryImpl struct {
}

func NewManhourRepository() ManhourRepository {
	return &ManhourRepositoryImpl{}
}

func (repository *ManhourRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Manhours {
	Manhours := domain.Manhours{}
	tx := db.Table("MF_WH.manhour").
		Select("periode, batch_no, product_id, product_name, subtahap_name, subtahap, manhour, jumlah_orang, manhour_non_jumlah_orang, divisi").Find(&Manhours)

	helper.PanicIfError(tx.Error)

	return Manhours
}
