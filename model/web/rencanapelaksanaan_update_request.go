package web

type RencanaPelaksanaanUpdateRequest struct {
	Id               int                             `json:"id"`
	KodeOpd          string                          `json:"kode_opd"`
	IdKebutuhan      int                             `json:"id_kebutuhan"`
	IdSasaranKinerja int                             `json:"id_sasaran_kinerja"`
	TahunPelaksanaan []TahunPelaksanaanUpdateRequest `json:"tahun_pelaksanaan"`
}

type TahunPelaksanaanUpdateRequest struct {
	IdRencanaPelaksana int `json:"id_rencana_pelaksana"`
	Tahun              int `json:"tahun"`
}
