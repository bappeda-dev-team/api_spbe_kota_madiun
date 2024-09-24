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
	UpdateKeterangan(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error)
	UpdatePenanggungJawab(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error)
	Delete(ctx context.Context, tx *sql.Tx, kebutuhanSPBEId int) error
	FindById(ctx context.Context, tx *sql.Tx, kebutuhanSPBEId int) (domain.KebutuhanSPBE, error)
	FindJenisKebutuhanByKebutuhanId(ctx context.Context, tx *sql.Tx, kebutuhanId int) ([]domain.JenisKebutuhan, error)
	FindKondisiAwalByJenisKebutuhanId(ctx context.Context, tx *sql.Tx, jenisKebutuhanId int) ([]domain.KondisiAwal, error)
	FindPenanggungJawab(ctx context.Context, tx *sql.Tx, pj string, tahun int) ([]domain.KebutuhanSPBE, error)
	FindIdPenanggungJawab(ctx context.Context, tx *sql.Tx, kebutuhanId int, pj string) (domain.KebutuhanSPBE, error)
}
