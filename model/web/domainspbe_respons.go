package web

type DomainSPBEResponse struct {
	Id         int    `json:"id"`
	NamaDomain string `json:"nama_domain"`
	KodeDomain string `json:"kode_domain"`
	Tahun      int    `json:"tahun"`
}
