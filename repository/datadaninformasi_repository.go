package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type DataDanInformasiRepository interface {
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.DataDanInformasi, error)
	FindById(ctx context.Context, tx *sql.Tx, dataId int) (domain.DataDanInformasi, error)
	Insert(ctx context.Context, tx *sql.Tx, datainformasi domain.DataDanInformasi) domain.DataDanInformasi
	Update(ctx context.Context, tx *sql.Tx, datainformasi domain.DataDanInformasi) domain.DataDanInformasi
	Delete(ctx context.Context, tx *sql.Tx, datainformasi domain.DataDanInformasi)
}
