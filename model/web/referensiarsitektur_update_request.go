package web

import "time"

type ReferensiArsitekturUpdateRequest struct {
	Id 				int
	Kode_referensi  string    `json:"kode_referensi" validate:"required,uppercase"`
	Nama_referensi  string    `json:"nama_referensi" validate:"required"`
	Level_referensi int       `json:"level_referensi" validate:"required,numeric"`
	Jenis_referensi string	  `json:"jenis_referensi" validate:"required"`
	Updated_at      time.Time `json:"updated_at"`
	Tahun 			int 	  `json:"tahun" validate:"required"`
}
