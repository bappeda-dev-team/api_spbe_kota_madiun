package domain

import "time"

type BidangUrusan struct {
	ID               int
	KodeBidangUrusan string
	BidangUrusan     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
