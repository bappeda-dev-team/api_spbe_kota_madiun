package domain

import "time"

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
	RabLevel4ID      int
	RabLevel5ID      int
	RabLevel6ID      int
	Tahun            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
