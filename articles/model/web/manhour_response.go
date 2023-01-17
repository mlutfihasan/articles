package web

type ManhourResponse struct {
	Periode               string  `json:"periode"`
	BatchNo               string  `json:"batch_no"`
	ProductID             string  `json:"product_id"`
	ProductName           string  `json:"product_name"`
	SubtahapName          string  `json:"subtahap_name"`
	Subtahap              string  `json:"subtahap"`
	Manhour               float32 `json:"manhour"`
	JumlahOrang           int     `json:"jumlah_orang"`
	ManhourNonJumlahOrang float32 `json:"manhour_non_jumlah_orang"`
	Divisi                string  `json:"divisi"`
}
