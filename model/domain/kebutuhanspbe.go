package domain

import "database/sql"

type KebutuhanSPBE struct {
	ID              int              `json:"id"`
	Keterangan      string           `json:"keterangan"`
	KodeOpd         string           `json:"kode_opd"`
	Tahun           int              `json:"tahun"`
	NamaDomain      sql.NullString   `json:"nama_domain"`
	IdProsesbisnis  int              `json:"id_prosesbisnis"`
	JenisKebutuhan  []JenisKebutuhan `json:"jenis_kebutuhan"`
	IndikatorPj     sql.NullString
	PenanggungJawab sql.NullString
}

type JenisKebutuhan struct {
	Id          int           `json:"id"`
	KebutuhanId int           `json:"kebutuhan_id"`
	Kebutuhan   string        `json:"kebutuhan"`
	KondisiAwal []KondisiAwal `json:"kondisi_awal"`
}

type KondisiAwal struct {
	Id               int    `json:"id"`
	JenisKebutuhanId int    `json:"jenis_kebutuhan_id"`
	Keterangan       string `json:"keterangan"`
	Tahun            int    `json:"tahun"`
}
