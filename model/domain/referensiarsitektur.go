package domain

import "time"

type ReferensiArsitektur struct {
	IdReferensi     int
	Kode_referensi  string
	Nama_referensi  string
	Level_referensi int
	Jenis_referensi string
	Created_at      time.Time
	Updated_at 		time.Time
	Tahun 			int
}
