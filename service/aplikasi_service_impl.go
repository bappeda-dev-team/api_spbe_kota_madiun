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

type AplikasiServiceImpl struct {
	AplikasiRepository            repository.AplikasiRepository
	PohonkinerjaRepository        repository.PohonKinerjaRepository
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewAplikasiServiceImpl(aplikasiRepository repository.AplikasiRepository, pohonkinerjaRepository repository.PohonKinerjaRepository, referensiarsitektuRepository repository.ReferensiArsitekturRepository, DB *sql.DB, validate *validator.Validate) *AplikasiServiceImpl {
	return &AplikasiServiceImpl{
		AplikasiRepository:            aplikasiRepository,
		PohonkinerjaRepository:        pohonkinerjaRepository,
		ReferensiArsitekturRepository: referensiarsitektuRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func (service *AplikasiServiceImpl) FindByKodeOpd(ctx context.Context, kodeOPD string, tahun int) ([]web.AplikasiRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aplikasiList, err := service.AplikasiRepository.FindByKodeOpd(ctx, tx, kodeOPD, tahun)
	helper.PanicIfError(err)

	var responses []web.AplikasiRespons
	for _, aplikasi := range aplikasiList {
		var strategicid, tacticalid, operational *web.AplikasiPohonRespons

		if aplikasi.StrategicId.Valid {
			strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.StrategicId.Int32))
			helper.PanicIfError(err)
			strategicid = &web.AplikasiPohonRespons{
				ID:         strategicData.ID,
				NamaPohon:  strategicData.NamaPohon,
				LevelPohon: strategicData.LevelPohon,
			}
		}

		if aplikasi.TacticalId.Valid {
			tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.TacticalId.Int32))
			helper.PanicIfError(err)
			tacticalid = &web.AplikasiPohonRespons{
				ID:         tacticalData.ID,
				NamaPohon:  tacticalData.NamaPohon,
				LevelPohon: tacticalData.LevelPohon,
			}
		}

		if aplikasi.OperationalId.Valid {
			operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.OperationalId.Int32))
			helper.PanicIfError(err)
			operational = &web.AplikasiPohonRespons{
				ID:         operationalData.ID,
				NamaPohon:  operationalData.NamaPohon,
				LevelPohon: operationalData.LevelPohon,
			}
		}

		var raaLevel1, raaLevel2, raaLevel3 *web.AplikasiReferensiArsitekturRespons
		if aplikasi.RaaLevel1id.Valid {
			raaLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(aplikasi.RaaLevel1id.Int32))
			helper.PanicIfError(err)
			raaLevel1 = &web.AplikasiReferensiArsitekturRespons{
				Id:              raaLevel1Data.IdReferensi,
				Kode_referensi:  raaLevel1Data.Kode_referensi,
				Nama_referensi:  raaLevel1Data.Nama_referensi,
				Level_referensi: raaLevel1Data.Level_referensi,
			}
		}
		if aplikasi.RaaLevel2id.Valid {
			raaLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(aplikasi.RaaLevel2id.Int32))
			helper.PanicIfError(err)
			raaLevel2 = &web.AplikasiReferensiArsitekturRespons{
				Id:              raaLevel2Data.IdReferensi,
				Kode_referensi:  raaLevel2Data.Kode_referensi,
				Nama_referensi:  raaLevel2Data.Nama_referensi,
				Level_referensi: raaLevel2Data.Level_referensi,
			}
		}
		if aplikasi.RaaLevel3id.Valid {
			raaLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(aplikasi.RaaLevel3id.Int32))
			helper.PanicIfError(err)
			raaLevel3 = &web.AplikasiReferensiArsitekturRespons{
				Id:              raaLevel3Data.IdReferensi,
				Kode_referensi:  raaLevel3Data.Kode_referensi,
				Nama_referensi:  raaLevel3Data.Nama_referensi,
				Level_referensi: raaLevel3Data.Level_referensi,
			}
		}

		response := web.AplikasiRespons{
			Id:                     aplikasi.Id,
			NamaAplikasi:           aplikasi.NamaAplikasi,
			FungsiAplikasi:         aplikasi.FungsiAplikasi,
			JenisAplikasi:          aplikasi.JenisAplikasi,
			ProdusenAplikasi:       aplikasi.ProdusenAplikasi,
			PjAplikasi:             aplikasi.PjAplikasi,
			KodeOPD:                aplikasi.KodeOPD,
			InformasiTerkaitInput:  aplikasi.InformasiTerkaitInput,
			InformasiTerkaitOutput: aplikasi.InformasiTerkaitOutput,
			Interoprabilitas:       aplikasi.Interoprabilitas,
			Keterangan:             nil,
			Tahun:                  aplikasi.Tahun,
			CreatedAt:              aplikasi.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:              aplikasi.UpdatedAt.Format("2006-01-02 15:04:05"),
			RaaLevel1id:            raaLevel1,
			RaaLevel2id:            raaLevel2,
			RaaLevel3id:            raaLevel3,
			StrategicId:            strategicid,
			TacticalId:             tacticalid,
			OperationalId:          operational,
		}
		if aplikasi.Keterangan.Valid && aplikasi.Keterangan.String != "" {
			response.Keterangan = &aplikasi.Keterangan.String
		}

		responses = append(responses, response)
	}

	return responses, nil

}
func (service *AplikasiServiceImpl) FindById(ctx context.Context, aplikasiId int, kodeOPD string) (web.AplikasiRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.AplikasiRespons{}, err
	}
	defer helper.CommitOrRollback(tx)

	aplikasi, err := service.AplikasiRepository.FindById(ctx, tx, aplikasiId)
	if err != nil {
		return web.AplikasiRespons{}, err
	}

	if kodeOPD != "" && aplikasi.KodeOPD != kodeOPD {
		return web.AplikasiRespons{}, errors.New("aplikasi tidak ditemukan untuk OPD ini")
	}

	var strategicid, tacticalid, operational *web.AplikasiPohonRespons

	if aplikasi.StrategicId.Valid {
		strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.StrategicId.Int32))
		helper.PanicIfError(err)
		strategicid = &web.AplikasiPohonRespons{
			ID:         strategicData.ID,
			NamaPohon:  strategicData.NamaPohon,
			LevelPohon: strategicData.LevelPohon,
		}
	}

	if aplikasi.TacticalId.Valid {
		tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.TacticalId.Int32))
		helper.PanicIfError(err)
		tacticalid = &web.AplikasiPohonRespons{
			ID:         tacticalData.ID,
			NamaPohon:  tacticalData.NamaPohon,
			LevelPohon: tacticalData.LevelPohon,
		}
	}

	if aplikasi.OperationalId.Valid {
		operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.OperationalId.Int32))
		helper.PanicIfError(err)
		operational = &web.AplikasiPohonRespons{
			ID:         operationalData.ID,
			NamaPohon:  operationalData.NamaPohon,
			LevelPohon: operationalData.LevelPohon,
		}
	}

	var raaLevel1, raaLevel2, raaLevel3 *web.AplikasiReferensiArsitekturRespons
	if aplikasi.RaaLevel1id.Valid {
		raaLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(aplikasi.RaaLevel1id.Int32))
		helper.PanicIfError(err)
		raaLevel1 = &web.AplikasiReferensiArsitekturRespons{
			Id:              raaLevel1Data.IdReferensi,
			Kode_referensi:  raaLevel1Data.Kode_referensi,
			Nama_referensi:  raaLevel1Data.Nama_referensi,
			Level_referensi: raaLevel1Data.Level_referensi,
		}
	}
	if aplikasi.RaaLevel2id.Valid {
		raaLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(aplikasi.RaaLevel2id.Int32))
		helper.PanicIfError(err)
		raaLevel2 = &web.AplikasiReferensiArsitekturRespons{
			Id:              raaLevel2Data.IdReferensi,
			Kode_referensi:  raaLevel2Data.Kode_referensi,
			Nama_referensi:  raaLevel2Data.Nama_referensi,
			Level_referensi: raaLevel2Data.Level_referensi,
		}
	}
	if aplikasi.RaaLevel3id.Valid {
		raaLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(aplikasi.RaaLevel3id.Int32))
		helper.PanicIfError(err)
		raaLevel3 = &web.AplikasiReferensiArsitekturRespons{
			Id:              raaLevel3Data.IdReferensi,
			Kode_referensi:  raaLevel3Data.Kode_referensi,
			Nama_referensi:  raaLevel3Data.Nama_referensi,
			Level_referensi: raaLevel3Data.Level_referensi,
		}
	}

	response := web.AplikasiRespons{
		Id:                     aplikasi.Id,
		NamaAplikasi:           aplikasi.NamaAplikasi,
		FungsiAplikasi:         aplikasi.FungsiAplikasi,
		JenisAplikasi:          aplikasi.JenisAplikasi,
		ProdusenAplikasi:       aplikasi.ProdusenAplikasi,
		PjAplikasi:             aplikasi.PjAplikasi,
		KodeOPD:                aplikasi.KodeOPD,
		InformasiTerkaitInput:  aplikasi.InformasiTerkaitInput,
		InformasiTerkaitOutput: aplikasi.InformasiTerkaitOutput,
		Interoprabilitas:       aplikasi.Interoprabilitas,
		Keterangan:             nil,
		Tahun:                  aplikasi.Tahun,
		CreatedAt:              aplikasi.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:              aplikasi.UpdatedAt.Format("2006-01-02 15:04:05"),
		RaaLevel1id:            raaLevel1,
		RaaLevel2id:            raaLevel2,
		RaaLevel3id:            raaLevel3,
		StrategicId:            strategicid,
		TacticalId:             tacticalid,
		OperationalId:          operational,
	}
	if aplikasi.Keterangan.Valid && aplikasi.Keterangan.String != "" {
		response.Keterangan = &aplikasi.Keterangan.String
	}

	return response, nil
}

