package web

type KeteranganGapRespons struct {
	Id             int    `json:"id_keterangangap"`
	IdProsesBisnis int    `json:"id_prosesbisnis"`
	KodeOpd        string `json:"kode_opd"`
	KeteranganGap  string `json:"keterangan_gap"`
}
