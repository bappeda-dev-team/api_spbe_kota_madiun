package domain

import "time"

type PohonKinerja struct {
	ID         int
	Parent     string
	JenisPohon string
	LevelPohon int
	NamaPohon  string
	Tahun      string
	KodeOpd    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
