package web

import "time"

type SasaranKotaRespons struct {
	ID           int
	TujuanKota   string
	StrategiKota string
	Sasaran      string
	Tahun        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
