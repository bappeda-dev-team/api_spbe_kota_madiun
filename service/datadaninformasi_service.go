package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type DataDanInformasiService interface {
	FindByKodeOpd(ctx context.Context, kodeOPD string, tahun int) ([]web.DataDanInformasiRespons, error)
	FindById(ctx context.Context, dataId int, kodeOPD string) (web.DataDanInformasiRespons, error)
	Insert(ctx context.Context, request web.DataDanInformasiCreateRequest) web.DataDanInformasiRespons
	Update(ctx context.Context, request web.DataDanInformasiUpdateRequest) web.DataDanInformasiRespons
	Delete(ctx context.Context, dataId int, kodeOPD string, role string) error
}
