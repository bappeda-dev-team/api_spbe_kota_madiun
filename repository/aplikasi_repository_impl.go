package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type AplikasiRepositoryImpl struct {
}

func NewAplikasiRepositoryImpl() *AplikasiRepositoryImpl {
	return &AplikasiRepositoryImpl{}
}

func (repository *AplikasiRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.Aplikasi, error) {
	script := `
		SELECT id, nama_aplikasi, fungsi_aplikasi, jenis_aplikasi, produsen_aplikasi, pj_aplikasi, informasi_terkait_input, informasi_terkait_output, interoprabilitas, keterangan, kode_opd, tahun, created_at, updated_at, raa_level_1_id, raa_level_2_id, raa_level_3_id, strategic_id, tactical_id, operational_id
		FROM aplikasi
		WHERE 1=1
	`

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

	var aplikasiList []domain.Aplikasi
	for rows.Next() {
		var aplikasi domain.Aplikasi
		err := rows.Scan(
			&aplikasi.Id,
			&aplikasi.NamaAplikasi,
			&aplikasi.FungsiAplikasi,
			&aplikasi.JenisAplikasi,
			&aplikasi.ProdusenAplikasi,
			&aplikasi.PjAplikasi,
			&aplikasi.InformasiTerkaitInput,
			&aplikasi.InformasiTerkaitOutput,
			&aplikasi.Interoprabilitas,
			&aplikasi.Keterangan,
			&aplikasi.KodeOPD,
			&aplikasi.Tahun,
			&aplikasi.CreatedAt,
			&aplikasi.UpdatedAt,
			&aplikasi.RaaLevel1id,
			&aplikasi.RaaLevel2id,
			&aplikasi.RaaLevel3id,
			&aplikasi.StrategicId,
			&aplikasi.TacticalId,
			&aplikasi.OperationalId,
		)
		if err != nil {
			return nil, err
		}
		aplikasiList = append(aplikasiList, aplikasi)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return aplikasiList, nil
}

func (repository *AplikasiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, aplikasiId int) (domain.Aplikasi, error) {
	script := "SELECT id, nama_aplikasi, fungsi_aplikasi, jenis_aplikasi, produsen_aplikasi, pj_aplikasi, informasi_terkait_input, informasi_terkait_output, interoprabilitas, keterangan, kode_opd, tahun, created_at, updated_at, raa_level_1_id, raa_level_2_id, raa_level_3_id, strategic_id, tactical_id, operational_id FROM aplikasi WHERE id = ?"
	rows, err := tx.QueryContext(ctx, script, aplikasiId)
	helper.PanicIfError(err)
	defer rows.Close()

	aplikasi := domain.Aplikasi{}

	if rows.Next() {
		err := rows.Scan(
			&aplikasi.Id,
			&aplikasi.NamaAplikasi,
			&aplikasi.FungsiAplikasi,
			&aplikasi.JenisAplikasi,
			&aplikasi.ProdusenAplikasi,
			&aplikasi.PjAplikasi,
			&aplikasi.InformasiTerkaitInput,
			&aplikasi.InformasiTerkaitOutput,
			&aplikasi.Interoprabilitas,
			&aplikasi.Keterangan,
			&aplikasi.KodeOPD,
			&aplikasi.Tahun,
			&aplikasi.CreatedAt,
			&aplikasi.UpdatedAt,
			&aplikasi.RaaLevel1id,
			&aplikasi.RaaLevel2id,
			&aplikasi.RaaLevel3id,
			&aplikasi.StrategicId,
			&aplikasi.TacticalId,
			&aplikasi.OperationalId,
		)
		helper.PanicIfError(err)
		return aplikasi, nil
	} else {
		return aplikasi, errors.New("aplikasi is not found")
	}
}

func (repository *AplikasiRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, aplikasi domain.Aplikasi) domain.Aplikasi {
	currentTime := time.Now()
	aplikasi.CreatedAt = currentTime
	aplikasi.UpdatedAt = currentTime

	script := "INSERT INTO aplikasi (nama_aplikasi, fungsi_aplikasi, jenis_aplikasi, produsen_aplikasi, pj_aplikasi, informasi_terkait_input, informasi_terkait_output, interoprabilitas, keterangan, kode_opd, tahun, created_at, updated_at, raa_level_1_id, raa_level_2_id, raa_level_3_id, strategic_id, tactical_id, operational_id) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, script,
		aplikasi.NamaAplikasi,
		aplikasi.FungsiAplikasi,
		aplikasi.JenisAplikasi,
		aplikasi.ProdusenAplikasi,
		aplikasi.PjAplikasi,
		aplikasi.InformasiTerkaitInput,
		aplikasi.InformasiTerkaitOutput,
		aplikasi.Interoprabilitas,
		aplikasi.Keterangan,
		aplikasi.KodeOPD,
		aplikasi.Tahun,
		aplikasi.CreatedAt,
		aplikasi.UpdatedAt,
		aplikasi.RaaLevel1id,
		aplikasi.RaaLevel2id,
		aplikasi.RaaLevel3id,
		aplikasi.StrategicId,
		aplikasi.TacticalId,
		aplikasi.OperationalId,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	aplikasi.Id = int(id)
	return aplikasi
}

func (repository *AplikasiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, aplikasi domain.Aplikasi) domain.Aplikasi {
	currentTime := time.Now()
	aplikasi.UpdatedAt = currentTime

	script := `
		UPDATE aplikasi SET 
			nama_aplikasi = ?, 
			fungsi_aplikasi = ?, 
			jenis_aplikasi = ?, 
			produsen_aplikasi = ?, 
			pj_aplikasi = ?, 
			informasi_terkait_input = ?, 
			informasi_terkait_output = ?, 
			interoprabilitas = ?, 
			keterangan = ?,
			kode_opd = ?, 
			tahun = ?, 
			updated_at = ?, 
			raa_level_1_id = ?, 
			raa_level_2_id = ?, 
			raa_level_3_id = ?, 
			strategic_id = ?, 
			tactical_id = ?, 
			operational_id = ? 
		WHERE id = ?
	`

	_, err := tx.ExecContext(ctx, script,
		aplikasi.NamaAplikasi,
		aplikasi.FungsiAplikasi,
		aplikasi.JenisAplikasi,
		aplikasi.ProdusenAplikasi,
		aplikasi.PjAplikasi,
		aplikasi.InformasiTerkaitInput,
		aplikasi.InformasiTerkaitOutput,
		aplikasi.Interoprabilitas,
		aplikasi.Keterangan,
		aplikasi.KodeOPD,
		aplikasi.Tahun,
		aplikasi.UpdatedAt,
		aplikasi.RaaLevel1id,
		aplikasi.RaaLevel2id,
		aplikasi.RaaLevel3id,
		aplikasi.StrategicId,
		aplikasi.TacticalId,
		aplikasi.OperationalId,
		aplikasi.Id,
	)
	helper.PanicIfError(err)

	return aplikasi
}
func (repository *AplikasiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, aplikasi domain.Aplikasi) {
	script := "delete from aplikasi where id =?"
	_, err := tx.ExecContext(ctx, script, aplikasi.Id)
	helper.PanicIfError(err)
}
