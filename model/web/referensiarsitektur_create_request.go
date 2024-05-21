package web

import "time"

type ReferensiArsitekturCreateRequest struct {
	Kode_referensi  string    `json:"kode_referensi" validate:"required,uppercase"`
	Nama_referensi  string    `json:"nama_referensi" validate:"required"`
	Level_referensi int       `json:"level_referensi" validate:"required,numeric"`
	Jenis_referensi string	  `json:"jenis_referensi" validate:"required"`
	Created_at      time.Time
}