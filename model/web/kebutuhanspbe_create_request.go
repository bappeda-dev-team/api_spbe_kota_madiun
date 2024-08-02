package web

type KebutuhanSPBECreateRequest struct {
	KodeOpd        string                 `validate:"required" json:"kode_opd"`
	Tahun          int                    `validate:"required" json:"tahun"`
	NamaDomain     string                 `validate:"required" json:"nama_domain"`
	IdProsesbisnis int                    `validate:"required" json:"id_prosesbisnis"`
	JenisKebutuhan []JenisKebutuhanCreate `validate:"required" json:"jenis_kebutuhan"`
}

type JenisKebutuhanCreate struct {
	Kebutuhan   string              `json:"kebutuhan"`
	KondisiAwal []KondisiAwalCreate `json:"kondisi_awal"`
}

type KondisiAwalCreate struct {
	Keterangan string `json:"keterangan"`
	Tahun      int    `json:"tahun"`
}
