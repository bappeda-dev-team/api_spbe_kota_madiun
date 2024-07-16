package domain

import (
	"database/sql"
	"time"
)

type LayananSPBE struct {
	Id                int
	NamaLayanan       string
	KodeLayanan       string
	TujuanLayananId   sql.NullInt32
	FungsiLayanan     string
	Tahun             int
	KodeOPD           string
	KementrianTerkait string
	MetodeLayanan     string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	RalLevel1id       sql.NullInt32
	RalLevel2id       sql.NullInt32
	RalLevel3id       sql.NullInt32
	RalLevel4id       sql.NullInt32
	StrategicId       sql.NullInt32
	TacticalId        sql.NullInt32
	OperationalId     sql.NullInt32
}
