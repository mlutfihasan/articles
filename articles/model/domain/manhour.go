package domain

import (
	"gitlab.com/VNEU/mf-micro-service-cogs-data/model/web"
)

type Manhours []Manhour

type Manhour struct {
	// Fields
	Periode               string  `gorm:"periode"`
	BatchNo               string  `gorm:"batch_no"`
	ProductID             string  `gorm:"product_id"`
	ProductName           string  `gorm:"product_name"`
	SubtahapName          string  `gorm:"subtahap_name"`
	Subtahap              string  `gorm:"subtahap"`
	Manhour               float32 `gorm:"manhour"`
	JumlahOrang           int     `gorm:"jumlah_orang"`
	ManhourNonJumlahOrang float32 `gorm:"manhour_non_jumlah_orang"`
	Divisi                string  `gorm:"divisi"`
}

func (manhour *Manhour) ToManhourResponse() web.ManhourResponse {
	return web.ManhourResponse{
		// fields
		Periode:               manhour.Periode,
		BatchNo:               manhour.BatchNo,
		ProductID:             manhour.ProductID,
		ProductName:           manhour.ProductName,
		SubtahapName:          manhour.SubtahapName,
		Subtahap:              manhour.Subtahap,
		Manhour:               manhour.Manhour,
		JumlahOrang:           manhour.JumlahOrang,
		ManhourNonJumlahOrang: manhour.ManhourNonJumlahOrang,
		Divisi:                manhour.Divisi,
	}
}

func (manhours Manhours) ToManhourResponses() []web.ManhourResponse {
	manhourResponses := []web.ManhourResponse{}
	for _, manhour := range manhours {
		manhourResponses = append(manhourResponses, manhour.ToManhourResponse())
	}
	return manhourResponses
}
