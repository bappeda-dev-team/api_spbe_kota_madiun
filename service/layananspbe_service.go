package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type LayananSpbeService interface {
	FindByKodeOpd(ctx context.Context, kodeOPD string, tahun int) ([]web.LayananSpbeRespons, error)
	FindById(ctx context.Context, layananspbeId int) (web.LayananSpbeRespons, error)
	Insert(ctx context.Context, request web.LayananSpbeCreateRequest) web.LayananSpbeRespons
	Update(ctx context.Context, request web.LayananSpbeUpdateRequest) web.LayananSpbeRespons
	Delete(ctx context.Context, layananspbeId int)
}
