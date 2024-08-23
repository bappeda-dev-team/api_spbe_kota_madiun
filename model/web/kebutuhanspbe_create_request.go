package web

type KebutuhanSPBECreateRequest struct {
	KeteranganGap  string                 `json:"keterangan_gap"`
	KodeOpd        string                 `json:"kode_opd"`
	Tahun          int                    `json:"tahun"`
	NamaDomain     string                 `json:"nama_domain"`
	IdProsesbisnis int                    `json:"id_prosesbisnis"`
	JenisKebutuhan []JenisKebutuhanCreate `json:"jenis_kebutuhan"`
}

type JenisKebutuhanCreate struct {
	Kebutuhan   string              `json:"kebutuhan"`
	KondisiAwal []KondisiAwalCreate `json:"kondisi_awal"`
}

type KondisiAwalCreate struct {
	Keterangan string `json:"keterangan"`
	Tahun      int    `json:"tahun"`
}
