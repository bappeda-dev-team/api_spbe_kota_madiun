package web

import "time"

type ReferensiArsitekturUpdateRequest struct {
	Id              int
	Kode_referensi  string `validate:"required"`
	Nama_referensi  string `validate:"required"`
	Level_referensi int    `validate:"required"`
	Jenis_referensi string `validate:"required"`
	Updated_at      time.Time
}
