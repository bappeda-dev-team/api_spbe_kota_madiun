package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type ProsesBisnisService interface {
	GetProsesBisnis(ctx context.Context) ([]web.ProsesBisnisRespons, error)
}