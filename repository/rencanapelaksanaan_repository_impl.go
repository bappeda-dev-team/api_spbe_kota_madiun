package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"log"
)

type RencanaPelaksanaanRepositoryImpl struct{}

func NewRencanaPelaksanaanRepositoryImpl() *RencanaPelaksanaanRepositoryImpl {
	return &RencanaPelaksanaanRepositoryImpl{}
}

func (repository *RencanaPelaksanaanRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, rencanaPelaksanaan domain.RencanaPelaksanaanPegawai) (domain.RencanaPelaksanaanPegawai, error) {
	query := "INSERT INTO rencana_pelaksanaan (kode_opd, id_kebutuhan, id_sasarankinerja) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, rencanaPelaksanaan.KodeOpd, rencanaPelaksanaan.IdKebutuhan, rencanaPelaksanaan.IdSasaranKinerja)
	if err != nil {
		return rencanaPelaksanaan, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return rencanaPelaksanaan, err
	}
	rencanaPelaksanaan.Id = int(id)

	for _, tahun := range rencanaPelaksanaan.TahunPelaksanaan {
		tahunQuery := "INSERT INTO tahun_pelaksanaan (id_rencana_pelaksana, tahun) VALUES (?, ?)"
		_, err := tx.ExecContext(ctx, tahunQuery, rencanaPelaksanaan.Id, tahun.Tahun)
		if err != nil {
			return rencanaPelaksanaan, err
		}
	}

	return rencanaPelaksanaan, nil
}

func (repository *RencanaPelaksanaanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, rencanaPelaksanaan domain.RencanaPelaksanaanPegawai) (domain.RencanaPelaksanaanPegawai, error) {
	query := "UPDATE rencana_pelaksanaan SET kode_opd = ?, id_kebutuhan = ?, id_sasarankinerja = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, rencanaPelaksanaan.KodeOpd, rencanaPelaksanaan.IdKebutuhan, rencanaPelaksanaan.IdSasaranKinerja, rencanaPelaksanaan.Id)
	if err != nil {
		return rencanaPelaksanaan, err
	}

	// Hapus tahun pelaksanaan yang ada
	deleteQuery := "DELETE FROM tahun_pelaksanaan WHERE id_rencana_pelaksana = ?"
	_, err = tx.ExecContext(ctx, deleteQuery, rencanaPelaksanaan.Id)
	if err != nil {
		return rencanaPelaksanaan, err
	}

	// Tambahkan tahun pelaksanaan baru
	for _, tahun := range rencanaPelaksanaan.TahunPelaksanaan {
		tahunQuery := "INSERT INTO tahun_pelaksanaan (id_rencana_pelaksana, tahun) VALUES (?, ?)"
		_, err := tx.ExecContext(ctx, tahunQuery, rencanaPelaksanaan.Id, tahun.Tahun)
		if err != nil {
			return rencanaPelaksanaan, err
		}
	}

	return rencanaPelaksanaan, nil
}

func (repository *RencanaPelaksanaanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, rencanaPelaksanaanId int) error {
	log.Printf("Menghapus rencana pelaksanaan dengan ID: %d", rencanaPelaksanaanId)

	deleteTahunQuery := "DELETE FROM tahun_pelaksanaan WHERE id_rencana_pelaksana = ?"
	_, err := tx.ExecContext(ctx, deleteTahunQuery, rencanaPelaksanaanId)
	if err != nil {
		return err
	}

	log.Printf("Menghapus rencana pelaksanaan dengan ID: %d", rencanaPelaksanaanId)
	deleteRencanaQuery := "DELETE FROM rencana_pelaksanaan WHERE id = ?"
	_, err = tx.ExecContext(ctx, deleteRencanaQuery, rencanaPelaksanaanId)
	return err
}

func (repository *RencanaPelaksanaanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, kebutuhanId int) ([]domain.RencanaPelaksanaanPegawai, error) {
	query := "SELECT id, kode_opd, id_kebutuhan, id_sasarankinerja FROM rencana_pelaksanaan WHERE 1=1"
	var args []interface{}

	if kodeOpd != "" {
		query += " AND kode_opd = ?"
		args = append(args, kodeOpd)
	}

	if kebutuhanId != 0 {
		query += " AND id_kebutuhan = ?"
		args = append(args, kebutuhanId)
	}

	query += " ORDER BY id DESC"

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var rencana []domain.RencanaPelaksanaanPegawai
	for rows.Next() {
		var rencanaPelaksanaan domain.RencanaPelaksanaanPegawai
		if err := rows.Scan(&rencanaPelaksanaan.Id, &rencanaPelaksanaan.KodeOpd, &rencanaPelaksanaan.IdKebutuhan, &rencanaPelaksanaan.IdSasaranKinerja); err != nil {
			return nil, err
		}
		rencana = append(rencana, rencanaPelaksanaan)
	}
	return rencana, nil
}

func (repository *RencanaPelaksanaanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, rencanaId int) (domain.RencanaPelaksanaanPegawai, error) {
	query := "SELECT id, kode_opd, id_kebutuhan, id_sasarankinerja FROM rencana_pelaksanaan WHERE id = ?"
	args := []interface{}{rencanaId}

	row := tx.QueryRowContext(ctx, query, args...)
	var rencanaPelaksanaan domain.RencanaPelaksanaanPegawai
	err := row.Scan(&rencanaPelaksanaan.Id, &rencanaPelaksanaan.KodeOpd, &rencanaPelaksanaan.IdKebutuhan, &rencanaPelaksanaan.IdSasaranKinerja)
	if err != nil {
		return rencanaPelaksanaan, err
	}
	return rencanaPelaksanaan, nil
}

func (repository *RencanaPelaksanaanRepositoryImpl) FindIdTahunPelaksanaan(ctx context.Context, tx *sql.Tx, tahunId int) ([]domain.TahunPelaksanaan, error) {
	query := "SELECT id, id_rencana_pelaksana, tahun FROM tahun_pelaksanaan WHERE id_rencana_pelaksana = ?"
	rows, err := tx.QueryContext(ctx, query, tahunId)
	helper.PanicIfError(err)
	defer rows.Close()

	var tahunPelaksanaan []domain.TahunPelaksanaan
	for rows.Next() {
		var tahun domain.TahunPelaksanaan
		if err := rows.Scan(&tahun.Id, &tahun.IdRencanaPelaksana, &tahun.Tahun); err != nil {
			return nil, err
		}
		tahunPelaksanaan = append(tahunPelaksanaan, tahun)
	}
	return tahunPelaksanaan, nil
}

func (repository *RencanaPelaksanaanRepositoryImpl) FindByKebutuhanId(ctx context.Context, tx *sql.Tx, kebutuhanId int) ([]domain.RencanaPelaksanaanPegawai, error) {
	query := "SELECT id, kode_opd, id_kebutuhan, id_sasarankinerja FROM rencana_pelaksanaan WHERE id_kebutuhan = ?"
	rows, err := tx.QueryContext(ctx, query, kebutuhanId)
	helper.PanicIfError(err)
	defer rows.Close()

	var rencana []domain.RencanaPelaksanaanPegawai
	for rows.Next() {
		var rencanaPelaksanaan domain.RencanaPelaksanaanPegawai
		if err := rows.Scan(&rencanaPelaksanaan.Id, &rencanaPelaksanaan.KodeOpd, &rencanaPelaksanaan.IdKebutuhan, &rencanaPelaksanaan.IdSasaranKinerja); err != nil {
			return nil, err
		}
		rencana = append(rencana, rencanaPelaksanaan)
	}
	return rencana, nil

}
