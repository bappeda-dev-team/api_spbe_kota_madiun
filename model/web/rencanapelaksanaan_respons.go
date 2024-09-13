package web

type RencanaPelaksanaanResponse struct {
	Id               int                           `json:"id"`
	KodeOpd          string                        `json:"kode_opd"`
	IdKebutuhan      int                           `json:"id_kebutuhan"`
	SasaranKinerja   SasaranKinerjaPegawaiResponse `json:"sasaran_kinerja"`
	IndikatorPD      string                        `json:"indikator_pd"`
	PerangkatDaerah  string                        `json:"perangkat_daerah"`
	TahunPelaksanaan []TahunPelaksanaanResponse    `json:"tahun_pelaksanaan"`
}

type TahunPelaksanaanResponse struct {
	Id    int `json:"id_tahun_pelaksanaan"`
	Tahun int `json:"tahun"`
}
