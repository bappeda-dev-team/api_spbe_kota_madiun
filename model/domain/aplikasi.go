package domain

import (
	"database/sql"
	"time"
)

type Aplikasi struct {
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
	Keterangan             sql.NullString
	Tahun                  int
	CreatedAt              time.Time
	UpdatedAt              time.Time
	RaaLevel1id            sql.NullInt32
	RaaLevel2id            sql.NullInt32
	RaaLevel3id            sql.NullInt32
	StrategicId            sql.NullInt32
	TacticalId             sql.NullInt32
	OperationalId          sql.NullInt32
}
