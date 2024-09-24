package web

type BIDANGURUSANOPD struct {
	Id               int    `json:"id_bidang_urusan"`
	KodeBidangUrusan string `json:"kode_bidang_urusan"`
	BidangUrusan     string `json:"bidang_urusan"`
}

type URUSANOPD struct {
	Id              int               `json:"id_urusan"`
	KodeUrusan      string            `json:"kode_urusan"`
	Urusan          string            `json:"urusan"`
	BidangUrusanOPD []BIDANGURUSANOPD `json:"bidang_urusan"`
}

type OPD struct {
	KodeOpd   string      `json:"kode_opd"`
	NamaOpd   string      `json:"nama_opd"`
	UrusanOPD []URUSANOPD `json:"urusan_opd"`
}
