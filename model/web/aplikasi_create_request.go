package web

type AplikasiCreateRequest struct {
	NamaAplikasi           string `json:"nama_aplikasi"`
	FungsiAplikasi         string `json:"fungsi_aplikasi"`
	JenisAplikasi          string `json:"jenis_aplikasi"`
	ProdusenAplikasi       string `json:"produsen_aplikasi"`
	PjAplikasi             string `json:"pj_aplikasi"`
	KodeOPD                string `json:"kode_opd"`
	InformasiTerkaitInput  string `json:"informasi_terkait_input"`
	InformasiTerkaitOutput string `json:"informasi_terkait_output"`
	Interoprabilitas       string `json:"interoprabilitas"`
	Tahun                  int    `json:"tahun"`
	RaaLevel1id            *int   `json:"raa_level_1_id"`
	RaaLevel2id            *int   `json:"raa_level_2_id"`
	RaaLevel3id            *int   `json:"raa_level_3_id"`
	StrategicId            *int   `json:"strategic_id"`
	TacticalId             *int   `json:"tactical_id"`
	OperationalId          *int   `json:"operational_id"`
}
