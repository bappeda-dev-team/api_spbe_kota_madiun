package web

type ProsesBisnisRespons struct {
	ID               int                                `json:"id"`
	KodeOPD          string                             `json:"kode_opd"`
	NamaProsesBisnis string                             `json:"nama_proses_bisnis"`
	SasaranKota      *ProsbisSasaranKotaRespons         `json:"sasaran_kota"`
	KodeProsesBisnis string                             `json:"kode_proses_bisnis"`
	BidangUrusan     *ProsBisBidangUrusanRespons        `json:"bidang_urusan_id"`
	RabLevel1        *ProsBisReferensiArsitekturRespons `json:"rab_level_1"`
	RabLevel2        *ProsBisReferensiArsitekturRespons `json:"rab_level_2"`
	RabLevel3        *ProsBisReferensiArsitekturRespons `json:"rab_level_3"`
	RabLevel4        *ProsBisPohonKinerjaRespons        `json:"rab_level_4"`
	RabLevel5        *ProsBisPohonKinerjaRespons        `json:"rab_level_5"`
	RabLevel6        *ProsBisPohonKinerjaRespons        `json:"rab_level_6"`
	Tahun            int                                `json:"tahun"`
	CreatedAt        string                             `json:"created_at"`
	UpdatedAt        string                             `json:"updated_at"`
}
