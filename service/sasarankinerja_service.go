package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type SasaranKinerjaPegawaiService interface {
	FetchsasaranKinerja(ctx context.Context, kodeOPD string, tahun string) (web.ApiResponsSasaranKinerjaPegawai, error)
	FindByKodeOpdAndTahun(ctx context.Context, kodeOPD string, tahun string) []web.SasaranKinerjaPegawaiResponse
	FindById(ctx context.Context, sasaranKinerjaId int) (web.SasaranKinerjaPegawaiResponse, error)
}
