package domain

import (
	"database/sql"
)

type Petarencana struct {
	ID               int                       `json:"id"`
	GeneratedID      string                    `json:"generated_id"`
	KodeOpd          string                    `json:"kode_opd"`
	Tahun            int                       `json:"tahun"`
	NamaProsesBisnis string                    `json:"nama_proses_bisnis"`
	KodeProsesBisnis string                    `json:"kode_proses_bisnis"`
	Layanan          []RencanaLayanan          `json:"layanans"`
	DataDanInformasi []RencanaDataDanInformasi `json:"data_dan_informasi"`
	Aplikasi         []RencanaAplikasi         `json:"aplikasi"`
	Keterangan       []RencanaKeterangan       `json:"keterangan"`
}

type RencanaLayanan struct {
	NamaLayanan sql.NullString `json:"nama_layanan"`
}

type RencanaDataDanInformasi struct {
	NamaData sql.NullString `json:"nama_data"`
}

type RencanaAplikasi struct {
	NamaAplikasi sql.NullString `json:"nama_aplikasi"`
}

type RencanaKeterangan struct {
	IdKeterangan sql.NullInt32 `json:"id_keterangan"`
}