func (service *AplikasiServiceImpl) Insert(ctx context.Context, request web.AplikasiCreateRequest) web.AplikasiRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	currentTime := time.Now()

	aplikasi := domain.Aplikasi{
		NamaAplikasi:           request.NamaAplikasi,
		FungsiAplikasi:         request.FungsiAplikasi,
		JenisAplikasi:          request.JenisAplikasi,
		ProdusenAplikasi:       request.ProdusenAplikasi,
		PjAplikasi:             request.PjAplikasi,
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
		RaaLevel1id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RaaLevel2id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RaaLevel3id: sql.NullInt32{
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
		aplikasi.Keterangan = sql.NullString{
			String: *request.Keterangan,
			Valid:  *request.Keterangan != "",
		}
	}

	if request.RaaLevel1id != nil {
		aplikasi.RaaLevel1id = sql.NullInt32{
			Int32: int32(*request.RaaLevel1id),
			Valid: true,
		}
	}

	if request.RaaLevel2id != nil {
		aplikasi.RaaLevel2id = sql.NullInt32{
			Int32: int32(*request.RaaLevel2id),
			Valid: true,
		}
	}

	if request.RaaLevel3id != nil {
		aplikasi.RaaLevel3id = sql.NullInt32{
			Int32: int32(*request.RaaLevel3id),
			Valid: true,
		}
	}

	if request.StrategicId != nil {
		aplikasi.StrategicId = sql.NullInt32{
			Int32: int32(*request.StrategicId),
			Valid: true,
		}
	}

	if request.TacticalId != nil {
		aplikasi.TacticalId = sql.NullInt32{
			Int32: int32(*request.TacticalId),
			Valid: true,
		}
	}

	if request.OperationalId != nil {
		aplikasi.OperationalId = sql.NullInt32{
			Int32: int32(*request.OperationalId),
			Valid: true,
		}
	}

	// Jika TacticalId terisi, ambil nama pohonnya dan isi FungsiAplikasi
	if aplikasi.TacticalId.Valid {
		tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.TacticalId.Int32))
		helper.PanicIfError(err)
		aplikasi.FungsiAplikasi = tacticalData.NamaPohon
	}

	aplikasData := service.AplikasiRepository.Insert(ctx, tx, aplikasi)
	return helper.ToAplikasiRespons(aplikasData)
}

