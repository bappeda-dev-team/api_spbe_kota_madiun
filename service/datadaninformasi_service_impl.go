package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type DataDanInformasiServiceImpl struct {
	DataDanInformasiRepository    repository.DataDanInformasiRepository
	PohonkinerjaRepository        repository.PohonKinerjaRepository
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewDataDanInformasiServiceImpl(datadaninformasiRepository repository.DataDanInformasiRepository, pohonkinerjaRepository repository.PohonKinerjaRepository, referensiarsitekturRepository repository.ReferensiArsitekturRepository, DB *sql.DB, validate *validator.Validate) *DataDanInformasiServiceImpl {
	return &DataDanInformasiServiceImpl{
		DataDanInformasiRepository:    datadaninformasiRepository,
		PohonkinerjaRepository:        pohonkinerjaRepository,
		ReferensiArsitekturRepository: referensiarsitekturRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func (service *DataDanInformasiServiceImpl) FindByKodeOpd(ctx context.Context, kodeOPD string, tahun int) ([]web.DataDanInformasiRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataDanInformasiList, err := service.DataDanInformasiRepository.FindByKodeOpd(ctx, tx, kodeOPD, tahun)
	helper.PanicIfError(err)

	var responses []web.DataDanInformasiRespons
	for _, dataDanInformasi := range dataDanInformasiList {
		var strategicid, tacticalid, operational *web.DataDanInformasiPohonResponns

		if dataDanInformasi.StrategicId.Valid {
			strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(dataDanInformasi.StrategicId.Int32))
			helper.PanicIfError(err)
			strategicid = &web.DataDanInformasiPohonResponns{
				ID:         strategicData.ID,
				NamaPohon:  strategicData.NamaPohon,
				LevelPohon: strategicData.LevelPohon,
			}
		}

		if dataDanInformasi.TacticalId.Valid {
			tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(dataDanInformasi.TacticalId.Int32))
			helper.PanicIfError(err)
			tacticalid = &web.DataDanInformasiPohonResponns{
				ID:         tacticalData.ID,
				NamaPohon:  tacticalData.NamaPohon,
				LevelPohon: tacticalData.LevelPohon,
			}
		}

		if dataDanInformasi.OperationalId.Valid {
			operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(dataDanInformasi.OperationalId.Int32))
			helper.PanicIfError(err)
			operational = &web.DataDanInformasiPohonResponns{
				ID:         operationalData.ID,
				NamaPohon:  operationalData.NamaPohon,
				LevelPohon: operationalData.LevelPohon,
			}
		}

		var radLevel1, radLevel2, radLevel3, radLevel4 *web.DataDanInformasiReferensiArsitekturRespons
		if dataDanInformasi.RadLevel1id.Valid {
			radLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel1id.Int32))
			helper.PanicIfError(err)
			radLevel1 = &web.DataDanInformasiReferensiArsitekturRespons{
				Id:              radLevel1Data.IdReferensi,
				Kode_referensi:  radLevel1Data.Kode_referensi,
				Nama_referensi:  radLevel1Data.Nama_referensi,
				Level_referensi: radLevel1Data.Level_referensi,
			}
		}
		if dataDanInformasi.RadLevel2id.Valid {
			radLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel2id.Int32))
			helper.PanicIfError(err)
			radLevel2 = &web.DataDanInformasiReferensiArsitekturRespons{
				Id:              radLevel2Data.IdReferensi,
				Kode_referensi:  radLevel2Data.Kode_referensi,
				Nama_referensi:  radLevel2Data.Nama_referensi,
				Level_referensi: radLevel2Data.Level_referensi,
			}
		}
		if dataDanInformasi.RadLevel3id.Valid {
			radLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel3id.Int32))
			helper.PanicIfError(err)
			radLevel3 = &web.DataDanInformasiReferensiArsitekturRespons{
				Id:              radLevel3Data.IdReferensi,
				Kode_referensi:  radLevel3Data.Kode_referensi,
				Nama_referensi:  radLevel3Data.Nama_referensi,
				Level_referensi: radLevel3Data.Level_referensi,
			}
		}
		if dataDanInformasi.RadLevel4id.Valid {
			radLevel4Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel4id.Int32))
			helper.PanicIfError(err)
			radLevel4 = &web.DataDanInformasiReferensiArsitekturRespons{
				Id:              radLevel4Data.IdReferensi,
				Kode_referensi:  radLevel4Data.Kode_referensi,
				Nama_referensi:  radLevel4Data.Nama_referensi,
				Level_referensi: radLevel4Data.Level_referensi,
			}
		}

		response := web.DataDanInformasiRespons{
			Id:                     dataDanInformasi.Id,
			NamaData:               dataDanInformasi.NamaData,
			SifatData:              dataDanInformasi.SifatData,
			JenisData:              dataDanInformasi.JenisData,
			ProdusenData:           dataDanInformasi.ProdusenData,
			UraianData:             dataDanInformasi.UraianData,
			ValiditasData:          dataDanInformasi.ValiditasData,
			PjData:                 dataDanInformasi.PjData,
			KodeOPD:                dataDanInformasi.KodeOPD,
			InformasiTerkaitInput:  dataDanInformasi.InformasiTerkaitInput,
			InformasiTerkaitOutput: dataDanInformasi.InformasiTerkaitOutput,
			Interoprabilitas:       dataDanInformasi.Interoprabilitas,
			Keterangan:             nil,
			Tahun:                  dataDanInformasi.Tahun,
			CreatedAt:              dataDanInformasi.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:              dataDanInformasi.UpdatedAt.Format("2006-01-02 15:04:05"),
			RadLevel1id:            radLevel1,
			RadLevel2id:            radLevel2,
			RadLevel3id:            radLevel3,
			RadLevel4id:            radLevel4,
			StrategicId:            strategicid,
			TacticalId:             tacticalid,
			OperationalId:          operational,
		}

		if dataDanInformasi.Keterangan.Valid && dataDanInformasi.Keterangan.String != "" {
			response.Keterangan = &dataDanInformasi.Keterangan.String
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *DataDanInformasiServiceImpl) FindById(ctx context.Context, dataId int, kodeOPD string) (web.DataDanInformasiRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.DataDanInformasiRespons{}, err
	}
	defer helper.CommitOrRollback(tx)

	dataDanInformasi, err := service.DataDanInformasiRepository.FindById(ctx, tx, dataId)
	if err != nil {
		return web.DataDanInformasiRespons{}, err
	}

	if kodeOPD != "" && dataDanInformasi.KodeOPD != kodeOPD {
		return web.DataDanInformasiRespons{}, errors.New("data dan informasi tidak ditemukan untuk OPD ini")
	}

	var strategicid, tacticalid, operational *web.DataDanInformasiPohonResponns

	if dataDanInformasi.StrategicId.Valid {
		strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(dataDanInformasi.StrategicId.Int32))
		helper.PanicIfError(err)
		strategicid = &web.DataDanInformasiPohonResponns{
			ID:         strategicData.ID,
			NamaPohon:  strategicData.NamaPohon,
			LevelPohon: strategicData.LevelPohon,
		}
	}

	if dataDanInformasi.TacticalId.Valid {
		tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(dataDanInformasi.TacticalId.Int32))
		helper.PanicIfError(err)
		tacticalid = &web.DataDanInformasiPohonResponns{
			ID:         tacticalData.ID,
			NamaPohon:  tacticalData.NamaPohon,
			LevelPohon: tacticalData.LevelPohon,
		}
	}

	if dataDanInformasi.OperationalId.Valid {
		operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(dataDanInformasi.OperationalId.Int32))
		helper.PanicIfError(err)
		operational = &web.DataDanInformasiPohonResponns{
			ID:         operationalData.ID,
			NamaPohon:  operationalData.NamaPohon,
			LevelPohon: operationalData.LevelPohon,
		}
	}

	var radLevel1, radLevel2, radLevel3, radLevel4 *web.DataDanInformasiReferensiArsitekturRespons
	if dataDanInformasi.RadLevel1id.Valid {
		radLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel1id.Int32))
		helper.PanicIfError(err)
		radLevel1 = &web.DataDanInformasiReferensiArsitekturRespons{
			Id:              radLevel1Data.IdReferensi,
			Kode_referensi:  radLevel1Data.Kode_referensi,
			Nama_referensi:  radLevel1Data.Nama_referensi,
			Level_referensi: radLevel1Data.Level_referensi,
		}
	}
	if dataDanInformasi.RadLevel2id.Valid {
		radLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel2id.Int32))
		helper.PanicIfError(err)
		radLevel2 = &web.DataDanInformasiReferensiArsitekturRespons{
			Id:              radLevel2Data.IdReferensi,
			Kode_referensi:  radLevel2Data.Kode_referensi,
			Nama_referensi:  radLevel2Data.Nama_referensi,
			Level_referensi: radLevel2Data.Level_referensi,
		}
	}
	if dataDanInformasi.RadLevel3id.Valid {
		radLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel3id.Int32))
		helper.PanicIfError(err)
		radLevel3 = &web.DataDanInformasiReferensiArsitekturRespons{
			Id:              radLevel3Data.IdReferensi,
			Kode_referensi:  radLevel3Data.Kode_referensi,
			Nama_referensi:  radLevel3Data.Nama_referensi,
			Level_referensi: radLevel3Data.Level_referensi,
		}
	}
	if dataDanInformasi.RadLevel4id.Valid {
		radLevel4Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(dataDanInformasi.RadLevel4id.Int32))
		helper.PanicIfError(err)
		radLevel4 = &web.DataDanInformasiReferensiArsitekturRespons{
			Id:              radLevel4Data.IdReferensi,
			Kode_referensi:  radLevel4Data.Kode_referensi,
			Nama_referensi:  radLevel4Data.Nama_referensi,
			Level_referensi: radLevel4Data.Level_referensi,
		}
	}

	response := web.DataDanInformasiRespons{
		Id:                     dataDanInformasi.Id,
		NamaData:               dataDanInformasi.NamaData,
		SifatData:              dataDanInformasi.SifatData,
		JenisData:              dataDanInformasi.JenisData,
		ProdusenData:           dataDanInformasi.ProdusenData,
		UraianData:             dataDanInformasi.UraianData,
		ValiditasData:          dataDanInformasi.ValiditasData,
		PjData:                 dataDanInformasi.PjData,
		KodeOPD:                dataDanInformasi.KodeOPD,
		InformasiTerkaitInput:  dataDanInformasi.InformasiTerkaitInput,
		InformasiTerkaitOutput: dataDanInformasi.InformasiTerkaitOutput,
		Interoprabilitas:       dataDanInformasi.Interoprabilitas,
		Keterangan:             nil,
		Tahun:                  dataDanInformasi.Tahun,
		CreatedAt:              dataDanInformasi.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:              dataDanInformasi.UpdatedAt.Format("2006-01-02 15:04:05"),
		RadLevel1id:            radLevel1,
		RadLevel2id:            radLevel2,
		RadLevel3id:            radLevel3,
		RadLevel4id:            radLevel4,
		StrategicId:            strategicid,
		TacticalId:             tacticalid,
		OperationalId:          operational,
	}

	if dataDanInformasi.Keterangan.Valid && dataDanInformasi.Keterangan.String != "" {
		response.Keterangan = &dataDanInformasi.Keterangan.String
	}

	return response, nil
}

