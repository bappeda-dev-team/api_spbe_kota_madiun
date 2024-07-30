package web

type DataDanInformasiRespons struct {
	Id                     int
	NamaData               string
	SifatData              string
	JenisData              string
	ProdusenData           string
	UraianData             string
	ValiditasData          string
	PjData                 string
	KodeOPD                string
	InformasiTerkaitInput  string
	InformasiTerkaitOutput string
	Interoprabilitas       string
	Tahun                  int
	CreatedAt              string
	UpdatedAt              string
	RadLevel1id            *DataDanInformasiReferensiArsitekturRespons
	RadLevel2id            *DataDanInformasiReferensiArsitekturRespons
	RadLevel3id            *DataDanInformasiReferensiArsitekturRespons
	RadLevel4id            *DataDanInformasiReferensiArsitekturRespons
	StrategicId            *DataDanInformasiPohonResponns
	TacticalId             *DataDanInformasiPohonResponns
	OperationalId          *DataDanInformasiPohonResponns
}
