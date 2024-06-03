package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type ProsesBisnisRepository interface {
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int)([]domain.ProsesBisnis, error)
	FindById(ctx context.Context, tx *sql.Tx, prosesbisnisId int) (domain.ProsesBisnis, error)
	Insert(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) domain.ProsesBisnis
	Update(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) domain.ProsesBisnis
	Delete(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis)
}