func (service *DataDanInformasiServiceImpl) Insert(ctx context.Context, request web.DataDanInformasiCreateRequest) web.DataDanInformasiRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	currentTime := time.Now()

	dataDanInformasi := domain.DataDanInformasi{
		NamaData:               request.NamaData,
		SifatData:              request.SifatData,
		JenisData:              request.JenisData,
		ProdusenData:           request.ProdusenData,
		UraianData:             request.UraianData,
		ValiditasData:          request.ValiditasData,
		PjData:                 request.PjData,
		KodeOPD:                request.KodeOPD,
		InformasiTerkaitInput:  request.InformasiTerkaitInput,
		InformasiTerkaitOutput: request.InformasiTerkaitOutput,
		Interoprabilitas:       request.Interoprabilitas,
		Keterangan: sql.NullString{
			String: "",
			Valid:  false,
		},
		Tahun:     request.Tahun,
		CreatedAt: currentTime,
		RadLevel1id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RadLevel2id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RadLevel3id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RadLevel4id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		StrategicId: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		TacticalId: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		OperationalId: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
	}

	if request.Keterangan != nil {
		dataDanInformasi.Keterangan = sql.NullString{
			String: *request.Keterangan,
			Valid:  *request.Keterangan != "",
		}
	}

	if request.RadLevel1id != nil {
		dataDanInformasi.RadLevel1id = sql.NullInt32{
			Int32: int32(*request.RadLevel1id),
			Valid: true,
		}
	}

	if request.RadLevel2id != nil {
		dataDanInformasi.RadLevel2id = sql.NullInt32{
			Int32: int32(*request.RadLevel2id),
			Valid: true,
		}
	}

	if request.RadLevel3id != nil {
		dataDanInformasi.RadLevel3id = sql.NullInt32{
			Int32: int32(*request.RadLevel3id),
			Valid: true,
		}
	}

	if request.RadLevel4id != nil {
		dataDanInformasi.RadLevel4id = sql.NullInt32{
			Int32: int32(*request.RadLevel4id),
			Valid: true,
		}
	}

	if request.StrategicId != nil {
		dataDanInformasi.StrategicId = sql.NullInt32{
			Int32: int32(*request.StrategicId),
			Valid: true,
		}
	}

	if request.TacticalId != nil {
		dataDanInformasi.TacticalId = sql.NullInt32{
			Int32: int32(*request.TacticalId),
			Valid: true,
		}
	}

	if request.OperationalId != nil {
		dataDanInformasi.OperationalId = sql.NullInt32{
			Int32: int32(*request.OperationalId),
			Valid: true,
		}
	}

	data := service.DataDanInformasiRepository.Insert(ctx, tx, dataDanInformasi)
	return helper.ToDataDanInformasiRespons(data)
}

