package web

type BidangUrusanOPD struct {
	KodeBidangUrusan string `json:"kode_bidang_urusan"`
	BidangUrusan     string `json:"bidang_urusan"`
}

type UrusanOPD struct {
	KodeUrusan      string            `json:"kode_urusan"`
	Urusan          string            `json:"urusan"`
	BidangUrusanOPD []BidangUrusanOPD `json:"bidang_urusan_opd"`
}

type Opd struct {
	KodeOpd   string      `json:"kode_opd"`
	NamaOpd   string      `json:"nama_opd"`
	UrusanOPD []UrusanOPD `json:"urusan_opd"`
}

type ResponsesAPI struct {
	Results []Opd `json:"results"`
}
