package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type RencanaPelaksanaanService interface {
	Create(ctx context.Context, request web.RencanaPelaksanaanCreateRequest) (web.RencanaPelaksanaanResponse, error)
	Update(ctx context.Context, request web.RencanaPelaksanaanUpdateRequest) (web.RencanaPelaksanaanResponse, error)
	Delete(ctx context.Context, rencanaId int, kodeOpd string, role string) error
	FindById(ctx context.Context, rencanaId int, kodeOpd string) (web.RencanaPelaksanaanResponse, error)
	FindAll(ctx context.Context, kodeOpd string, kebutuhanid int) ([]web.RencanaPelaksanaanResponse, error)
}