func (service *DataDanInformasiServiceImpl) Update(ctx context.Context, request web.DataDanInformasiUpdateRequest) web.DataDanInformasiRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataInformasi, err := service.DataDanInformasiRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	dataInformasi.NamaData = request.NamaData
	dataInformasi.SifatData = request.SifatData
	dataInformasi.JenisData = request.JenisData
	dataInformasi.ProdusenData = request.ProdusenData
	dataInformasi.UraianData = request.UraianData
	dataInformasi.ValiditasData = request.ValiditasData
	dataInformasi.PjData = request.PjData
	dataInformasi.KodeOPD = request.KodeOPD
	dataInformasi.InformasiTerkaitInput = request.InformasiTerkaitInput
	dataInformasi.InformasiTerkaitOutput = request.InformasiTerkaitOutput
	dataInformasi.Interoprabilitas = request.Interoprabilitas
	if request.Keterangan != nil {
		dataInformasi.Keterangan = sql.NullString{
			String: *request.Keterangan,
			Valid:  true,
		}
	} else {
		dataInformasi.Keterangan = sql.NullString{
			Valid: false,
		}
	}
	dataInformasi.Tahun = request.Tahun
	dataInformasi.RadLevel1id = sql.NullInt32{Int32: int32(request.RadLevel1id), Valid: request.RadLevel1id != 0}
	dataInformasi.RadLevel2id = sql.NullInt32{Int32: int32(request.RadLevel2id), Valid: request.RadLevel2id != 0}
	dataInformasi.RadLevel3id = sql.NullInt32{Int32: int32(request.RadLevel3id), Valid: request.RadLevel3id != 0}
	dataInformasi.RadLevel4id = sql.NullInt32{Int32: int32(request.RadLevel4id), Valid: request.RadLevel4id != 0}
	dataInformasi.StrategicId = sql.NullInt32{Int32: int32(request.StrategicId), Valid: request.StrategicId != 0}
	dataInformasi.TacticalId = sql.NullInt32{Int32: int32(request.TacticalId), Valid: request.TacticalId != 0}
	dataInformasi.OperationalId = sql.NullInt32{Int32: int32(request.OperationalId), Valid: request.OperationalId != 0}

	dataInformasi = service.DataDanInformasiRepository.Update(ctx, tx, dataInformasi)

	return helper.ToDataDanInformasiRespons(dataInformasi)

}

func (service *DataDanInformasiServiceImpl) Delete(ctx context.Context, dataId int, kodeOPD string, role string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataInformasi, err := service.DataDanInformasiRepository.FindById(ctx, tx, dataId)
	helper.PanicIfError(err)

	if role == "admin_kota" {
		// Admin kota dapat menghapus semua data tanpa memeriksa kode OPD
		service.DataDanInformasiRepository.Delete(ctx, tx, dataInformasi)
	} else if role == "admin_opd" || role == "asn" {
		// Admin OPD dan ASN hanya dapat menghapus data berdasarkan kode OPD mereka
		if dataInformasi.KodeOPD != kodeOPD {
			return errors.New("data informasi tidak ditemukan untuk OPD ini")
		}
		service.DataDanInformasiRepository.Delete(ctx, tx, dataInformasi)
	} else {
		return errors.New("role tidak memiliki izin untuk menghapus data informasi")
	}
	return nil
}
