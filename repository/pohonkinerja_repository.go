package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type PohonKinerjaRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, pohonkinerjaId int) (domain.PohonKinerja, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.PohonKinerja
}
