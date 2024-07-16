package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type UrusanService interface {
	FetchUrusan(ctx context.Context) (web.UrusanOPD, error)
}
