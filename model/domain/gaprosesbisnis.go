package domain

import "database/sql"

type GapProsesBisnis struct {
	ID               int                   `json:"id"`
	KodeOpd          string                `json:"kode_opd"`
	Tahun            int                   `json:"tahun"`
	NamaProsesBisnis string                `json:"nama_proses_bisnis"`
	KodeProsesBisnis string                `json:"kode_proses_bisnis"`
	Layanan          []GapLayanan          `json:"layanans"`
	DataDanInformasi []GapDataDanInformasi `json:"data_dan_informasi"`
	Aplikasi         []GapAplikasi         `json:"aplikasi"`
	KeteranganGap    []GapKeteranganGap    `json:"keterangan_gap"`
	Keterangan       []GapKeterangan       `json:"keterangan"`
}

type GapLayanan struct {
	NamaLayanan sql.NullString `json:"nama_layanan"`
}

type GapDataDanInformasi struct {
	NamaData sql.NullString `json:"nama_data"`
}

type GapAplikasi struct {
	NamaAplikasi sql.NullString `json:"nama_aplikasi"`
}

type GapKeterangan struct {
	IdKeterangan sql.NullInt32  `json:"id_keterangan"`
	Keterangan   sql.NullString `json:"keterangan"`
}

type GapKeteranganGap struct {
	IdKeteranganGap sql.NullInt32  `json:"id_keterangan_gap"`
	KeteranganGap   sql.NullString `json:"keterangan_gap"`
}
