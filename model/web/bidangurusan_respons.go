package web

type BidangUrusanRespons struct {
	Id               int    `json:"id"`
	KodeOpd          string `json:"kode_opd"`
	KodeBidangUrusan string `json:"kode_bidang_urusan"`
	BidangUrusan     string `json:"bidang_urusan"`
}
