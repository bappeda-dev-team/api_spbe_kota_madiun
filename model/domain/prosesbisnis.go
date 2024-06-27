package domain

import (
	"database/sql"
	"time"
)

type ProsesBisnis struct {
	ID               int
	NamaProsesBisnis string
	SasaranKotaId    sql.NullInt32
	KodeProsesBisnis string
	KodeOPD          string
	BidangUrusanId   sql.NullInt32
	RabLevel1ID      sql.NullInt32
	RabLevel2ID      sql.NullInt32
	RabLevel3ID      sql.NullInt32
	RabLevel4ID      sql.NullInt32
	RabLevel5ID      sql.NullInt32
	RabLevel6ID      sql.NullInt32
	Tahun            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
