package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type AplikasiRepository interface {
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.Aplikasi, error)
	FindById(ctx context.Context, tx *sql.Tx, aplikasiId int) (domain.Aplikasi, error)
	Insert(ctx context.Context, tx *sql.Tx, aplikasi domain.Aplikasi) domain.Aplikasi
	Update(ctx context.Context, tx *sql.Tx, aplikasi domain.Aplikasi) domain.Aplikasi
	Delete(ctx context.Context, tx *sql.Tx, aplikasi domain.Aplikasi)
}