func (service *AplikasiServiceImpl) Update(ctx context.Context, request web.AplikasiUpdateRequest) web.AplikasiRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aplikasi, err := service.AplikasiRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	aplikasi.NamaAplikasi = request.NamaAplikasi
	aplikasi.FungsiAplikasi = request.FungsiAplikasi
	aplikasi.JenisAplikasi = request.JenisAplikasi
	aplikasi.ProdusenAplikasi = request.ProdusenAplikasi
	aplikasi.PjAplikasi = request.PjAplikasi
	aplikasi.KodeOPD = request.KodeOPD
	aplikasi.InformasiTerkaitInput = request.InformasiTerkaitInput
	aplikasi.InformasiTerkaitOutput = request.InformasiTerkaitOutput
	aplikasi.Interoprabilitas = request.Interoprabilitas
	aplikasi.Tahun = request.Tahun
	aplikasi.Keterangan = sql.NullString{String: string(request.Keterangan), Valid: request.Keterangan != ""}
	aplikasi.RaaLevel1id = sql.NullInt32{Int32: int32(request.RaaLevel1id), Valid: request.RaaLevel1id != 0}
	aplikasi.RaaLevel2id = sql.NullInt32{Int32: int32(request.RaaLevel2id), Valid: request.RaaLevel2id != 0}
	aplikasi.RaaLevel3id = sql.NullInt32{Int32: int32(request.RaaLevel3id), Valid: request.RaaLevel3id != 0}
	aplikasi.StrategicId = sql.NullInt32{Int32: int32(request.StrategicId), Valid: request.StrategicId != 0}
	aplikasi.TacticalId = sql.NullInt32{Int32: int32(request.TacticalId), Valid: request.TacticalId != 0}
	aplikasi.OperationalId = sql.NullInt32{Int32: int32(request.OperationalId), Valid: request.OperationalId != 0}

	if aplikasi.TacticalId.Valid {
		tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(aplikasi.TacticalId.Int32))
		helper.PanicIfError(err)
		aplikasi.FungsiAplikasi = tacticalData.NamaPohon
	}

	aplikasi = service.AplikasiRepository.Update(ctx, tx, aplikasi)
	return helper.ToAplikasiRespons(aplikasi)
}

func (service *AplikasiServiceImpl) Delete(ctx context.Context, aplikasiId int, kodeOPD string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aplikasi, err := service.AplikasiRepository.FindById(ctx, tx, aplikasiId)
	helper.PanicIfError(err)

	if aplikasi.KodeOPD != kodeOPD {
		panic(errors.New("aplikasi tidak ditemukan untuk OPD ini"))
	}

	service.AplikasiRepository.Delete(ctx, tx, aplikasi)
	return nil
}
