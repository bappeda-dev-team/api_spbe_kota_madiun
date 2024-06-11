package web

type PohonKinerjaApi struct {
	ID         int    `json:"id"`
	JenisPohon string `json:"jenis_pohon"`
	LevelPohon int    `json:"level_pohon"`
	Strategi   string `json:"strategi"`
}

type Data struct {
	Tahun         string            `json:"tahun"`
	KodeOpd       string            `json:"kode_opd"`
	NamaOpd       string            `json:"nama_opd"`
	PohonKinerjas []PohonKinerjaApi `json:"pohon_kinerjas"`
}

type Results struct {
	Data Data `json:"data"`
}

type APIResponse struct {
	Results Results `json:"results"`
}
