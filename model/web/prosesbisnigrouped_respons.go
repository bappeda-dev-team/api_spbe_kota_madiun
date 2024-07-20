package web

type ProsesBisnisGroupedResponse struct {
	Data []ProsesBisnisData `json:"data"`
}

type ProsesBisnisData struct {
	ProsesBisnis []ProsesBisnisResponse `json:"proses_bisnis"`
}

type ProsesBisnisResponse struct {
	NamaProsesBisnis string             `json:"nama_proses_bisnis"`
	KodeProsesBisnis string             `json:"kode_proses_bisnis"`
	KodeOpd          string             `json:"kode_opd"`
	Tahun            int                `json:"tahun"`
	Layanans         []Layanan          `json:"layanans"`
	DataDanInformasi []DataDanInformasi `json:"data_dan_informasi"`
	Aplikasi         []Aplikasi         `json:"aplikasi"`
}

type Layanan struct {
	NamaLayanan string `json:"nama_layanan"`
}

type DataDanInformasi struct {
	NamaData string `json:"nama_data"`
}

type Aplikasi struct {
	NamaAplikasi string `json:"nama_aplikasi"`
}
