package domain

type RencanaPelaksanaanPegawai struct {
	Id               int                `json:"id"`
	KodeOpd          string             `json:"kode_opd"`
	IdKebutuhan      int                `json:"id_kebutuhan"`
	IdSasaranKinerja int                `json:"id_sasaran_kinerja"`
	TahunPelaksanaan []TahunPelaksanaan `json:"tahun_pelaksanaan"`
}

type TahunPelaksanaan struct {
	Id                 int `json:"id_tahun_pelaksanaan"`
	IdRencanaPelaksana int `json:"id_rencana_pelaksana"`
	Tahun              int `json:"tahun"`
}
