package domain

import "time"

type PohonKinerja struct {
	ID         int
	JenisPohon string
	LevelPohon int
	NamaPohon  string
	Tahun      string
	KodeOpd    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Data struct {
	Tahun         string            `json:"tahun"`
	KodeOpd       string            `json:"kode_opd"`
	NamaOpd       string            `json:"nama_opd"`
	PohonKinerjas []PohonKinerja `json:"pohon_kinerjas"`
}

type Results struct {
	Data Data `json:"data"`
}

type APIResponse struct {
	Results Results `json:"results"`
}
