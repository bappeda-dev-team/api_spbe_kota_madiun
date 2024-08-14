package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type KebutuhanSPBEService interface {
	Create(ctx context.Context, request web.KebutuhanSPBECreateRequest) (web.KebutuhanSPBEResponse, error)
	Update(ctx context.Context, request web.KebutuhanSPBEUpdateRequest) (web.KebutuhanSPBEResponse, error)
	Delete(ctx context.Context, kebutuhanSPBEId int, kodeOpd string) error
	FindById(ctx context.Context, kebutuhanSPBEId int, kodeOPD string) (web.KebutuhanSPBEResponse, error)
	FindByKodeOpdAndTahun(ctx context.Context, kodeOpd string, tahun int) ([]web.KebutuhanSPBEResponse, error)
}
