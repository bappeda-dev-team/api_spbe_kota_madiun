package web

import "time"

type ProsesBisnisCreateRequest struct {
	NamaProsesBisnis string `json:"nama_proses_bisnis" validate:"required"`
	SasaranKota      string `json:"sasaran_kota" validate:"required"`
	KodeProsesBisnis string `json:"kode_proses_bisnis" validate:"required"`
	KodeOPD          string `json:"kode_opd" validate:"required"`
	BidangUrusan     string `json:"bidang_urusan" validate:"required"`
	RabLevel1ID      int    `json:"rab_level_1_id"`
	RabLevel2ID      int    `json:"rab_level_2_id"`
	RabLevel3ID      int    `json:"rab_level_3_id"`
	RabLevel4ID      int    `json:"rab_level_4_id"`
	RabLevel5ID      int    `json:"rab_level_5_id"`
	RabLevel6ID      int    `json:"rab_level_6_id"`
	Tahun            int    `json:"tahun"`
	CreatedAt        time.Time
}
