package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

type LayananSpbeServiceImpl struct {
	LayananspbeRepository         repository.LayananSPBERepository
	PohonkinerjaRepository        repository.PohonKinerjaRepository
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewLayananSpbeService(layananspbeRepository repository.LayananSPBERepository, pohonkinerjaRepository repository.PohonKinerjaRepository, referensiarsitekturRepository repository.ReferensiArsitekturRepository, DB *sql.DB, validate *validator.Validate) LayananSpbeService {
	return &LayananSpbeServiceImpl{
		LayananspbeRepository:         layananspbeRepository,
		PohonkinerjaRepository:        pohonkinerjaRepository,
		ReferensiArsitekturRepository: referensiarsitekturRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func (service *LayananSpbeServiceImpl) FindByKodeOpd(ctx context.Context, kodeOPD string, tahun int) ([]web.LayananSpbeRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	layananspbeList, err := service.LayananspbeRepository.FindByKodeOpd(ctx, tx, kodeOPD, tahun)
	helper.PanicIfError(err)
	helper.PanicIfError(err)

	var responses []web.LayananSpbeRespons
	for _, layananspbe := range layananspbeList {
		var strategicid, tacticalid, operational, tujuanlayanan *web.LayananspbePohonRespons

		if layananspbe.TujuanLayananId.Valid {
			tujuanlayananData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.TujuanLayananId.Int32))
			helper.PanicIfError(err)
			tujuanlayanan = &web.LayananspbePohonRespons{
				ID:         tujuanlayananData.ID,
				NamaPohon:  tujuanlayananData.NamaPohon,
				LevelPohon: tujuanlayananData.LevelPohon,
			}
		}

		if layananspbe.StrategicId.Valid {
			strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.StrategicId.Int32))
			helper.PanicIfError(err)
			strategicid = &web.LayananspbePohonRespons{
				ID:         strategicData.ID,
				NamaPohon:  strategicData.NamaPohon,
				LevelPohon: strategicData.LevelPohon,
			}
		}

		if layananspbe.TacticalId.Valid {
			tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.TacticalId.Int32))
			helper.PanicIfError(err)
			tacticalid = &web.LayananspbePohonRespons{
				ID:         tacticalData.ID,
				NamaPohon:  tacticalData.NamaPohon,
				LevelPohon: tacticalData.LevelPohon,
			}
		}

		if layananspbe.OperationalId.Valid {
			operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.OperationalId.Int32))
			helper.PanicIfError(err)
			operational = &web.LayananspbePohonRespons{
				ID:         operationalData.ID,
				NamaPohon:  operationalData.NamaPohon,
				LevelPohon: operationalData.LevelPohon,
			}
		}

		var ralLevel1, ralLevel2, ralLevel3, ralLevel4 *web.LayananSpbeReferensiArsitekturRespons
		if layananspbe.RalLevel1id.Valid {
			ralLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel1id.Int32))
			helper.PanicIfError(err)
			ralLevel1 = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel1Data.IdReferensi,
				Kode_referensi:  ralLevel1Data.Kode_referensi,
				Nama_referensi:  ralLevel1Data.Nama_referensi,
				Level_referensi: ralLevel1Data.Level_referensi,
			}
		}
		if layananspbe.RalLevel2id.Valid {
			ralLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel2id.Int32))
			helper.PanicIfError(err)
			ralLevel2 = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel2Data.IdReferensi,
				Kode_referensi:  ralLevel2Data.Kode_referensi,
				Nama_referensi:  ralLevel2Data.Nama_referensi,
				Level_referensi: ralLevel2Data.Level_referensi,
			}
		}
		if layananspbe.RalLevel3id.Valid {
			ralLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel3id.Int32))
			helper.PanicIfError(err)
			ralLevel3 = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel3Data.IdReferensi,
				Kode_referensi:  ralLevel3Data.Kode_referensi,
				Nama_referensi:  ralLevel3Data.Nama_referensi,
				Level_referensi: ralLevel3Data.Level_referensi,
			}
		}
		if layananspbe.RalLevel4id.Valid {
			ralLevel4Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel4id.Int32))
			helper.PanicIfError(err)
			ralLevel4 = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel4Data.IdReferensi,
				Kode_referensi:  ralLevel4Data.Kode_referensi,
				Nama_referensi:  ralLevel4Data.Nama_referensi,
				Level_referensi: ralLevel4Data.Level_referensi,
			}
		}

		response := web.LayananSpbeRespons{
			Id:                layananspbe.Id,
			NamaLayanan:       layananspbe.NamaLayanan,
			KodeLayanan:       layananspbe.KodeLayanan,
			TujuanLayananId:   tujuanlayanan,
			FungsiLayanan:     layananspbe.FungsiLayanan,
			Tahun:             layananspbe.Tahun,
			KodeOPD:           layananspbe.KodeOPD,
			KementrianTerkait: layananspbe.KementrianTerkait,
			MetodeLayanan:     layananspbe.MetodeLayanan,
			CreatedAt:         layananspbe.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:         layananspbe.UpdatedAt.Format("2006-01-02 15:04:05"),
			RalLevel1id:       ralLevel1,
			RalLevel2id:       ralLevel2,
			RalLevel3id:       ralLevel3,
			RalLevel4id:       ralLevel4,
			StrategicId:       strategicid,
			TacticalId:        tacticalid,
			OperationalId:     operational,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *LayananSpbeServiceImpl) FindById(ctx context.Context, layananspbeId int) (web.LayananSpbeRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.LayananSpbeRespons{}, err
	}
	defer helper.CommitOrRollback(tx)

	layananspbe, err := service.LayananspbeRepository.FindById(ctx, tx, layananspbeId)
	if err != nil {
		return web.LayananSpbeRespons{}, err
	}

	var tujuanlayanan, strategicid, tacticalid, operational *web.LayananspbePohonRespons

	if layananspbe.TujuanLayananId.Valid {
		tujuanlayananData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.TujuanLayananId.Int32))
		helper.PanicIfError(err)
		tujuanlayanan = &web.LayananspbePohonRespons{
			ID:         tujuanlayananData.ID,
			NamaPohon:  tujuanlayananData.NamaPohon,
			LevelPohon: tujuanlayananData.LevelPohon,
		}
	}

	if layananspbe.StrategicId.Valid {
		strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.StrategicId.Int32))
		helper.PanicIfError(err)
		strategicid = &web.LayananspbePohonRespons{
			ID:         strategicData.ID,
			NamaPohon:  strategicData.NamaPohon,
			LevelPohon: strategicData.LevelPohon,
		}
	}

	if layananspbe.TacticalId.Valid {
		tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.TacticalId.Int32))
		helper.PanicIfError(err)
		tacticalid = &web.LayananspbePohonRespons{
			ID:         tacticalData.ID,
			NamaPohon:  tacticalData.NamaPohon,
			LevelPohon: tacticalData.LevelPohon,
		}
	}

	if layananspbe.OperationalId.Valid {
		operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananspbe.OperationalId.Int32))
		helper.PanicIfError(err)
		operational = &web.LayananspbePohonRespons{
			ID:         operationalData.ID,
			NamaPohon:  operationalData.NamaPohon,
			LevelPohon: operationalData.LevelPohon,
		}
	}

	var ralLevel1, ralLevel2, ralLevel3, ralLevel4 *web.LayananSpbeReferensiArsitekturRespons
	if layananspbe.RalLevel1id.Valid {
		ralLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel1id.Int32))
		helper.PanicIfError(err)
		ralLevel1 = &web.LayananSpbeReferensiArsitekturRespons{
			Id:              ralLevel1Data.IdReferensi,
			Kode_referensi:  ralLevel1Data.Kode_referensi,
			Nama_referensi:  ralLevel1Data.Nama_referensi,
			Level_referensi: ralLevel1Data.Level_referensi,
		}
	}
	if layananspbe.RalLevel2id.Valid {
		ralLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel2id.Int32))
		helper.PanicIfError(err)
		ralLevel2 = &web.LayananSpbeReferensiArsitekturRespons{
			Id:              ralLevel2Data.IdReferensi,
			Kode_referensi:  ralLevel2Data.Kode_referensi,
			Nama_referensi:  ralLevel2Data.Nama_referensi,
			Level_referensi: ralLevel2Data.Level_referensi,
		}
	}
	if layananspbe.RalLevel3id.Valid {
		ralLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel3id.Int32))
		helper.PanicIfError(err)
		ralLevel3 = &web.LayananSpbeReferensiArsitekturRespons{
			Id:              ralLevel3Data.IdReferensi,
			Kode_referensi:  ralLevel3Data.Kode_referensi,
			Nama_referensi:  ralLevel3Data.Nama_referensi,
			Level_referensi: ralLevel3Data.Level_referensi,
		}
	}
	if layananspbe.RalLevel4id.Valid {
		ralLevel4Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananspbe.RalLevel4id.Int32))
		helper.PanicIfError(err)
		ralLevel4 = &web.LayananSpbeReferensiArsitekturRespons{
			Id:              ralLevel4Data.IdReferensi,
			Kode_referensi:  ralLevel4Data.Kode_referensi,
			Nama_referensi:  ralLevel4Data.Nama_referensi,
			Level_referensi: ralLevel4Data.Level_referensi,
		}
	}

	response := web.LayananSpbeRespons{
		Id:                layananspbe.Id,
		NamaLayanan:       layananspbe.NamaLayanan,
		KodeLayanan:       layananspbe.KodeLayanan,
		TujuanLayananId:   tujuanlayanan,
		FungsiLayanan:     layananspbe.FungsiLayanan,
		Tahun:             layananspbe.Tahun,
		KodeOPD:           layananspbe.KodeOPD,
		KementrianTerkait: layananspbe.KementrianTerkait,
		MetodeLayanan:     layananspbe.MetodeLayanan,
		CreatedAt:         layananspbe.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:         layananspbe.UpdatedAt.Format("2006-01-02 15:04:05"),
		RalLevel1id:       ralLevel1,
		RalLevel2id:       ralLevel2,
		RalLevel3id:       ralLevel3,
		RalLevel4id:       ralLevel4,
		StrategicId:       strategicid,
		TacticalId:        tacticalid,
		OperationalId:     operational,
	}

	return response, nil
}
func (service *LayananSpbeServiceImpl) Insert(ctx context.Context, request web.LayananSpbeCreateRequest) web.LayananSpbeRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	currentTime := time.Now()
	kodeLayanan := helper.GenerateRandomKode()

	layananspbe := domain.LayananSPBE{
		NamaLayanan: request.NamaLayanan,
		KodeLayanan: kodeLayanan,
		TujuanLayananId: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		FungsiLayanan:     request.FungsiLayanan,
		Tahun:             request.Tahun,
		KodeOPD:           request.KodeOPD,
		KementrianTerkait: request.KementrianTerkait,
		MetodeLayanan:     request.MetodeLayanan,
		CreatedAt:         currentTime,
		RalLevel1id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RalLevel2id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RalLevel3id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		}, RalLevel4id: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		}, StrategicId: sql.NullInt32{
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

	if request.TujuanLayananId != nil {
		layananspbe.TujuanLayananId = sql.NullInt32{
			Int32: int32(*request.TujuanLayananId),
			Valid: true,
		}
	}

	if request.RalLevel1id != nil {
		layananspbe.RalLevel1id = sql.NullInt32{
			Int32: int32(*request.RalLevel1id),
			Valid: true,
		}
	}

	if request.RalLevel2id != nil {
		layananspbe.RalLevel2id = sql.NullInt32{
			Int32: int32(*request.RalLevel2id),
			Valid: true,
		}
	}

	if request.RalLevel3id != nil {
		layananspbe.RalLevel3id = sql.NullInt32{
			Int32: int32(*request.RalLevel3id),
			Valid: true,
		}
	}

	if request.RalLevel4id != nil {
		layananspbe.RalLevel4id = sql.NullInt32{
			Int32: int32(*request.RalLevel4id),
			Valid: true,
		}
	}

	if request.StrategicId != nil {
		layananspbe.StrategicId = sql.NullInt32{
			Int32: int32(*request.StrategicId),
			Valid: true,
		}
	}

	if request.TacticalId != nil {
		layananspbe.TacticalId = sql.NullInt32{
			Int32: int32(*request.TacticalId),
			Valid: true,
		}
	}

	if request.OperationalId != nil {
		layananspbe.OperationalId = sql.NullInt32{
			Int32: int32(*request.OperationalId),
			Valid: true,
		}
	}

	layananSpbe := service.LayananspbeRepository.Insert(ctx, tx, layananspbe)
	return helper.ToLayananSpbeRespons(layananSpbe)

}
func (service *LayananSpbeServiceImpl) Update(ctx context.Context, request web.LayananSpbeUpdateRequest) web.LayananSpbeRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	layananSpbe, err := service.LayananspbeRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	layananSpbe.NamaLayanan = request.NamaLayanan
	if layananSpbe.KodeLayanan == "" {
		layananSpbe.KodeLayanan = helper.GenerateRandomKode()
	}
	layananSpbe.TujuanLayananId = sql.NullInt32{Int32: int32(request.TujuanLayananId), Valid: request.TujuanLayananId != 0}
	layananSpbe.FungsiLayanan = request.FungsiLayanan
	layananSpbe.Tahun = request.Tahun
	layananSpbe.KodeOPD = request.KodeOPD
	layananSpbe.KementrianTerkait = request.KementrianTerkait
	layananSpbe.MetodeLayanan = request.MetodeLayanan
	layananSpbe.RalLevel1id = sql.NullInt32{Int32: int32(request.RalLevel1id), Valid: request.RalLevel1id != 0}
	layananSpbe.RalLevel2id = sql.NullInt32{Int32: int32(request.RalLevel2id), Valid: request.RalLevel2id != 0}
	layananSpbe.RalLevel3id = sql.NullInt32{Int32: int32(request.RalLevel3id), Valid: request.RalLevel3id != 0}
	layananSpbe.RalLevel4id = sql.NullInt32{Int32: int32(request.RalLevel4id), Valid: request.RalLevel4id != 0}
	layananSpbe.StrategicId = sql.NullInt32{Int32: int32(request.StrategicId), Valid: request.StrategicId != 0}
	layananSpbe.TacticalId = sql.NullInt32{Int32: int32(request.TacticalId), Valid: request.TacticalId != 0}
	layananSpbe.OperationalId = sql.NullInt32{Int32: int32(request.OperationalId), Valid: request.OperationalId != 0}

	layananSpbe = service.LayananspbeRepository.Update(ctx, tx, layananSpbe)
	return helper.ToLayananSpbeRespons(layananSpbe)
}

