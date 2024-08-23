package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type KebutuhanSPBERepository interface {
	Insert(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error)
	FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int, prosesbisnis int) ([]domain.KebutuhanSPBE, error)
	Update(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error)
	Delete(ctx context.Context, tx *sql.Tx, kebutuhanSPBEId int) error
	FindById(ctx context.Context, tx *sql.Tx, kebutuhanSPBEId int) (domain.KebutuhanSPBE, error)
	FindJenisKebutuhanByKebutuhanId(ctx context.Context, tx *sql.Tx, kebutuhanId int) ([]domain.JenisKebutuhan, error)
	FindKondisiAwalByJenisKebutuhanId(ctx context.Context, tx *sql.Tx, jenisKebutuhanId int) ([]domain.KondisiAwal, error)
}
