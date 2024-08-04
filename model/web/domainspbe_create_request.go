package web

type DomainSPBECreateRequest struct {
	NamaDomain string `json:"nama_domain"`
	KodeDomain string `json:"kode_domain"`
	Tahun      int    `json:"tahun"`
}
