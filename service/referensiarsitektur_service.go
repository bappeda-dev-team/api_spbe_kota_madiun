package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type ReferensiArsitekturService interface {
	Insert(ctx context.Context, request web.ReferensiArsitekturCreateRequest) web.ReferensiArsitekturResponse
	Update(ctx context.Context, request web.ReferensiArsitekturUpdateRequest) web.ReferensiArsitekturResponse
	Delete(ctx context.Context, referensiarsitekturId int)
	FindAll(ctx context.Context) []web.ReferensiArsitekturResponse
	// FindById(ctx context.Context, referensiarsitekturId int) web.ReferensiArsitekturResponse
}
