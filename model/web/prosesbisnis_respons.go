package web

import "time"

type ProsesBinsisRespons struct {
	Id                 int
	Nama_proses_bisnis string
	Sasaran_kota       string
	Kode_proses_bisnis string
	Bidang_urusan      string
	Kode_opd           string
	Rad_level_1        string
	Rad_level_2        string
	Rad_level_3        string
	Rad_level_4        string
	Rad_level_5        string
	Rad_level_6        string
	Created_at         time.Time
	Updated_at         time.Time
}