package web

import (
	"time"
)

type ProsesBisnisCreateRequest struct {
	NamaProsesBisnis string `json:"nama_proses_bisnis"`
	SasaranKotaId    *int   `json:"sasaran_kota_id"`
	KodeOPD          string `json:"kode_opd"`
	BidangUrusanId   *int   `json:"bidang_urusan_id"`
	RabLevel1ID      *int   `json:"rab_level_1_id"`
	RabLevel2ID      *int   `json:"rab_level_2_id"`
	RabLevel3ID      *int   `json:"rab_level_3_id"`
	RabLevel4ID      *int   `json:"rab_level_4_id"`
	RabLevel5ID      *int   `json:"rab_level_5_id"`
	RabLevel6ID      *int   `json:"rab_level_6_id"`
	Tahun            int    `json:"tahun"`
	CreatedAt        time.Time
}
