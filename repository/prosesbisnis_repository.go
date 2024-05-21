package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type ProsesBisnisRepository interface {
	Save(ctx context.Context, tx *sql.Tx, prosesbisnis domain.ProsesBisnis)domain.ProsesBisnis
	Update(ctx context.Context, tx *sql.Tx, prosesbisnis domain.ProsesBisnis)domain.ProsesBisnis
	Delete(ctx context.Context, tx *sql.Tx, prosesbisnis domain.ProsesBisnis)
	FindById(ctx context.Context, tx *sql.Tx, Id int)(domain.ProsesBisnis, error)
	FindAll(ctx context.Context, tx *sql.Tx)[]domain.ProsesBisnis
	FindByKodeOdp(ctx context.Context, tx *sql.Tx, kode_opd string)(domain.ProsesBisnis, error)
}