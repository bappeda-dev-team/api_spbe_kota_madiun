package web

type KebutuhanSPBEUpdateRequest struct {
	ID             int                    `validate:"required" json:"id"`
	KeteranganGap  string                 `validate:"required" json:"keterangan_gap"`
	KodeOpd        string                 `validate:"required" json:"kode_opd"`
	Tahun          int                    `validate:"required" json:"tahun"`
	NamaDomain     string                 `validate:"required" json:"nama_domain"`
	IdProsesbisnis int                    `validate:"required" json:"id_prosesbisnis"`
	JenisKebutuhan []JenisKebutuhanUpdate `validate:"required" json:"jenis_kebutuhan"`
}

type JenisKebutuhanUpdate struct {
	ID               int                 `json:"id,omitempty"`
	JenisKebutuhanId int                 `json:"kebutuhan_id"`
	Kebutuhan        string              `json:"kebutuhan"`
	KondisiAwal      []KondisiAwalUpdate `json:"kondisi_awal"`
}

type KondisiAwalUpdate struct {
	ID               int    `json:"id,omitempty"`
	JenisKebutuhanId int    `json:"jenis_kebutuhan_id"`
	Keterangan       string `json:"keterangan"`
	Tahun            int    `json:"tahun"`
}