func (service *LayananSpbeServiceImpl) Delete(ctx context.Context, layananspbeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	layananSpbe, err := service.LayananspbeRepository.FindById(ctx, tx, layananspbeId)
	helper.PanicIfError(err)

	service.LayananspbeRepository.Delete(ctx, tx, layananSpbe)
}

func (service *LayananSpbeServiceImpl) FindByNull(ctx context.Context) ([]web.LayananSpbeRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	layananSpbe, err := service.LayananspbeRepository.FindByNull(ctx, tx)
	helper.PanicIfError(err)

	var responseList []web.LayananSpbeRespons
	for _, layananSpbe := range layananSpbe {
		response := web.LayananSpbeRespons{
			Id:                layananSpbe.Id,
			NamaLayanan:       layananSpbe.NamaLayanan,
			KodeLayanan:       layananSpbe.KodeLayanan,
			FungsiLayanan:     layananSpbe.FungsiLayanan,
			Tahun:             layananSpbe.Tahun,
			KodeOPD:           layananSpbe.KodeOPD,
			KementrianTerkait: layananSpbe.KementrianTerkait,
			MetodeLayanan:     layananSpbe.MetodeLayanan,
		}

		if layananSpbe.TujuanLayananId.Valid {
			tujuanlayananData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananSpbe.TujuanLayananId.Int32))
			helper.PanicIfError(err)
			response.TujuanLayananId = &web.LayananspbePohonRespons{
				ID:         tujuanlayananData.ID,
				NamaPohon:  tujuanlayananData.NamaPohon,
				LevelPohon: tujuanlayananData.LevelPohon,
			}
		}

		if layananSpbe.RalLevel1id.Valid {
			ralLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananSpbe.RalLevel1id.Int32))
			helper.PanicIfError(err)
			response.RalLevel1id = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel1Data.IdReferensi,
				Kode_referensi:  ralLevel1Data.Kode_referensi,
				Nama_referensi:  ralLevel1Data.Nama_referensi,
				Level_referensi: ralLevel1Data.Level_referensi,
			}
		}

		if layananSpbe.RalLevel2id.Valid {
			ralLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananSpbe.RalLevel2id.Int32))
			helper.PanicIfError(err)
			response.RalLevel2id = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel2Data.IdReferensi,
				Kode_referensi:  ralLevel2Data.Kode_referensi,
				Nama_referensi:  ralLevel2Data.Nama_referensi,
				Level_referensi: ralLevel2Data.Level_referensi,
			}
		}

		if layananSpbe.RalLevel3id.Valid {
			ralLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananSpbe.RalLevel3id.Int32))
			helper.PanicIfError(err)
			response.RalLevel3id = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel3Data.IdReferensi,
				Kode_referensi:  ralLevel3Data.Kode_referensi,
				Nama_referensi:  ralLevel3Data.Nama_referensi,
				Level_referensi: ralLevel3Data.Level_referensi,
			}
		}

		if layananSpbe.RalLevel4id.Valid {
			ralLevel4Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(layananSpbe.RalLevel4id.Int32))
			helper.PanicIfError(err)
			response.RalLevel4id = &web.LayananSpbeReferensiArsitekturRespons{
				Id:              ralLevel4Data.IdReferensi,
				Kode_referensi:  ralLevel4Data.Kode_referensi,
				Nama_referensi:  ralLevel4Data.Nama_referensi,
				Level_referensi: ralLevel4Data.Level_referensi,
			}
		}

		if layananSpbe.StrategicId.Valid {
			strategicData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananSpbe.StrategicId.Int32))
			helper.PanicIfError(err)
			response.StrategicId = &web.LayananspbePohonRespons{
				ID:         strategicData.ID,
				NamaPohon:  strategicData.NamaPohon,
				LevelPohon: strategicData.LevelPohon,
			}
		}

		if layananSpbe.TacticalId.Valid {
			tacticalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananSpbe.TacticalId.Int32))
			helper.PanicIfError(err)
			response.TacticalId = &web.LayananspbePohonRespons{
				ID:         tacticalData.ID,
				NamaPohon:  tacticalData.NamaPohon,
				LevelPohon: tacticalData.LevelPohon,
			}
		}

		if layananSpbe.OperationalId.Valid {
			operationalData, err := service.PohonkinerjaRepository.FindById(ctx, tx, int(layananSpbe.OperationalId.Int32))
			helper.PanicIfError(err)
			response.OperationalId = &web.LayananspbePohonRespons{
				ID:         operationalData.ID,
				NamaPohon:  operationalData.NamaPohon,
				LevelPohon: operationalData.LevelPohon,
			}
		}

		responseList = append(responseList, response)
	}

	return responseList, nil
}
