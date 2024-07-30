package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type DataDanInformasiRepositoryImpl struct {
}

func NewDataDanInformasiRepository() DataDanInformasiRepository {
	return &DataDanInformasiRepositoryImpl{}
}

func (repository *DataDanInformasiRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun int) ([]domain.DataDanInformasi, error) {
	script := "SELECT id, nama_data, sifat_data, jenis_data, produsen_data, uraian_data, validitas_data, pj_data, informasi_terkait_input,informasi_terkait_output,interoprabilitas, kode_opd, tahun, created_at, updated_at, rad_level_1_id, rad_level_2_id, rad_level_3_id, rad_level_4_id, strategic_id, tactical_id, operational_id FROM data_dan_informasi WHERE 1=1"
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

	var datainformasiList []domain.DataDanInformasi
	for rows.Next() {
		var data domain.DataDanInformasi
		err := rows.Scan(&data.Id, &data.NamaData, &data.SifatData, &data.JenisData, &data.ProdusenData, &data.UraianData, &data.ValiditasData, &data.PjData, &data.InformasiTerkaitInput, &data.InformasiTerkaitOutput, &data.Interoprabilitas, &data.KodeOPD, &data.Tahun, &data.CreatedAt, &data.UpdatedAt, &data.RadLevel1id, &data.RadLevel2id, &data.RadLevel3id, &data.RadLevel4id, &data.StrategicId, &data.TacticalId, &data.OperationalId)
		helper.PanicIfError(err)
		datainformasiList = append(datainformasiList, data)
	}
	return datainformasiList, nil
}
func (repository *DataDanInformasiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, dataId int) (domain.DataDanInformasi, error) {
	script := "SELECT id, nama_data, sifat_data, jenis_data, produsen_data, uraian_data, validitas_data, pj_data, informasi_terkait_input,informasi_terkait_output,interoprabilitas, kode_opd, tahun,created_at, updated_at, rad_level_1_id, rad_level_2_id, rad_level_3_id, rad_level_4_id, strategic_id, tactical_id, operational_id FROM data_dan_informasi WHERE id =?"
	rows, err := tx.QueryContext(ctx, script, dataId)
	helper.PanicIfError(err)
	defer rows.Close()

	data := domain.DataDanInformasi{}

	if rows.Next() {
		err := rows.Scan(&data.Id, &data.NamaData, &data.SifatData, &data.JenisData, &data.ProdusenData, &data.UraianData, &data.ValiditasData, &data.PjData, &data.InformasiTerkaitInput, &data.InformasiTerkaitOutput, &data.Interoprabilitas, &data.KodeOPD, &data.Tahun, &data.CreatedAt, &data.UpdatedAt, &data.RadLevel1id, &data.RadLevel2id, &data.RadLevel3id, &data.RadLevel4id, &data.StrategicId, &data.TacticalId, &data.OperationalId)
		helper.PanicIfError(err)
		return data, nil
	} else {
		return data, errors.New("data & informasi is not found")
	}
}
func (repository *DataDanInformasiRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, datainformasi domain.DataDanInformasi) domain.DataDanInformasi {
	currentTime := time.Now()
	datainformasi.CreatedAt = currentTime
	datainformasi.UpdatedAt = currentTime
	script := "INSERT INTO data_dan_informasi (nama_data, sifat_data, jenis_data, produsen_data, uraian_data, validitas_data, pj_data, informasi_terkait_input,informasi_terkait_output,interoprabilitas, kode_opd, tahun, created_at, updated_at, rad_level_1_id, rad_level_2_id, rad_level_3_id, rad_level_4_id, strategic_id, tactical_id, operational_id ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, script,
		datainformasi.NamaData,
		datainformasi.SifatData,
		datainformasi.JenisData,
		datainformasi.ProdusenData,
		datainformasi.UraianData,
		datainformasi.ValiditasData,
		datainformasi.PjData,
		datainformasi.InformasiTerkaitInput,
		datainformasi.InformasiTerkaitOutput,
		datainformasi.Interoprabilitas,
		datainformasi.KodeOPD,
		datainformasi.Tahun,
		datainformasi.CreatedAt,
		datainformasi.UpdatedAt,
		datainformasi.RadLevel1id,
		datainformasi.RadLevel2id,
		datainformasi.RadLevel3id,
		datainformasi.RadLevel4id,
		datainformasi.StrategicId,
		datainformasi.TacticalId,
		datainformasi.OperationalId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	datainformasi.Id = int(id)
	return datainformasi
}
func (repository *DataDanInformasiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, datainformasi domain.DataDanInformasi) domain.DataDanInformasi {
	currentTime := time.Now()
	datainformasi.UpdatedAt = currentTime

	script := "UPDATE data_dan_informasi SET nama_data = ?, sifat_data = ?, jenis_data = ?, produsen_data = ?, uraian_data = ?, validitas_data = ?,pj_data = ?, informasi_terkait_input = ?, informasi_terkait_output = ?, interoprabilitas = ?, kode_opd = ?, tahun = ?, updated_at = ?, rad_level_1_id = ?, rad_level_2_id = ?, rad_level_3_id = ?, 	rad_level_4_id = ?, strategic_id = ?, tactical_id = ?, operational_id = ?	WHERE id = ?"

	_, err := tx.ExecContext(ctx, script,
		datainformasi.NamaData,
		datainformasi.SifatData,
		datainformasi.ProdusenData,
		datainformasi.JenisData,
		datainformasi.UraianData,
		datainformasi.ValiditasData,
		datainformasi.PjData,
		datainformasi.InformasiTerkaitInput,
		datainformasi.InformasiTerkaitOutput,
		datainformasi.Interoprabilitas,
		datainformasi.KodeOPD,
		datainformasi.Tahun,
		datainformasi.UpdatedAt,
		datainformasi.RadLevel1id,
		datainformasi.RadLevel2id,
		datainformasi.RadLevel3id,
		datainformasi.RadLevel4id,
		datainformasi.StrategicId,
		datainformasi.TacticalId,
		datainformasi.OperationalId,
		datainformasi.Id,
	)
	helper.PanicIfError(err)

	return datainformasi
}

func (repository *DataDanInformasiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, datainformasi domain.DataDanInformasi) {
	script := "delete from data_dan_informasi where id =?"
	_, err := tx.ExecContext(ctx, script, datainformasi.Id)
	helper.PanicIfError(err)
}
