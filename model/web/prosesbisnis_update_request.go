package web

import "time"

type ProsesBisnisUpdateRequest struct {
	Id               int
	NamaProsesBisnis string `json:"nama_proses_bisnis"`
	SasaranKota      string `json:"sasaran_kota"`
	KodeProsesBisnis string `json:"kode_proses_bisnis"`
	KodeOPD          string `json:"kode_opd"`
	BidangUrusan     string `json:"bidang_urusan"`
	RabLevel1ID      int    `json:"rab_level_1_id"`
	RabLevel2ID      int    `json:"rab_level_2_id"`
	RabLevel3ID      int    `json:"rab_level_3_id"`
	RabLevel4ID      int    `json:"rab_level_4_id"`
	RabLevel5ID      int    `json:"rab_level_5_id"`
	RabLevel6ID      int    `json:"rab_level_6_id"`
	Tahun            int    `json:"tahun"`
	UpdatedAt        time.Time
}
