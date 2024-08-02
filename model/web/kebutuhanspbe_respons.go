package web

type KebutuhanSPBEResponse struct {
	ID             int                      `json:"id"`
	KodeOpd        string                   `json:"kode_opd"`
	Tahun          int                      `json:"tahun"`
	NamaDomain     string                   `json:"nama_domain"`
	IdProsesbisnis int                      `json:"id_prosesbisnis"`
	JenisKebutuhan []JenisKebutuhanResponse `json:"jenis_kebutuhan"`
}

type JenisKebutuhanResponse struct {
	Id          int                   `json:"id"`
	KebutuhanId int                   `json:"kebutuhan_id"`
	Kebutuhan   string                `json:"kebutuhan"`
	KondisiAwal []KondisiAwalResponse `json:"kondisi_awal"`
}

type KondisiAwalResponse struct {
	Id               int    `json:"id"`
	JenisKebutuhanId int    `json:"jenis_kebutuhan_id"`
	Keterangan       string `json:"keterangan"`
	Tahun            int    `json:"tahun"`
}
