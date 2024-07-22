package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type SasaranKotaService interface {
	FindById(ctx context.Context, sasarankotaId int) web.SasaranKotaRespons
	FindAll(ctx context.Context) []web.SasaranKotaRespons
	InsertApi(ctx context.Context) (web.SasaranKotaApi, error)
}
