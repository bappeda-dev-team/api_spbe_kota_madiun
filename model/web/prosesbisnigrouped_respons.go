package web

import (
	"database/sql"
	"encoding/json"
)

type GapProsesBisnis struct {
	ID               int                   `json:"id"`
	KodeOpd          string                `json:"kode_opd"`
	Tahun            int                   `json:"tahun"`
	NamaProsesBisnis string                `json:"nama_proses_bisnis"`
	KodeProsesBisnis string                `json:"kode_proses_bisnis"`
	Layanans         []GapLayanan          `json:"layanans"`
	DataDanInformasi []GapDataDanInformasi `json:"data_dan_informasi"`
	Aplikasi         []GapAplikasi         `json:"aplikasi"`
	Keterangan       []GapKeterangan       `json:"keterangan"`
}

type GapLayanan struct {
	NamaLayanan NullString `json:"nama_layanan"`
}

type GapDataDanInformasi struct {
	NamaData NullString `json:"nama_data"`
}

type GapAplikasi struct {
	NamaAplikasi NullString `json:"nama_aplikasi"`
}

type GapKeterangan struct {
	Keterangan NullString `json:"keterangan"`
}

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	ns.String = s
	ns.Valid = true
	return nil
}
