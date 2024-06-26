package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ProsesBisnisRepositoryImpl struct {
}

func NewProsesBisnisRepository() ProsesBisnisRepository {
	return &ProsesBisnisRepositoryImpl{}
}

func (repository *ProsesBisnisRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.ProsesBisnis, error) {
	script := "SELECT id, nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, rab_level_4_id, rab_level_5_id, rab_level_6_id, tahun, created_at, updated_at FROM proses_bisnis WHERE 1=1"
	var args []interface{}

	if kodeOPD != "" {
		script += " AND kode_opd = ?"
		args = append(args, kodeOPD)
	}
	if tahun != 0 {
		script += " AND tahun = ?"
		args = append(args, tahun)
	}

	rows, err := tx.QueryContext(ctx, script, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var prosesBisnisList []domain.ProsesBisnis
	for rows.Next() {
		var prosesBisnis domain.ProsesBisnis
		err := rows.Scan(&prosesBisnis.ID, &prosesBisnis.NamaProsesBisnis, &prosesBisnis.SasaranKotaId, &prosesBisnis.KodeProsesBisnis, &prosesBisnis.KodeOPD, &prosesBisnis.BidangUrusanId, &prosesBisnis.RabLevel1ID, &prosesBisnis.RabLevel2ID, &prosesBisnis.RabLevel3ID, &prosesBisnis.RabLevel4ID, &prosesBisnis.RabLevel5ID, &prosesBisnis.RabLevel6ID, &prosesBisnis.Tahun, &prosesBisnis.CreatedAt, &prosesBisnis.UpdatedAt)
		helper.PanicIfError(err)
		prosesBisnisList = append(prosesBisnisList, prosesBisnis)
	}
	return prosesBisnisList, nil
}

func (repository *ProsesBisnisRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, prosesbisnisId int) (domain.ProsesBisnis, error) {
	script := "select id, nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, rab_level_4_id, rab_level_5_id, rab_level_6_id, tahun, created_at, updated_at from proses_bisnis where id = ?"
	rows, err := tx.QueryContext(ctx, script, prosesbisnisId)
	helper.PanicIfError(err)
	defer rows.Close()

	prosesBisnis := domain.ProsesBisnis{}
	if rows.Next() {
		err := rows.Scan(&prosesBisnis.ID, &prosesBisnis.NamaProsesBisnis, &prosesBisnis.SasaranKotaId, &prosesBisnis.KodeProsesBisnis, &prosesBisnis.KodeOPD, &prosesBisnis.BidangUrusanId, &prosesBisnis.RabLevel1ID, &prosesBisnis.RabLevel2ID, &prosesBisnis.RabLevel3ID, &prosesBisnis.RabLevel4ID, &prosesBisnis.RabLevel5ID, &prosesBisnis.RabLevel6ID, &prosesBisnis.Tahun, &prosesBisnis.CreatedAt, &prosesBisnis.UpdatedAt)
		helper.PanicIfError(err)
		return prosesBisnis, nil
	} else {
		return prosesBisnis, errors.New("Proses bisnis is not found")
	}
}

func (repository *ProsesBisnisRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) domain.ProsesBisnis {
	currentTime := time.Now()
	prosesBisnis.CreatedAt = currentTime
	prosesBisnis.UpdatedAt = currentTime

	script := "insert into proses_bisnis (nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, rab_level_4_id, rab_level_5_id, rab_level_6_id, tahun, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)"
	result, err := tx.ExecContext(ctx, script,
		prosesBisnis.NamaProsesBisnis,
		prosesBisnis.SasaranKotaId,
		prosesBisnis.KodeProsesBisnis,
		prosesBisnis.KodeOPD,
		prosesBisnis.BidangUrusanId,
		prosesBisnis.RabLevel1ID,
		prosesBisnis.RabLevel2ID,
		prosesBisnis.RabLevel3ID,
		prosesBisnis.RabLevel4ID,
		prosesBisnis.RabLevel5ID,
		prosesBisnis.RabLevel6ID,
		prosesBisnis.Tahun,
		prosesBisnis.CreatedAt,
		prosesBisnis.UpdatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	prosesBisnis.ID = int(id)
	return prosesBisnis
}

func (repository *ProsesBisnisRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) domain.ProsesBisnis {
	currentTime := time.Now()
	prosesBisnis.UpdatedAt = currentTime

	script := "update proses_bisnis set nama_proses_bisnis = ?, sasaran_kota_id = ?, kode_proses_bisnis = ?, kode_opd = ?, bidang_urusan_id = ?, rab_level_1_id = ?, rab_level_2_id = ?, rab_level_3_id = ?, rab_level_4_id = ?, rab_level_5_id = ?, rab_level_6_id = ?, tahun = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, script, prosesBisnis.NamaProsesBisnis, prosesBisnis.SasaranKotaId, prosesBisnis.KodeProsesBisnis, prosesBisnis.KodeOPD, prosesBisnis.BidangUrusanId, prosesBisnis.RabLevel1ID, prosesBisnis.RabLevel2ID, prosesBisnis.RabLevel3ID, prosesBisnis.RabLevel4ID, prosesBisnis.RabLevel5ID, prosesBisnis.RabLevel6ID, prosesBisnis.Tahun, prosesBisnis.UpdatedAt, prosesBisnis.ID)
	helper.PanicIfError(err)

	return prosesBisnis
}

func (repository *ProsesBisnisRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) {
	script := "delete from proses_bisnis where id =?"
	_, err := tx.ExecContext(ctx, script, prosesBisnis.ID)
	helper.PanicIfError(err)
}

func (repository *ProsesBisnisRepositoryImpl) FindByNull(ctx context.Context, tx *sql.Tx) ([]domain.ProsesBisnis, error) {
	script := "SELECT id, nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, rab_level_4_id, rab_level_5_id, rab_level_6_id, tahun, created_at, updated_at FROM proses_bisnis WHERE nama_proses_bisnis is null or sasaran_kota_id is null or kode_proses_bisnis is null or kode_opd is null or bidang_urusan_id is null or rab_level_1_id is null or rab_level_2_id is null or rab_level_3_id is null or rab_level_4_id is null or rab_level_5_id is null or rab_level_6_id is null or tahun is null"

	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var prosesBisnisList []domain.ProsesBisnis
	for rows.Next() {
		var prosesBisnis domain.ProsesBisnis
		err := rows.Scan(&prosesBisnis.ID, &prosesBisnis.NamaProsesBisnis, &prosesBisnis.SasaranKotaId, &prosesBisnis.KodeProsesBisnis, &prosesBisnis.KodeOPD, &prosesBisnis.BidangUrusanId, &prosesBisnis.RabLevel1ID, &prosesBisnis.RabLevel2ID, &prosesBisnis.RabLevel3ID, &prosesBisnis.RabLevel4ID, &prosesBisnis.RabLevel5ID, &prosesBisnis.RabLevel6ID, &prosesBisnis.Tahun, &prosesBisnis.CreatedAt, &prosesBisnis.UpdatedAt)
		helper.PanicIfError(err)
		prosesBisnisList = append(prosesBisnisList, prosesBisnis)
	}
	return prosesBisnisList, nil

}
