package web

import "time"

type SasaranKotaRespons struct {
	ID           int
	Sasaran      string
	TujuanKota   string
	StrategiKota string
	Tahun        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
