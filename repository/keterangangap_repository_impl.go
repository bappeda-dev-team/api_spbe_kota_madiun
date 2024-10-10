package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type KeteranganGapRepositoryImpl struct {
}

func NewKeteranganGapRepositoryImpl() *KeteranganGapRepositoryImpl {
	return &KeteranganGapRepositoryImpl{}
}

func (repository *KeteranganGapRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, keteranganGap domain.KeteranganGap) (domain.KeteranganGap, error) {
	SQL := `INSERT INTO keterangan_gap (id, kode_opd, id_prosesbisnis, keterangan_gap) VALUES (?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, keteranganGap.Id, keteranganGap.KodeOpd, keteranganGap.IdProsesBisnis, keteranganGap.KeteranganGap)
	if err != nil {
		return domain.KeteranganGap{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.KeteranganGap{}, err
	}

	keteranganGap.Id = int(id)
	return keteranganGap, nil
}

func (repository *KeteranganGapRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.KeteranganGap, error) {
	SQL := `SELECT id, kode_opd, id_prosesbisnis, keterangan_gap FROM keterangan_gap WHERE id = ?`
	var keteranganGap domain.KeteranganGap
	err := tx.QueryRowContext(ctx, SQL, id).Scan(&keteranganGap.Id, &keteranganGap.KodeOpd, &keteranganGap.IdProsesBisnis, &keteranganGap.KeteranganGap)
	if err != nil {
		return domain.KeteranganGap{}, err
	}
	return keteranganGap, nil
}

func (repository *KeteranganGapRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, keteranganGap domain.KeteranganGap) (domain.KeteranganGap, error) {
	SQL := `UPDATE keterangan_gap SET kode_opd = ?, id_prosesbisnis = ?, keterangan_gap = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, keteranganGap.KodeOpd, keteranganGap.IdProsesBisnis, keteranganGap.KeteranganGap, keteranganGap.Id)
	if err != nil {
		return domain.KeteranganGap{}, err
	}
	return keteranganGap, nil
}

func (repository *KeteranganGapRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := `DELETE FROM keterangan_gap WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}
