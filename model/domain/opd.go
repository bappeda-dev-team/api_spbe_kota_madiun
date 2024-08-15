package domain

import "time"

type Opd struct {
	Id        string
	KodeOpd   string `json:"kode_opd"`
	NamaOpd   string `json:"nama_opd"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
