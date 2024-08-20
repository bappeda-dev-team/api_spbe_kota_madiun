package web

type DataDanInformasiCreateRequest struct {
	NamaData               string  `json:"nama_data" validate:"required"`
	SifatData              string  `json:"sifat_data" validate:"required"`
	JenisData              string  `json:"jenis_data" validate:"required"`
	ProdusenData           string  `json:"produsen_data" validate:"required"`
	UraianData             string  `json:"uraian_data" validate:"required"`
	ValiditasData          string  `json:"validitas_data" validate:"required"`
	PjData                 string  `json:"pj_data" validate:"required"`
	KodeOPD                string  `json:"kode_opd"`
	InformasiTerkaitInput  string  `json:"informasi_terkait_input" validate:"required"`
	InformasiTerkaitOutput string  `json:"informasi_terkait_output" validate:"required"`
	Interoprabilitas       string  `json:"interoprabilitas" validate:"required"`
	Keterangan             *string `json:"keterangan"`
	Tahun                  int     `json:"tahun" validate:"required"`
	RadLevel1id            *int    `json:"rad_level_1_id"`
	RadLevel2id            *int    `json:"rad_level_2_id"`
	RadLevel3id            *int    `json:"rad_level_3_id"`
	RadLevel4id            *int    `json:"rad_level_4_id"`
	StrategicId            *int    `json:"strategic_id"`
	TacticalId             *int    `json:"tactical_id"`
	OperationalId          *int    `json:"operational_id"`
}
