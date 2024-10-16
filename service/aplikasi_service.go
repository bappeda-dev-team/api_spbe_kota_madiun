package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type AplikasiService interface {
	FindByKodeOpd(ctx context.Context, kodeOPD string, tahun int) ([]web.AplikasiRespons, error)
	FindById(ctx context.Context, aplikasiId int, kodeOPD string) (web.AplikasiRespons, error)
	Insert(ctx context.Context, request web.AplikasiCreateRequest) web.AplikasiRespons
	Update(ctx context.Context, request web.AplikasiUpdateRequest) web.AplikasiRespons
	Delete(ctx context.Context, aplikasiId int, kodeOPD string, role string) error
}
