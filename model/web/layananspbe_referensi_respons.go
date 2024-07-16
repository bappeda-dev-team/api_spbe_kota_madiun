package web

type LayananSpbeReferensiArsitekturRespons struct {
	Id              int
	Kode_referensi  string `json:"kode_referensi"`
	Nama_referensi  string `json:"nama_referensi"`
	Level_referensi int    `json:"level_referensi"`
}
