package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type SasaranKotaRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, sasarankotaId int) (domain.SasaranKota, error)
}
