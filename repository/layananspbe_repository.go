package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type LayananSPBERepository interface {
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.LayananSPBE, error)
	FindById(ctx context.Context, tx *sql.Tx, LayananspbeId int) (domain.LayananSPBE, error)
	Insert(ctx context.Context, tx *sql.Tx, layananspbe domain.LayananSPBE) domain.LayananSPBE
	Update(ctx context.Context, tx *sql.Tx, layananspbe domain.LayananSPBE) domain.LayananSPBE
	Delete(ctx context.Context, tx *sql.Tx, layananspbe domain.LayananSPBE)
}
