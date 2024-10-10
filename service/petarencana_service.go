package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type PetarencanaService interface {
	FindAll(ctx context.Context, kodeOpd string, tahun int) ([]web.PetarencanaRespons, error)
}
