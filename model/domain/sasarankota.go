package domain

import "time"

type SasaranKotaApi struct {
	ID           int
	TujuanKota   string
	StrategiKota string
	Sasaran      string
	Tahun        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
