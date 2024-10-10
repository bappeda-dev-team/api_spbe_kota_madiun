package web

type PetarencanaRespons struct {
	GeneratedID      string                    `json:"generated_id"`
	ID               int                       `json:"id"`
	KodeOpd          string                    `json:"kode_opd"`
	Tahun            int                       `json:"tahun"`
	NamaProsesBisnis string                    `json:"nama_proses_bisnis"`
	KodeProsesBisnis string                    `json:"kode_proses_bisnis"`
	Layanan          []RencanaLayanan          `json:"layanans"`
	DataDanInformasi []RencanaDataDanInformasi `json:"data_dan_informasi"`
	Aplikasi         []RencanaAplikasi         `json:"aplikasi"`
	Keterangan       []PjKebutuhanSPBEResponse `json:"kebutuhan_spbe"`
}

type RencanaLayanan struct {
	NamaLayanan NullString `json:"nama_layanan"`
}

type RencanaDataDanInformasi struct {
	NamaData NullString `json:"nama_data"`
}

type RencanaAplikasi struct {
	NamaAplikasi NullString `json:"nama_aplikasi"`
}
