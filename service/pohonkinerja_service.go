package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type PohonKinerjaService interface {
	FindById(ctx context.Context, pohonId int) web.PohonKinerjaRespons
	FindAll(ctx context.Context, kodeOpd string, tahun int) []web.PohonKinerjaRespons
	InsertApi(ctx context.Context) (web.PohonKinerjaApi, error)
	FindByOperational(ctx context.Context, pohonId int) web.PohonKinerjaHierarchyResponse
}
