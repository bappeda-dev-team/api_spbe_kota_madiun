package web

import "time"

type PohonKinerjaRespons struct {
	ID         int    `json:"id"`
	Parent     string `json:"parent"`
	NamaPohon  string `json:"nama_pohon"`
	JenisPohon string `json:"jenis_pohon"`
	LevelPohon int    `json:"level_pohon"`
	KodeOpd    string `json:"kode_opd"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Tahun      string `json:"tahun"`
}
