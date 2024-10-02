package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type KeteranganGapService interface {
	Insert(ctx context.Context, keteranganGap web.KeteranganGapCreateRequest) (web.KeteranganGapRespons, error)
	FindById(ctx context.Context, id int) (web.KeteranganGapRespons, error)
	Update(ctx context.Context, keteranganGap web.KeteranganGapUpdateRequest) (web.KeteranganGapRespons, error)
	Delete(ctx context.Context, id int) (web.KeteranganGapRespons, error)
}
