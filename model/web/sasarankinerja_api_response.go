package web

type ApiResponsSasaranKinerjaPegawai struct {
	Results []SasaranKinerjaPegawaiApi `json:"results"`
}

type SasaranKinerjaPegawaiApi struct {
	Id                 int    `json:"id"`
	KodeOPD            string `json:"kode_opd"`
	KodeSasaranPegawai string `json:"kode_sasaran"`
	SasaranPegawai     string `json:"sasaran"`
	AnggaranSasaran    string `json:"anggaran_sasaran"`
	PelaksanaSasaran   string `json:"pelaksana_sasaran"`
	TahunSasaran       string `json:"tahun_sasaran"`
	KodeSubKegiatan    string `json:"kode_subkegiatan_sasaran"`
	SubKegiatan        string `json:"subkegiatan_sasaran"`
}
