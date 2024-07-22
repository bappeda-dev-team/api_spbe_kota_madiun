package repository

import (
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type UrusanReposiotry interface {
	FetchUrusan(ctx context.Context, tx *sql.Tx) (web.UrusanOPD, error)
}
