package repository

import (
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type UrusanRepository interface {
	FetchUrusan(ctx context.Context, tx *sql.Tx) (web.UrusanOPD, error)
}
