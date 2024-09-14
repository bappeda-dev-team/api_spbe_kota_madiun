package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type SasaranKinerjaPegawaiRepository interface {
	FetchsasaranKinerja(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun string) (web.ApiResponsSasaranKinerjaPegawai, error)
	FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun string) []domain.SasaranKinerja
	FindById(ctx context.Context, tx *sql.Tx, sasarankinerjaId int) (domain.SasaranKinerja, error)
}
