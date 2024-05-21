package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type ProsesBisnisService interface {
	FindAll(ctx context.Context) []web.ProsesBinsisRespons
}