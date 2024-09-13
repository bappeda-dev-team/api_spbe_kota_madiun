package web

type SasaranKinerjaPegawaiResponse struct {
	Id                    int    `json:"id"`
	KodeOpd               string `json:"kode_opd"`
	KodeSasaran           string `json:"kode_sasaran"`
	Tahun                 string `json:"tahun_sasaran"`
	SasaranKinerjaPegawai string `json:"sasaran_kinerja"`
	AnggaranSasaran       string `json:"anggaran_sasaran"`
	PelaksanaSasaran      string `json:"pelaksana_sasaran"`
	KodeSubKegiatan       string `json:"kode_subkegiatan_sasaran"`
	SubKegiatan           string `json:"subkegiatan_sasaran"`
}
