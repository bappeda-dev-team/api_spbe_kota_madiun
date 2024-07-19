package web

type DataDanInformasiUpdateRequest struct {
	Id                     int
	NamaData               string `json:"nama_data"`
	SifatData              string `json:"sifat_data"`
	JenisData              string `json:"jenis_data"`
	ProdusenData           string `json:"produsen_data"`
	PjData                 string `json:"pj_data"`
	KodeOPD                string `json:"kode_opd"`
	InformasiTerkaitInput  string `json:"informasi_terkait_input"`
	InformasiTerkaitOutput string `json:"informasi_terkait_output"`
	Interoprabilitas       string `json:"interoprabilitas"`
	Tahun                  int    `json:"tahun"`
	RadLevel1id            int    `json:"rad_level_1_id"`
	RadLevel2id            int    `json:"rad_level_2_id"`
	RadLevel3id            int    `json:"rad_level_3_id"`
	RadLevel4id            int    `json:"rad_level_4_id"`
	StrategicId            int    `json:"strategic_id"`
	TacticalId             int    `json:"tactical_id"`
	OperationalId          int    `json:"operational_id"`
}
