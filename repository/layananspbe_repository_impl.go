package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type LayananSPBERepositoryImpl struct {
}

func NewLayananSPBERepository() LayananSPBERepository {
	return &LayananSPBERepositoryImpl{}
}

func (repository *LayananSPBERepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.LayananSPBE, error) {
	script := "SELECT id, nama_layanan, kode_layanan, tujuan_layanan_id, fungsi_layanan, tahun, kode_opd, kementrian_terkait, metode_layanan, created_at, updated_at, ral_level_1_id, ral_level_2_id, ral_level_3_id, ral_level_4_id, strategic_id, tactical_id, operational_id FROM layanan_spbe WHERE 1=1"
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

	var layananspbeList []domain.LayananSPBE
	for rows.Next() {
		var layananspbe domain.LayananSPBE
		err := rows.Scan(&layananspbe.Id, &layananspbe.NamaLayanan, &layananspbe.KodeLayanan, &layananspbe.TujuanLayananId, &layananspbe.FungsiLayanan, &layananspbe.Tahun, &layananspbe.KodeOPD, &layananspbe.KementrianTerkait, &layananspbe.MetodeLayanan, &layananspbe.CreatedAt, &layananspbe.UpdatedAt, &layananspbe.RalLevel1id, &layananspbe.RalLevel2id, &layananspbe.RalLevel3id, &layananspbe.RalLevel4id, &layananspbe.StrategicId, &layananspbe.TacticalId, &layananspbe.OperationalId)
		helper.PanicIfError(err)
		layananspbeList = append(layananspbeList, layananspbe)
	}
	return layananspbeList, nil
}

func (repository *LayananSPBERepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, LayananspbeId int) (domain.LayananSPBE, error) {
	script := "SELECT id, nama_layanan, kode_layanan, tujuan_layanan_id, fungsi_layanan, tahun, kode_opd, kementrian_terkait, metode_layanan, created_at, updated_at, ral_level_1_id, ral_level_2_id, ral_level_3_id, ral_level_4_id, strategic_id, tactical_id, operational_id  FROM layanan_spbe WHERE id =?"
	rows, err := tx.QueryContext(ctx, script, LayananspbeId)
	helper.PanicIfError(err)
	defer rows.Close()

	layananspbe := domain.LayananSPBE{}

	if rows.Next() {
		err := rows.Scan(&layananspbe.Id, &layananspbe.NamaLayanan, &layananspbe.KodeLayanan, &layananspbe.TujuanLayananId, &layananspbe.FungsiLayanan, &layananspbe.Tahun, &layananspbe.KodeOPD, &layananspbe.KementrianTerkait, &layananspbe.MetodeLayanan, &layananspbe.CreatedAt, &layananspbe.UpdatedAt, &layananspbe.RalLevel1id, &layananspbe.RalLevel2id, &layananspbe.RalLevel3id, &layananspbe.RalLevel4id, &layananspbe.StrategicId, &layananspbe.TacticalId, &layananspbe.OperationalId)
		helper.PanicIfError(err)
		return layananspbe, nil
	} else {
		return layananspbe, errors.New("ayanan Spbe is not found")
	}
}

func (repository *LayananSPBERepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, layananspbe domain.LayananSPBE) domain.LayananSPBE {
	currentTime := time.Now()
	layananspbe.CreatedAt = currentTime
	layananspbe.UpdatedAt = currentTime
	script := "INSERT INTO layanan_spbe (nama_layanan, kode_layanan, tujuan_layanan_id, fungsi_layanan, tahun, kode_opd, kementrian_terkait, metode_layanan, created_at, updated_at, ral_level_1_id, ral_level_2_id, ral_level_3_id, ral_level_4_id, strategic_id, tactical_id, operational_id) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, script,
		layananspbe.NamaLayanan,
		layananspbe.KodeLayanan,
		layananspbe.TujuanLayananId,
		layananspbe.FungsiLayanan,
		layananspbe.Tahun,
		layananspbe.KodeOPD,
		layananspbe.KementrianTerkait,
		layananspbe.MetodeLayanan,
		layananspbe.CreatedAt,
		layananspbe.UpdatedAt,
		layananspbe.RalLevel1id,
		layananspbe.RalLevel2id,
		layananspbe.RalLevel3id,
		layananspbe.RalLevel4id,
		layananspbe.StrategicId,
		layananspbe.TacticalId,
		layananspbe.OperationalId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	layananspbe.Id = int(id)
	return layananspbe
}

func (repository *LayananSPBERepositoryImpl) Update(ctx context.Context, tx *sql.Tx, layananspbe domain.LayananSPBE) domain.LayananSPBE {
	script := "update layanan_spbe set nama_layanan = ?, kode_layanan = ?, tujuan_layanan_id = ?, fungsi_layanan = ?, tahun = ?, kode_opd = ?, kementrian_terkait = ?, metode_layanan = ?, ral_level_1_id = ?, ral_level_2_id = ?, ral_level_3_id = ?, ral_level_4_id = ?, strategic_id = ?, tactical_id  = ?, operational_id = ? where id = ?"

	_, err := tx.ExecContext(ctx, script,
		layananspbe.NamaLayanan,
		layananspbe.KodeLayanan,
		layananspbe.TujuanLayananId,
		layananspbe.FungsiLayanan,
		layananspbe.Tahun,
		layananspbe.KodeOPD,
		layananspbe.KementrianTerkait,
		layananspbe.MetodeLayanan,
		layananspbe.RalLevel1id,
		layananspbe.RalLevel2id,
		layananspbe.RalLevel3id,
		layananspbe.RalLevel4id,
		layananspbe.StrategicId,
		layananspbe.TacticalId,
		layananspbe.OperationalId,
		layananspbe.Id,
	)
	helper.PanicIfError(err)

	return layananspbe
}
func (repository *LayananSPBERepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, layananspbe domain.LayananSPBE) {
	script := "delete from layanan_spbe where id =?"
	_, err := tx.ExecContext(ctx, script, layananspbe.Id)
	helper.PanicIfError(err)

}
