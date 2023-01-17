package domain

import (
	"time"

	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
)

type CogsDatas []CogsData

type CogsData struct {
	// Fields
	Bulan         int       `gorm:"column:bulan"`
	Tahun         int       `gorm:"column:tahun"`
	Tanggal       int       `gorm:"column:tanggal"`
	PoNo          string    `gorm:"column:po_no"`
	MaterialID    string    `gorm:"column:material_id"`
	MaterialName  string    `gorm:"column:material_name"`
	Satuan        string    `gorm:"column:satuan"`
	Supplier      string    `gorm:"column:supplier"`
	Principal     string    `gorm:"column:principal"`
	PoHarga       float32   `gorm:"column:po_harga"`
	Ppn           float32   `gorm:"column:ppn"`
	PpnNilai      float32   `gorm:"column:ppn_nilai"`
	Pph           float32   `gorm:"column:pph"`
	Pph_Nilai     float32   `gorm:"column:pph_nilai"`
	DeliveryDate  time.Time `gorm:"column:delivery_date"`
	DpPersen      float32   `gorm:"column:dp_persen"`
	Dp            float32   `gorm:"column:dp"`
	DivisiNama    string    `gorm:"column:divisi_nama"`
	QtyTerima     float32   `gorm:"column:qty_terima"`
	TanggalTerima time.Time `gorm:"column:tanggal_terima"`
	QtyRilis      float32   `gorm:"column:qty_rilis"`
	ReleaseDate   time.Time `gorm:"column:release_date"`
	Na            string    `gorm:"column:na"`
}

func (cogsData *CogsData) ToCogsDataResponse() web.CogsDataResponse {
	return web.CogsDataResponse{
		// fields
		Bulan:         cogsData.Bulan,
		Tahun:         cogsData.Tahun,
		Tanggal:       cogsData.Tahun,
		PoNo:          cogsData.PoNo,
		MaterialID:    cogsData.MaterialID,
		MaterialName:  cogsData.MaterialName,
		Satuan:        cogsData.Satuan,
		Supplier:      cogsData.Supplier,
		Principal:     cogsData.Principal,
		PoHarga:       cogsData.PoHarga,
		Ppn:           cogsData.Ppn,
		PpnNilai:      cogsData.PpnNilai,
		Pph:           cogsData.Pph,
		Pph_Nilai:     cogsData.Pph_Nilai,
		DeliveryDate:  cogsData.DeliveryDate,
		DpPersen:      cogsData.DpPersen,
		Dp:            cogsData.Dp,
		DivisiNama:    cogsData.DivisiNama,
		QtyTerima:     cogsData.QtyTerima,
		TanggalTerima: cogsData.TanggalTerima,
		QtyRilis:      cogsData.QtyRilis,
		ReleaseDate:   cogsData.ReleaseDate,
		Na:            cogsData.Na,
	}
}

func (cogsDatas CogsDatas) ToCogsDataResponses() []web.CogsDataResponse {
	cogsDataResponses := []web.CogsDataResponse{}
	for _, cogsData := range cogsDatas {
		cogsDataResponses = append(cogsDataResponses, cogsData.ToCogsDataResponse())
	}
	return cogsDataResponses
}
