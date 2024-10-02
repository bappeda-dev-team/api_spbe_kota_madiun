package web

type KeteranganGapUpdateRequest struct {
	Id             int
	IdProsesBisnis int    `json:"id_prosesbisnis"`
	KodeOpd        string `json:"kode_opd"`
	KeteranganGap  string `json:"keterangan_gap"`
}
