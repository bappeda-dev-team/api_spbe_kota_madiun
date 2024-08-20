package domain

import (
	"database/sql"
	"time"
)

type DataDanInformasi struct {
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
	Keterangan             sql.NullString
	Tahun                  int
	CreatedAt              time.Time
	UpdatedAt              time.Time
	RadLevel1id            sql.NullInt32
	RadLevel2id            sql.NullInt32
	RadLevel3id            sql.NullInt32
	RadLevel4id            sql.NullInt32
	StrategicId            sql.NullInt32
	TacticalId             sql.NullInt32
	OperationalId          sql.NullInt32
}
