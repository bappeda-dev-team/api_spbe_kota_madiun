package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type PohonKinerjaRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, pohonkinerjaId int) (domain.PohonKinerja, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) []domain.PohonKinerja
	InsertApi(ctx context.Context, tx *sql.Tx) (web.PohonKinerjaApi, error)
}
