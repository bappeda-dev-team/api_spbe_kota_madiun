package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type PetarencanaRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) ([]domain.Petarencana, error)
}
