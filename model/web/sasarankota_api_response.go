package web

type SasaranKotaApi struct {
	ID           int    `json:"id"`
	TujuanKota   string `json:"tujuan_kota"`
	StrategiKota string `json:"strategi_kota"`
	Sasaran      string `json:"sasaran"`
}

type DataSasaranKota struct {
	Tahun         string           `json:"tahun"`
	Jumlahsasaran int              `json:"jumlah_sasaran"`
	SasaranKota   []SasaranKotaApi `json:"sasaran_kota_collections"`
}

type APIResponseDataSasaranKota struct {
	Data DataSasaranKota `json:"data"`
}
