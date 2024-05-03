package web

import "time"

type ReferensiArsitekturUpdateRequest struct {
	Id 				int
	Kode_referensi  string
	Nama_referensi  string
	Level_referensi string
	Jenis_referensi string
	Updated_at      time.Time
}