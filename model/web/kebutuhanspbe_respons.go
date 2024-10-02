package web

type KebutuhanSPBEResponse struct {
	ID              int                      `json:"id"`
	KeteranganGap   string                   `json:"keterangan"`
	KodeOpd         string                   `json:"kode_opd"`
	PerangkatDaerah OpdRespons               `json:"perangkat_daerah"`
	Tahun           int                      `json:"tahun"`
	NamaDomain      string                   `json:"nama_domain"`
	ProsesBisnis    ProsesBisnisResponse     `json:"proses_bisnis"`
	JenisKebutuhan  []JenisKebutuhanResponse `json:"jenis_kebutuhan"`
	IndikatorPj     string                   `json:"indikator_pj"`
	PenanggungJawab string                   `json:"penanggung_jawab"`
}

type PjKebutuhanSPBEResponse struct {
	ID                 int                          `json:"id"`
	KeteranganGap      string                       `json:"keterangan"`
	KodeOpd            string                       `json:"kode_opd"`
	PerangkatDaerah    OpdRespons                   `json:"perangkat_daerah"`
	Tahun              int                          `json:"tahun"`
	NamaDomain         string                       `json:"nama_domain"`
	ProsesBisnis       ProsesBisnisResponse         `json:"proses_bisnis"`
	JenisKebutuhan     []JenisKebutuhanResponse     `json:"jenis_kebutuhan"`
	IndikatorPj        string                       `json:"indikator_pj"`
	PenanggungJawab    OpdRespons                   `json:"penanggung_jawab"`
	RencanaPelaksanaan []RencanaPelaksanaanResponse `json:"rencana_pelaksanaan"`
}

type ProsesBisnisResponse struct {
	ID               int    `json:"id"`
	NamaProsesBisnis string `json:"nama_proses_bisnis"`
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

type KebutuhanSPBEKeteranganResponse struct {
	ID            int    `json:"id"`
	KeteranganGap string `json:"keterangan"`
}

type KebutuhanSPBEPjResponse struct {
	ID              int    `json:"id"`
	IndikatorPj     string `json:"indikator_pj"`
	PenanggungJawab string `json:"penanggung_jawab"`
}
