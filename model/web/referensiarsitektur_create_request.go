package web

import "time"

type ReferensiArsitekturCreateRequest struct {
	Kode_referensi  string
	Nama_referensi  string
	Level_referensi string
	Jenis_referensi string
	Created_at      time.Time
}