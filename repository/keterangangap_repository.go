package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type KeteranganGapRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, keteranganGap domain.KeteranganGap) (domain.KeteranganGap, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.KeteranganGap, error)
	Update(ctx context.Context, tx *sql.Tx, keteranganGap domain.KeteranganGap) (domain.KeteranganGap, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
