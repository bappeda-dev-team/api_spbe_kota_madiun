package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type ProsesBisnisService interface {
	GetProsesBisnis(ctx context.Context, kodeOPD string, tahun int) ([]web.ProsesBisnisRespons, error)
	FindById(ctx context.Context, prosesbisnisId int, kodeOPD string) (web.ProsesBisnisRespons, error)
	Insert(ctx context.Context, request web.ProsesBisnisCreateRequest) web.ProsesBisnisRespons
	Update(ctx context.Context, request web.ProsesBisnisUpdateRequest) web.ProsesBisnisRespons
	Delete(ctx context.Context, prosesbisnisId int, kodeOPD string, role string) error
	GetProsesBisnisGrouped(ctx context.Context, kodeOPD string, tahun int) ([]web.GapProsesBisnis, error)
	GetProsesBisnisNoGap(ctx context.Context, kodeOPD string, tahun int) ([]web.GapProsesBisnis, error)
}
