package web

type KeteranganGapCreateRequest struct {
	IdProsesBisnis int    `json:"id_prosesbisnis"`
	KodeOpd        string `json:"kode_opd"`
	KeteranganGap  string `json:"keterangan_gap"`
}
