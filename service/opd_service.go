package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type OpdService interface {
	FetchKodeOpd(ctx context.Context) (web.Opd, error)
}
