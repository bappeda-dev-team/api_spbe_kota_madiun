package web

import "time"

type ProsesBinsisCreateRequest struct {
	Id                 int
	Nama_proses_bisnis string `json:"nama_proses_bisnis"`
	Sasaran_kota       string `json:"sasaran_kota"`
	Kode_proses_bisnis string `json:"kode_proses_bisnis"`
	Bidang_urusan      string `json:"bidang_urusan"`
	Kode_opd           string `json:"kode_opd"`
	KodeReferensi      string `json:"kode_referensi"`
	Tahun				int `json:"tahun"`
	Created_at         time.Time

}