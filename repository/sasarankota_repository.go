package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type SasaranKotaRepository interface {
	InsertApi(ctx context.Context, tx *sql.Tx) (web.SasaranKotaApi, error)
	FindById(ctx context.Context, tx *sql.Tx, sasarankotaId int) (domain.SasaranKota, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.SasaranKota
}
