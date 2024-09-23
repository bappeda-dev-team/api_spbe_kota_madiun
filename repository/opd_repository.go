package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type OpdRepository interface {
	FetchAllData(ctx context.Context, tx *sql.Tx) ([]web.Opd, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOPD string) []domain.Opd
	FindById(ctx context.Context, tx *sql.Tx, opdId string) (domain.Opd, error)
	FindKodeOpdUrusan(ctx context.Context, tx *sql.Tx, kodeOPD string) ([]domain.OpdUrusanBidang, error)
}
