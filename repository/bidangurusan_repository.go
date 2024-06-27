package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type BidangUrusanRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, bidangurusanId int) (domain.BidangUrusan, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BidangUrusan
}
