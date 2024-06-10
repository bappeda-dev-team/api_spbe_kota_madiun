package domain

import (
	"database/sql"
	"time"
)

type ProsesBisnis struct {
	ID               int
	NamaProsesBisnis string
	SasaranKota      string
	KodeProsesBisnis string
	KodeOPD          string
	BidangUrusan     string
	RabLevel1ID      int
	RabLevel2ID      int
	RabLevel3ID      int
	RabLevel4ID      sql.NullInt64
	RabLevel5ID      sql.NullInt64
	RabLevel6ID      sql.NullInt64
	Tahun            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
