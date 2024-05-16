package web

import "time"

type ReferensiArsitekturUpdateRequest struct {
	Id 				int
	Kode_referensi  string `json:"kode_referensi" validate:"required"`
	Nama_referensi  string `json:"nama_referensi" validate:"required"`
	Level_referensi int 	`json:"level_referensi" validate:"required"`
	Jenis_referensi string	`json:"jenis_referensi" validate:"required"`
	Updated_at      time.Time 
}