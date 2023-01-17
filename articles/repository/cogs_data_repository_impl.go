package repository

import (
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/domain"

	"gorm.io/gorm"
)

type CogsDataRepositoryImpl struct {
}

func NewCogsDataRepository() CogsDataRepository {
	return &CogsDataRepositoryImpl{}
}

func (repository *CogsDataRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.CogsDatas {
	CogsDatas := domain.CogsDatas{}
	tx := db.Table("MF_WH.material").
		Select("bulan, tahun, tanggal, po_no, material_id, material_name, satuan, supplier, principal, po_harga, ppn, ppn_nilai, pph, pph_nilai, delivery_date, dp_persen, dp, divisi_nama, qty_terima, tanggal_terima, qty_rilis, release_date, na").Find(&CogsDatas)

	helper.PanicIfError(tx.Error)

	return CogsDatas
}
