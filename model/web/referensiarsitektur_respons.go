package web

import "time"

type ReferensiArsitekturResponse struct {
	Id              int
	Kode_referensi  string    `json:"kode_referensi"`
	Nama_referensi  string    `json:"nama_referensi"`
	Level_referensi int       `json:"level_referensi"`
	Jenis_referensi string    `json:"jenis_referensi"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
	Tahun           int       `json:"tahun"`
}
