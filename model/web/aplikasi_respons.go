package web

type AplikasiRespons struct {
	Id                     int
	NamaAplikasi           string
	FungsiAplikasi         string
	JenisAplikasi          string
	ProdusenAplikasi       string
	PjAplikasi             string
	KodeOPD                string
	InformasiTerkaitInput  string
	InformasiTerkaitOutput string
	Interoprabilitas       string
	Tahun                  int
	CreatedAt              string
	UpdatedAt              string
	RaaLevel1id            *AplikasiReferensiArsitekturRespons
	RaaLevel2id            *AplikasiReferensiArsitekturRespons
	RaaLevel3id            *AplikasiReferensiArsitekturRespons
	RaaLevel4id            *AplikasiReferensiArsitekturRespons
	StrategicId            *AplikasiPohonRespons
	TacticalId             *AplikasiPohonRespons
	OperationalId          *AplikasiPohonRespons
}
