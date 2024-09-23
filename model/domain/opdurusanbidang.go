package domain

import "database/sql"

type OpdUrusanBidang struct {
	ID           int
	KodeOpd      sql.NullString
	KodeUrusan   sql.NullString
	BidangUrusan sql.NullString
}
