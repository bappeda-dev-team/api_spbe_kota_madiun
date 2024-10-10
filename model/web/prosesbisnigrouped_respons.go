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
	KeteranganGap    []GapKeteranganGap    `json:"keterangan_gap"`
	Keterangan       []GapKeterangan       `json:"keterangan_kebutuhan"`
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
	IdKeterangan NullInt32  `json:"id_keterangan"`
	Keterangan   NullString `json:"keterangan"`
}

type GapKeteranganGap struct {
	IdKeteranganGap NullInt32  `json:"id_keterangan_gap"`
	KeteranganGap   NullString `json:"keterangan_gap"`
}

type NullString struct {
	sql.NullString
}

type NullInt32 struct {
	sql.NullInt32
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

func (ni NullInt32) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int32)
	}
	return json.Marshal(nil)
}

func (ni *NullInt32) UnmarshalJSON(b []byte) error {
	var i *int32
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Int32 = *i
		ni.Valid = true
	} else {
		ni.Valid = false
	}
	return nil
}
