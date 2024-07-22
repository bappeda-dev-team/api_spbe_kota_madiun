package web

type AplikasiPohonRespons struct {
	ID         int    `json:"id"`
	NamaPohon  string `json:"nama_pohon"`
	LevelPohon int    `json:"level_pohon"`
}
