package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type RencanaPelaksanaanRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, rencanaPelaksanaan domain.RencanaPelaksanaanPegawai) (domain.RencanaPelaksanaanPegawai, error)
	Update(ctx context.Context, tx *sql.Tx, rencanaPelaksanaan domain.RencanaPelaksanaanPegawai) (domain.RencanaPelaksanaanPegawai, error)
	Delete(ctx context.Context, tx *sql.Tx, rencanaPelaksanaanId int) error
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, kebutuhanId int) ([]domain.RencanaPelaksanaanPegawai, error)
	FindById(ctx context.Context, tx *sql.Tx, rencanaId int) (domain.RencanaPelaksanaanPegawai, error)
	FindIdTahunPelaksanaan(ctx context.Context, tx *sql.Tx, tahunId int) ([]domain.TahunPelaksanaan, error)
	FindByKebutuhanId(ctx context.Context, tx *sql.Tx, kebutuhanId int) ([]domain.RencanaPelaksanaanPegawai, error)
}
