package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type OpdService interface {
	FetchAllData(ctx context.Context) ([]web.Opd, error)
	FindAll(ctx context.Context, kodeOPD string) []web.Opd
	FindKodeOpdUrusan(ctx context.Context, kodeOPD string) ([]web.OPD, error)
}
