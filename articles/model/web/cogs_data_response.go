package web

import "time"

type CogsDataResponse struct {
	Bulan         int       `json:"bulan"`
	Tahun         int       `json:"tahun"`
	Tanggal       int       `json:"tanggal"`
	PoNo          string    `json:"po_no"`
	MaterialID    string    `json:"material_id"`
	MaterialName  string    `json:"material_name"`
	Satuan        string    `json:"satuan"`
	Supplier      string    `json:"supplier"`
	Principal     string    `json:"principal"`
	PoHarga       float32   `json:"po_harga"`
	Ppn           float32   `json:"ppn"`
	PpnNilai      float32   `json:"ppn_nilai"`
	Pph           float32   `json:"pph"`
	Pph_Nilai     float32   `json:"pph_nilai"`
	DeliveryDate  time.Time `json:"delivery_date"`
	DpPersen      float32   `json:"dp_persen"`
	Dp            float32   `json:"dp"`
	DivisiNama    string    `json:"divisi_nama"`
	QtyTerima     float32   `json:"qty_terima"`
	TanggalTerima time.Time `json:"tanggal_terima"`
	QtyRilis      float32   `json:"qty_rilis"`
	ReleaseDate   time.Time `json:"release_date"`
	Na            string    `json:"na"`
}
