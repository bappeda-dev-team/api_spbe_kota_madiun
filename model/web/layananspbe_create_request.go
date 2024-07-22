package web

import "time"

type LayananSpbeCreateRequest struct {
	NamaLayanan       string `json:"nama_layanan"`
	KodeLayanan       string `json:"kode_layanan"`
	TujuanLayananId   *int   `json:"tujuan_layanan_id"`
	FungsiLayanan     string `json:"fungsi_layanan"`
	Tahun             int    `json:"tahun"`
	KodeOPD           string `json:"kode_opd"`
	KementrianTerkait string `json:"kementrian_terkait"`
	MetodeLayanan     string `json:"metode_layanan"`
	CreatedAt         time.Time
	RalLevel1id       *int `json:"ral_level_1_id"`
	RalLevel2id       *int `json:"ral_level_2_id"`
	RalLevel3id       *int `json:"ral_level_3_id"`
	RalLevel4id       *int `json:"ral_level_4_id"`
	StrategicId       *int `json:"strategic_id"`
	TacticalId        *int `json:"tactical_id"`
	OperationalId     *int `json:"operational_id"`
}
