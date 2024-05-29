package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type ProsesBisnisRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx)([]domain.ProsesBisnis, error)
}