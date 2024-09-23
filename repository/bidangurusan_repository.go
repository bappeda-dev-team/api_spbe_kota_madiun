package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type BidangUrusanRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, bidangurusanId int) (domain.BidangUrusan, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BidangUrusan
	FetchBidangUrusan(ctx context.Context, tx *sql.Tx) (web.BidangUrusanOPD, error)
	FindByBidangUrusan(ctx context.Context, tx *sql.Tx, bidangUrusan string) (domain.BidangUrusan, error)
}
