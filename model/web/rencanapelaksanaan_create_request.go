package web

type RencanaPelaksanaanCreateRequest struct {
	KodeOpd          string                          `json:"kode_opd"`
	IdKebutuhan      int                             `json:"id_kebutuhan"`
	IdSasaranKinerja int                             `json:"id_sasaran_kinerja"`
	TahunPelaksanaan []TahunPelaksanaanCreateRequest `json:"tahun_pelaksanaan"`
}

type TahunPelaksanaanCreateRequest struct {
	Tahun int `json:"tahun"`
}
