package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type OpdRepository interface {
	FetchKodeOpd(ctx context.Context, tx *sql.Tx) (web.Opd, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOPD string) []domain.Opd
}
