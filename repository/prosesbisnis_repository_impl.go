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

func NewProsesBisnisRepositoryImpl() *ProsesBisnisRepositoryImpl {
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
        pb.id,
        pb.kode_opd,
        pb.tahun,
        pb.nama_proses_bisnis,
        pb.kode_proses_bisnis,
        l.nama_layanan,
        d.nama_data,
        a.nama_aplikasi
    FROM
        proses_bisnis pb
    LEFT JOIN
        layanan_spbe l ON (l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id AND l.operational_id = pb.operational_id  AND l.kode_opd = pb.kode_opd)
        OR (l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id  AND l.kode_opd = pb.kode_opd)
        OR (l.strategic_id = pb.strategic_id  AND l.kode_opd = pb.kode_opd)
    LEFT JOIN
        data_dan_informasi d ON (d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id AND d.operational_id = pb.operational_id  AND l.kode_opd = pb.kode_opd)
        OR (d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id  AND l.kode_opd = pb.kode_opd)
        OR (d.strategic_id = pb.strategic_id  AND d.kode_opd = pb.kode_opd)
    LEFT JOIN
        aplikasi a ON (a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id AND a.operational_id = pb.operational_id  AND l.kode_opd = pb.kode_opd)
        OR (a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id  AND l.kode_opd = pb.kode_opd)
        OR (a.strategic_id = pb.strategic_id  AND a.kode_opd = pb.kode_opd)
    WHERE 1=1
	`

	var args []interface{}
	if kodeOpd != "" {
		query += " AND pb.kode_opd = ?"
		args = append(args, kodeOpd)
	}
	if tahun != 0 {
		query += " AND pb.tahun = ?"
		args = append(args, tahun)
	}

	query += " ORDER BY pb.kode_opd, pb.id;"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prosesBisnisMap := make(map[int]*domain.GapProsesBisnis)

	for rows.Next() {
		var id int
		var kodeOpd string
		var tahun int
		var namaProsesBisnis, kodeProsesBisnis string
		var namaLayanan, namaData, namaAplikasi sql.NullString

		if err := rows.Scan(
			&id,
			&kodeOpd,
			&tahun,
			&namaProsesBisnis,
			&kodeProsesBisnis,
			&namaLayanan,
			&namaData,
			&namaAplikasi,
		); err != nil {
			return nil, err
		}

		pb, exists := prosesBisnisMap[id]
		if !exists {
			pb = &domain.GapProsesBisnis{
				ID:               id,
				KodeOpd:          kodeOpd,
				Tahun:            tahun,
				NamaProsesBisnis: namaProsesBisnis,
				KodeProsesBisnis: kodeProsesBisnis,
			}
			prosesBisnisMap[id] = pb
		}

		if namaLayanan.Valid {
			pb.Layanan = append(pb.Layanan, domain.GapLayanan{
				NamaLayanan: sql.NullString{
					String: namaLayanan.String,
					Valid:  true,
				},
			})
		} else {
			pb.Layanan = append(pb.Layanan, domain.GapLayanan{
				NamaLayanan: sql.NullString{},
			})
		}
		if namaData.Valid {
			pb.DataDanInformasi = append(
				pb.DataDanInformasi, domain.GapDataDanInformasi{
					NamaData: sql.NullString{
						String: namaData.String,
						Valid:  true,
					},
				})
		} else {
			pb.DataDanInformasi = append(pb.DataDanInformasi, domain.GapDataDanInformasi{
				NamaData: sql.NullString{},
			})
		}
		if namaAplikasi.Valid {
			pb.Aplikasi = append(
				pb.Aplikasi, domain.GapAplikasi{
					NamaAplikasi: sql.NullString{
						String: namaAplikasi.String,
						Valid:  true,
					},
				})
		} else {
			pb.Aplikasi = append(pb.Aplikasi, domain.GapAplikasi{
				NamaAplikasi: sql.NullString{},
			})
		}
	}

	var prosesBisnisList []domain.GapProsesBisnis
	for _, pb := range prosesBisnisMap {
		prosesBisnisList = append(prosesBisnisList, *pb)
	}

	return prosesBisnisList, nil
}

func (repository *ProsesBisnisRepositoryImpl) NoGapProsesBisnis(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) ([]domain.GapProsesBisnis, error) {
	query := `
	   SELECT
        pb.id,
        pb.kode_opd,
        pb.tahun,
        pb.nama_proses_bisnis,
        pb.kode_proses_bisnis,
        l.nama_layanan,
        d.nama_data,
        a.nama_aplikasi
    FROM
        proses_bisnis pb
    LEFT JOIN
        layanan_spbe l ON (l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id AND l.operational_id = pb.operational_id  AND l.kode_opd = pb.kode_opd)
        OR (l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id  AND l.kode_opd = pb.kode_opd)
        OR (l.strategic_id = pb.strategic_id  AND l.kode_opd = pb.kode_opd)
    LEFT JOIN
        data_dan_informasi d ON (d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id AND d.operational_id = pb.operational_id  AND l.kode_opd = pb.kode_opd)
        OR (d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id  AND l.kode_opd = pb.kode_opd)
        OR (d.strategic_id = pb.strategic_id  AND d.kode_opd = pb.kode_opd)
    LEFT JOIN
        aplikasi a ON (a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id AND a.operational_id = pb.operational_id  AND l.kode_opd = pb.kode_opd)
        OR (a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id  AND l.kode_opd = pb.kode_opd)
        OR (a.strategic_id = pb.strategic_id  AND a.kode_opd = pb.kode_opd)
    WHERE 1=1
    AND l.nama_layanan IS NOT NULL
    AND d.nama_data IS NOT NULL
    AND a.nama_aplikasi IS NOT NULL


	`

	var args []interface{}
	if kodeOpd != "" {
		query += " AND pb.kode_opd = ?"
		args = append(args, kodeOpd)
	}
	if tahun != 0 {
		query += " AND pb.tahun = ?"
		args = append(args, tahun)
	}

	query += " ORDER BY pb.kode_opd, pb.id;"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prosesBisnisMap := make(map[int]*domain.GapProsesBisnis)

	for rows.Next() {
		var id int
		var kodeOpd string
		var tahun int
		var namaProsesBisnis, kodeProsesBisnis string
		var namaLayanan, namaData, namaAplikasi sql.NullString

		if err := rows.Scan(
			&id,
			&kodeOpd,
			&tahun,
			&namaProsesBisnis,
			&kodeProsesBisnis,
			&namaLayanan,
			&namaData,
			&namaAplikasi,
		); err != nil {
			return nil, err
		}

		pb, exists := prosesBisnisMap[id]
		if !exists {
			pb = &domain.GapProsesBisnis{
				ID:               id,
				KodeOpd:          kodeOpd,
				Tahun:            tahun,
				NamaProsesBisnis: namaProsesBisnis,
				KodeProsesBisnis: kodeProsesBisnis,
			}
			prosesBisnisMap[id] = pb
		}

		if namaLayanan.Valid {
			pb.Layanan = append(pb.Layanan, domain.GapLayanan{
				NamaLayanan: sql.NullString{
					String: namaLayanan.String,
					Valid:  true,
				},
			})
		}
		if namaData.Valid {
			pb.DataDanInformasi = append(
				pb.DataDanInformasi, domain.GapDataDanInformasi{
					NamaData: sql.NullString{
						String: namaData.String,
						Valid:  true,
					},
				})
		}
		if namaAplikasi.Valid {
			pb.Aplikasi = append(
				pb.Aplikasi, domain.GapAplikasi{
					NamaAplikasi: sql.NullString{
						String: namaAplikasi.String,
						Valid:  true,
					},
				})
		}
	}

	var prosesBisnisList []domain.GapProsesBisnis
	for _, pb := range prosesBisnisMap {
		prosesBisnisList = append(prosesBisnisList, *pb)
	}

	return prosesBisnisList, nil
}
