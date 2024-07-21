package web

type LayananSpbeRespons struct {
	Id                int
	NamaLayanan       string
	KodeLayanan       string
	TujuanLayananId   *LayananspbePohonRespons
	FungsiLayanan     string
	Tahun             int
	KodeOPD           string
	KementrianTerkait string
	MetodeLayanan     string
	CreatedAt         string
	UpdatedAt         string
	RalLevel1id       *LayananSpbeReferensiArsitekturRespons
	RalLevel2id       *LayananSpbeReferensiArsitekturRespons
	RalLevel3id       *LayananSpbeReferensiArsitekturRespons
	RalLevel4id       *LayananSpbeReferensiArsitekturRespons
	StrategicId       *LayananspbePohonRespons
	TacticalId        *LayananspbePohonRespons
	OperationalId     *LayananspbePohonRespons
}
