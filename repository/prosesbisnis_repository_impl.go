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
	script := "SELECT id, nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, strategic_id, tactical_id, operational_id, tahun, created_at, updated_at FROM proses_bisnis WHERE 1=1"
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
		err := rows.Scan(&prosesBisnis.ID, &prosesBisnis.NamaProsesBisnis, &prosesBisnis.SasaranKotaId, &prosesBisnis.KodeProsesBisnis, &prosesBisnis.KodeOPD, &prosesBisnis.BidangUrusanId, &prosesBisnis.RabLevel1ID, &prosesBisnis.RabLevel2ID, &prosesBisnis.RabLevel3ID, &prosesBisnis.StrategicId, &prosesBisnis.TacticalId, &prosesBisnis.OperationalId, &prosesBisnis.Tahun, &prosesBisnis.CreatedAt, &prosesBisnis.UpdatedAt)
		helper.PanicIfError(err)
		prosesBisnisList = append(prosesBisnisList, prosesBisnis)
	}
	return prosesBisnisList, nil
}

func (repository *ProsesBisnisRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, prosesbisnisId int) (domain.ProsesBisnis, error) {
	script := "select id, nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, strategic_id, tactical_id, operational_id, tahun, created_at, updated_at from proses_bisnis where id = ?"
	rows, err := tx.QueryContext(ctx, script, prosesbisnisId)
	helper.PanicIfError(err)
	defer rows.Close()

	prosesBisnis := domain.ProsesBisnis{}
	if rows.Next() {
		err := rows.Scan(&prosesBisnis.ID, &prosesBisnis.NamaProsesBisnis, &prosesBisnis.SasaranKotaId, &prosesBisnis.KodeProsesBisnis, &prosesBisnis.KodeOPD, &prosesBisnis.BidangUrusanId, &prosesBisnis.RabLevel1ID, &prosesBisnis.RabLevel2ID, &prosesBisnis.RabLevel3ID, &prosesBisnis.StrategicId, &prosesBisnis.TacticalId, &prosesBisnis.OperationalId, &prosesBisnis.Tahun, &prosesBisnis.CreatedAt, &prosesBisnis.UpdatedAt)
		helper.PanicIfError(err)
		return prosesBisnis, nil
	} else {
		return prosesBisnis, errors.New("proses bisnis is not found")
	}
}

func (repository *ProsesBisnisRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) domain.ProsesBisnis {
	currentTime := time.Now()
	prosesBisnis.CreatedAt = currentTime
	prosesBisnis.UpdatedAt = currentTime

	script := "insert into proses_bisnis (nama_proses_bisnis, sasaran_kota_id, kode_proses_bisnis, kode_opd, bidang_urusan_id, rab_level_1_id, rab_level_2_id, rab_level_3_id, strategic_id, tactical_id, operational_id, tahun, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script,
		prosesBisnis.NamaProsesBisnis,
		prosesBisnis.SasaranKotaId,
		prosesBisnis.KodeProsesBisnis,
		prosesBisnis.KodeOPD,
		prosesBisnis.BidangUrusanId,
		prosesBisnis.RabLevel1ID,
		prosesBisnis.RabLevel2ID,
		prosesBisnis.RabLevel3ID,
		prosesBisnis.StrategicId,
		prosesBisnis.TacticalId,
		prosesBisnis.OperationalId,
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

	script := "update proses_bisnis set nama_proses_bisnis = ?, sasaran_kota_id = ?, kode_proses_bisnis = ?, kode_opd = ?, bidang_urusan_id = ?, rab_level_1_id = ?, rab_level_2_id = ?, rab_level_3_id = ?, strategic_id = ?, tactical_id = ?, operational_id = ?, tahun = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, script, prosesBisnis.NamaProsesBisnis, prosesBisnis.SasaranKotaId, prosesBisnis.KodeProsesBisnis, prosesBisnis.KodeOPD, prosesBisnis.BidangUrusanId, prosesBisnis.RabLevel1ID, prosesBisnis.RabLevel2ID, prosesBisnis.RabLevel3ID, prosesBisnis.StrategicId, prosesBisnis.TacticalId, prosesBisnis.OperationalId, prosesBisnis.Tahun, prosesBisnis.UpdatedAt, prosesBisnis.ID)
	helper.PanicIfError(err)

	return prosesBisnis
}

func (repository *ProsesBisnisRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, prosesBisnis domain.ProsesBisnis) {
	script := "delete from proses_bisnis where id =?"
	_, err := tx.ExecContext(ctx, script, prosesBisnis.ID)
	helper.PanicIfError(err)
}

func (repository *ProsesBisnisRepositoryImpl) GapProsesBisnis(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) ([]domain.GapProsesBisnis, error) {
	query := `
       SELECT
        pb.kode_opd,
		pb.tahun,
        pb.id,
        pb.nama_proses_bisnis,
        pb.kode_proses_bisnis,
        MAX(l.nama_layanan) AS nama_layanan,
        MAX(d.nama_data) AS nama_data,
        MAX(a.nama_aplikasi) AS nama_aplikasi
    FROM
        proses_bisnis pb
    LEFT JOIN
        layanan_spbe l ON l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id AND l.operational_id = pb.operational_id
        OR l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id OR l.strategic_id = pb.strategic_id
    LEFT JOIN
        data_dan_informasi d ON d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id AND d.operational_id = pb.operational_id
        OR d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id OR d.strategic_id = pb.strategic_id
    LEFT JOIN
        aplikasi a ON a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id AND a.operational_id = pb.operational_id
        OR a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id OR a.strategic_id = pb.strategic_id
    WHERE
        (l.strategic_id IS NOT NULL OR d.strategic_id IS NOT NULL OR a.strategic_id IS NOT NULL)
        AND (l.tactical_id IS NOT NULL OR d.tactical_id IS NOT NULL OR a.tactical_id IS NOT NULL)
    `

	var args []interface{}

	if kodeOpd != "" {
		query += " AND pb.kode_opd = ?"
		args = append(args, kodeOpd)
	}

	query += " GROUP BY pb.kode_opd, pb.id, pb.nama_proses_bisnis, pb.kode_proses_bisnis ORDER BY pb.id"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prosesBisnisList []domain.GapProsesBisnis

	for rows.Next() {
		var pb domain.GapProsesBisnis
		var namaLayanan sql.NullString
		var namaData sql.NullString
		var namaAplikasi sql.NullString

		if err := rows.Scan(
			&pb.KodeOpd,
			&pb.Tahun,
			&pb.ID,
			&pb.NamaProsesBisnis,
			&pb.KodeProsesBisnis,
			&namaLayanan,
			&namaData,
			&namaAplikasi,
		); err != nil {
			return nil, err
		}

		if namaLayanan.Valid {
			pb.Layanan = append(pb.Layanan, domain.GapLayanan{NamaLayanan: namaLayanan})
		}
		if namaData.Valid {
			pb.DataDanInformasi = append(pb.DataDanInformasi, domain.GapDataDanInformasi{NamaData: namaData})
		}
		if namaAplikasi.Valid {
			pb.Aplikasi = append(pb.Aplikasi, domain.GapAplikasi{NamaAplikasi: namaAplikasi})
		}

		prosesBisnisList = append(prosesBisnisList, pb)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prosesBisnisList, nil
}
