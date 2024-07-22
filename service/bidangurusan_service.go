package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type BidangUrusanService interface {
	FindAll(ctx context.Context) []web.BidangUrusanRespons
	FetchUrusan(ctx context.Context) (web.BidangUrusanOPD, error)
}
