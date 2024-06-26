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

type ProsesBisnisServiceImpl struct {
	ProsesBisnisRepository        repository.ProsesBisnisRepository
	SasaranKotaRepository         repository.SasaranKotaRepository
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	PohonKinerjaRepository        repository.PohonKinerjaRepository
	BidangUrusanRepository        repository.BidangUrusanRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewProsesBisnisService(referensiarsitekturRepository repository.ReferensiArsitekturRepository, sasarankotaRepository repository.SasaranKotaRepository, bidangUrusanRepository repository.BidangUrusanRepository, prosesbisnisRepository repository.ProsesBisnisRepository, pohonkinerjaRepository repository.PohonKinerjaRepository, DB *sql.DB, validate *validator.Validate) ProsesBisnisService {
	return &ProsesBisnisServiceImpl{
		ProsesBisnisRepository:        prosesbisnisRepository,
		SasaranKotaRepository:         sasarankotaRepository,
		ReferensiArsitekturRepository: referensiarsitekturRepository,
		BidangUrusanRepository:        bidangUrusanRepository,
		PohonKinerjaRepository:        pohonkinerjaRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func (service *ProsesBisnisServiceImpl) GetProsesBisnis(ctx context.Context, kodeOPD string, tahun int) ([]web.ProsesBisnisRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	prosesBisnisList, err := service.ProsesBisnisRepository.FindByKodeOpd(ctx, tx, kodeOPD, tahun)
	helper.PanicIfError(err)

	var responses []web.ProsesBisnisRespons
	for _, prosesBisnis := range prosesBisnisList {

		var sasaranKota *web.ProsbisSasaranKotaRespons
		if prosesBisnis.SasaranKotaId.Valid {
			sasarankotaData, err := service.SasaranKotaRepository.FindById(ctx, tx, int(prosesBisnis.SasaranKotaId.Int32))
			helper.PanicIfError(err)
			sasaranKota = &web.ProsbisSasaranKotaRespons{
				ID:      sasarankotaData.ID,
				Sasaran: sasarankotaData.Sasaran,
			}
		}

		var bidangUrusan *web.ProsBisBidangUrusanRespons

		if prosesBisnis.BidangUrusanId.Valid {
			bidangurusanData, err := service.BidangUrusanRepository.FindById(ctx, tx, int(prosesBisnis.BidangUrusanId.Int32))
			helper.PanicIfError(err)
			bidangUrusan = &web.ProsBisBidangUrusanRespons{
				Id:           bidangurusanData.ID,
				BidangUrusan: bidangurusanData.BidangUrusan,
			}
		}

		var rabLevel1, rabLevel2, rabLevel3 *web.ProsBisReferensiArsitekturRespons

		if prosesBisnis.RabLevel1ID.Valid {
			rabLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel1ID.Int32))
			helper.PanicIfError(err)
			rabLevel1 = &web.ProsBisReferensiArsitekturRespons{
				Id:              rabLevel1Data.IdReferensi,
				Kode_referensi:  rabLevel1Data.Kode_referensi,
				Nama_referensi:  rabLevel1Data.Nama_referensi,
				Level_referensi: rabLevel1Data.Level_referensi,
			}
		}

		if prosesBisnis.RabLevel2ID.Valid {
			rabLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel2ID.Int32))
			helper.PanicIfError(err)
			rabLevel2 = &web.ProsBisReferensiArsitekturRespons{
				Id:              rabLevel2Data.IdReferensi,
				Kode_referensi:  rabLevel2Data.Kode_referensi,
				Nama_referensi:  rabLevel2Data.Nama_referensi,
				Level_referensi: rabLevel2Data.Level_referensi,
			}
		}

		if prosesBisnis.RabLevel3ID.Valid {
			rabLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel3ID.Int32))
			helper.PanicIfError(err)
			rabLevel3 = &web.ProsBisReferensiArsitekturRespons{
				Id:              rabLevel3Data.IdReferensi,
				Kode_referensi:  rabLevel3Data.Kode_referensi,
				Nama_referensi:  rabLevel3Data.Nama_referensi,
				Level_referensi: rabLevel3Data.Level_referensi,
			}
		}

		var rabLevel4, rabLevel5, rabLevel6 *web.ProsBisPohonKinerjaRespons

		if prosesBisnis.RabLevel4ID.Valid {
			rabLevel4Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel4ID.Int32))
			helper.PanicIfError(err)
			rabLevel4 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel4Data.ID,
				NamaPohon:  rabLevel4Data.NamaPohon,
				LevelPohon: rabLevel4Data.LevelPohon,
			}
		}

		if prosesBisnis.RabLevel5ID.Valid {
			rabLevel5Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel5ID.Int32))
			helper.PanicIfError(err)
			rabLevel5 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel5Data.ID,
				NamaPohon:  rabLevel5Data.NamaPohon,
				LevelPohon: rabLevel5Data.LevelPohon,
			}
		}

		if prosesBisnis.RabLevel6ID.Valid {
			rabLevel6Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel6ID.Int32))
			helper.PanicIfError(err)
			rabLevel6 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel6Data.ID,
				NamaPohon:  rabLevel6Data.NamaPohon,
				LevelPohon: rabLevel6Data.LevelPohon,
			}
		}

		response := web.ProsesBisnisRespons{
			ID:               prosesBisnis.ID,
			NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
			SasaranKota:      sasaranKota,
			KodeProsesBisnis: prosesBisnis.KodeProsesBisnis,
			KodeOPD:          prosesBisnis.KodeOPD,
			BidangUrusan:     bidangUrusan,
			RabLevel1:        rabLevel1,
			RabLevel2:        rabLevel2,
			RabLevel3:        rabLevel3,
			RabLevel4:        rabLevel4,
			RabLevel5:        rabLevel5,
			RabLevel6:        rabLevel6,
			Tahun:            prosesBisnis.Tahun,
			CreatedAt:        prosesBisnis.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        prosesBisnis.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *ProsesBisnisServiceImpl) FindById(ctx context.Context, prosesbisnisId int) (web.ProsesBisnisRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.ProsesBisnisRespons{}, err
	}
	defer helper.CommitOrRollback(tx)

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, prosesbisnisId)
	if err != nil {
		return web.ProsesBisnisRespons{}, err
	}

	var sasaranKota *web.ProsbisSasaranKotaRespons
	if prosesBisnis.SasaranKotaId.Valid {
		sasarankotaData, err := service.SasaranKotaRepository.FindById(ctx, tx, int(prosesBisnis.SasaranKotaId.Int32))
		helper.PanicIfError(err)
		sasaranKota = &web.ProsbisSasaranKotaRespons{
			ID:      sasarankotaData.ID,
			Sasaran: sasarankotaData.Sasaran,
		}
	}

	var bidangUrusan *web.ProsBisBidangUrusanRespons

	if prosesBisnis.BidangUrusanId.Valid {
		bidangurusanData, err := service.BidangUrusanRepository.FindById(ctx, tx, int(prosesBisnis.BidangUrusanId.Int32))
		helper.PanicIfError(err)
		bidangUrusan = &web.ProsBisBidangUrusanRespons{
			Id:           bidangurusanData.ID,
			BidangUrusan: bidangurusanData.BidangUrusan,
		}
	}

	var rabLevel1, rabLevel2, rabLevel3 *web.ProsBisReferensiArsitekturRespons

	if prosesBisnis.RabLevel1ID.Valid {
		rabLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel1ID.Int32))
		helper.PanicIfError(err)
		rabLevel1 = &web.ProsBisReferensiArsitekturRespons{
			Id:              rabLevel1Data.IdReferensi,
			Kode_referensi:  rabLevel1Data.Kode_referensi,
			Nama_referensi:  rabLevel1Data.Nama_referensi,
			Level_referensi: rabLevel1Data.Level_referensi,
		}
	}

	if prosesBisnis.RabLevel2ID.Valid {
		rabLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel2ID.Int32))
		helper.PanicIfError(err)
		rabLevel2 = &web.ProsBisReferensiArsitekturRespons{
			Id:              rabLevel2Data.IdReferensi,
			Kode_referensi:  rabLevel2Data.Kode_referensi,
			Nama_referensi:  rabLevel2Data.Nama_referensi,
			Level_referensi: rabLevel2Data.Level_referensi,
		}
	}

	if prosesBisnis.RabLevel3ID.Valid {
		rabLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel3ID.Int32))
		helper.PanicIfError(err)
		rabLevel3 = &web.ProsBisReferensiArsitekturRespons{
			Id:              rabLevel3Data.IdReferensi,
			Kode_referensi:  rabLevel3Data.Kode_referensi,
			Nama_referensi:  rabLevel3Data.Nama_referensi,
			Level_referensi: rabLevel3Data.Level_referensi,
		}
	}

	var rabLevel4, rabLevel5, rabLevel6 *web.ProsBisPohonKinerjaRespons

	if prosesBisnis.RabLevel4ID.Valid {
		rabLevel4Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel4ID.Int32))
		helper.PanicIfError(err)
		rabLevel4 = &web.ProsBisPohonKinerjaRespons{
			ID:         rabLevel4Data.ID,
			NamaPohon:  rabLevel4Data.NamaPohon,
			LevelPohon: rabLevel4Data.LevelPohon,
		}
	}

	if prosesBisnis.RabLevel5ID.Valid {
		rabLevel5Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel5ID.Int32))
		helper.PanicIfError(err)
		rabLevel5 = &web.ProsBisPohonKinerjaRespons{
			ID:         rabLevel5Data.ID,
			NamaPohon:  rabLevel5Data.NamaPohon,
			LevelPohon: rabLevel5Data.LevelPohon,
		}
	}

	if prosesBisnis.RabLevel6ID.Valid {
		rabLevel6Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel6ID.Int32))
		helper.PanicIfError(err)
		rabLevel6 = &web.ProsBisPohonKinerjaRespons{
			ID:         rabLevel6Data.ID,
			NamaPohon:  rabLevel6Data.NamaPohon,
			LevelPohon: rabLevel6Data.LevelPohon,
		}
	}

	response := web.ProsesBisnisRespons{
		ID:               prosesBisnis.ID,
		NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
		SasaranKota:      sasaranKota,
		KodeProsesBisnis: prosesBisnis.KodeProsesBisnis,
		KodeOPD:          prosesBisnis.KodeOPD,
		BidangUrusan:     bidangUrusan,
		RabLevel1:        rabLevel1,
		RabLevel2:        rabLevel2,
		RabLevel3:        rabLevel3,
		RabLevel4:        rabLevel4,
		RabLevel5:        rabLevel5,
		RabLevel6:        rabLevel6,
		Tahun:            prosesBisnis.Tahun,
		CreatedAt:        prosesBisnis.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        prosesBisnis.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *ProsesBisnisServiceImpl) Insert(ctx context.Context, request web.ProsesBisnisCreateRequest) web.ProsesBisnisRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	currentTime := time.Now()
	kodeprosbis := helper.GenerateRandomKodeProsesBisnis()

	prosesBisnis := domain.ProsesBisnis{
		NamaProsesBisnis: request.NamaProsesBisnis,

		SasaranKotaId: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		KodeProsesBisnis: kodeprosbis,
		KodeOPD:          request.KodeOPD,
		BidangUrusanId: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RabLevel1ID: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RabLevel2ID: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RabLevel3ID: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RabLevel4ID: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RabLevel5ID: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		RabLevel6ID: sql.NullInt32{
			Int32: int32(0),
			Valid: false,
		},
		Tahun:     request.Tahun,
		CreatedAt: currentTime,
	}

	// Mengisi nilai jika tidak null
	if request.SasaranKotaId != nil {
		prosesBisnis.SasaranKotaId = sql.NullInt32{
			Int32: int32(*request.SasaranKotaId),
			Valid: true,
		}
	}

	if request.BidangUrusanId != nil {
		prosesBisnis.BidangUrusanId = sql.NullInt32{
			Int32: int32(*request.BidangUrusanId),
			Valid: true,
		}
	}

	if request.RabLevel1ID != nil {
		prosesBisnis.RabLevel1ID = sql.NullInt32{
			Int32: int32(*request.RabLevel1ID),
			Valid: true,
		}
	}

	if request.RabLevel2ID != nil {
		prosesBisnis.RabLevel2ID = sql.NullInt32{
			Int32: int32(*request.RabLevel2ID),
			Valid: true,
		}
	}

	if request.RabLevel3ID != nil {
		prosesBisnis.RabLevel3ID = sql.NullInt32{
			Int32: int32(*request.RabLevel3ID),
			Valid: true,
		}
	}

	if request.RabLevel4ID != nil {
		prosesBisnis.RabLevel4ID = sql.NullInt32{
			Int32: int32(*request.RabLevel4ID),
			Valid: true,
		}
	}

	if request.RabLevel5ID != nil {
		prosesBisnis.RabLevel5ID = sql.NullInt32{
			Int32: int32(*request.RabLevel5ID),
			Valid: true,
		}
	}

	if request.RabLevel6ID != nil {
		prosesBisnis.RabLevel6ID = sql.NullInt32{
			Int32: int32(*request.RabLevel6ID),
			Valid: true,
		}
	}

	prosesBisnis = service.ProsesBisnisRepository.Insert(ctx, tx, prosesBisnis)
	return helper.ToProsesBisnisResponse(prosesBisnis)
}

func (service *ProsesBisnisServiceImpl) Update(ctx context.Context, request web.ProsesBisnisUpdateRequest) web.ProsesBisnisRespons {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	prosesBisnis.NamaProsesBisnis = request.NamaProsesBisnis
	prosesBisnis.SasaranKotaId = sql.NullInt32{Int32: int32(request.SasaranKotaId), Valid: request.SasaranKotaId != 0}
	if prosesBisnis.KodeProsesBisnis == "" {
		prosesBisnis.KodeProsesBisnis = helper.GenerateRandomKodeProsesBisnis()
	}
	prosesBisnis.KodeOPD = request.KodeOPD
	prosesBisnis.BidangUrusanId = sql.NullInt32{Int32: int32(request.BidangUrusanId), Valid: request.BidangUrusanId != 0}
	prosesBisnis.RabLevel1ID = sql.NullInt32{Int32: int32(request.RabLevel1ID), Valid: request.RabLevel1ID != 0}
	prosesBisnis.RabLevel2ID = sql.NullInt32{Int32: int32(request.RabLevel2ID), Valid: request.RabLevel2ID != 0}
	prosesBisnis.RabLevel3ID = sql.NullInt32{Int32: int32(request.RabLevel3ID), Valid: request.RabLevel3ID != 0}
	prosesBisnis.RabLevel4ID = sql.NullInt32{Int32: int32(request.RabLevel4ID), Valid: request.RabLevel4ID != 0}
	prosesBisnis.RabLevel5ID = sql.NullInt32{Int32: int32(request.RabLevel5ID), Valid: request.RabLevel5ID != 0}
	prosesBisnis.RabLevel6ID = sql.NullInt32{Int32: int32(request.RabLevel6ID), Valid: request.RabLevel6ID != 0}
	prosesBisnis.Tahun = request.Tahun
	prosesBisnis.UpdatedAt = request.UpdatedAt

	prosesBisnis = service.ProsesBisnisRepository.Update(ctx, tx, prosesBisnis)
	return helper.ToProsesBisnisResponse(prosesBisnis)
}

func (service *ProsesBisnisServiceImpl) Delete(ctx context.Context, prosesbisnisId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, prosesbisnisId)
	helper.PanicIfError(err)

	service.ProsesBisnisRepository.Delete(ctx, tx, prosesBisnis)
}

func (service *ProsesBisnisServiceImpl) FindByNull(ctx context.Context) ([]web.ProsesBisnisRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	prosesBisnisList, err := service.ProsesBisnisRepository.FindByNull(ctx, tx)
	if err != nil {
		return nil, err
	}

	var responseList []web.ProsesBisnisRespons
	for _, prosesBisnis := range prosesBisnisList {
		response := web.ProsesBisnisRespons{
			ID:               prosesBisnis.ID,
			NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
			KodeOPD:          prosesBisnis.KodeOPD,
			KodeProsesBisnis: prosesBisnis.KodeProsesBisnis,
			Tahun:            prosesBisnis.Tahun,
			CreatedAt:        prosesBisnis.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        prosesBisnis.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		if prosesBisnis.SasaranKotaId.Valid {
			sasarankotaData, err := service.SasaranKotaRepository.FindById(ctx, tx, int(prosesBisnis.SasaranKotaId.Int32))
			if err != nil {
				return nil, err
			}
			response.SasaranKota = &web.ProsbisSasaranKotaRespons{
				ID:      sasarankotaData.ID,
				Sasaran: sasarankotaData.Sasaran,
			}
		}

		if prosesBisnis.BidangUrusanId.Valid {
			bidangurusanData, err := service.BidangUrusanRepository.FindById(ctx, tx, int(prosesBisnis.BidangUrusanId.Int32))
			if err != nil {
				return nil, err
			}
			response.BidangUrusan = &web.ProsBisBidangUrusanRespons{
				Id:           bidangurusanData.ID,
				BidangUrusan: bidangurusanData.BidangUrusan,
			}
		}

		if prosesBisnis.RabLevel1ID.Valid {
			rabLevel1Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel1ID.Int32))
			if err != nil {
				return nil, err
			}
			response.RabLevel1 = &web.ProsBisReferensiArsitekturRespons{
				Id:              rabLevel1Data.IdReferensi,
				Kode_referensi:  rabLevel1Data.Kode_referensi,
				Nama_referensi:  rabLevel1Data.Nama_referensi,
				Level_referensi: rabLevel1Data.Level_referensi,
			}
		}

		if prosesBisnis.RabLevel2ID.Valid {
			rabLevel2Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel2ID.Int32))
			if err != nil {
				return nil, err
			}
			response.RabLevel2 = &web.ProsBisReferensiArsitekturRespons{
				Id:              rabLevel2Data.IdReferensi,
				Kode_referensi:  rabLevel2Data.Kode_referensi,
				Nama_referensi:  rabLevel2Data.Nama_referensi,
				Level_referensi: rabLevel2Data.Level_referensi,
			}
		}

		if prosesBisnis.RabLevel3ID.Valid {
			rabLevel3Data, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel3ID.Int32))
			if err != nil {
				return nil, err
			}
			response.RabLevel3 = &web.ProsBisReferensiArsitekturRespons{
				Id:              rabLevel3Data.IdReferensi,
				Kode_referensi:  rabLevel3Data.Kode_referensi,
				Nama_referensi:  rabLevel3Data.Nama_referensi,
				Level_referensi: rabLevel3Data.Level_referensi,
			}
		}

		if prosesBisnis.RabLevel4ID.Valid {
			rabLevel4Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel4ID.Int32))
			if err != nil {
				return nil, err
			}
			response.RabLevel4 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel4Data.ID,
				NamaPohon:  rabLevel4Data.NamaPohon,
				LevelPohon: rabLevel4Data.LevelPohon,
			}
		}

		if prosesBisnis.RabLevel5ID.Valid {
			rabLevel5Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel5ID.Int32))
			if err != nil {
				return nil, err
			}
			response.RabLevel5 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel5Data.ID,
				NamaPohon:  rabLevel5Data.NamaPohon,
				LevelPohon: rabLevel5Data.LevelPohon,
			}
		}

		if prosesBisnis.RabLevel6ID.Valid {
			rabLevel6Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel6ID.Int32))
			if err != nil {
				return nil, err
			}
			response.RabLevel6 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel6Data.ID,
				NamaPohon:  rabLevel6Data.NamaPohon,
				LevelPohon: rabLevel6Data.LevelPohon,
			}
		}

		responseList = append(responseList, response)
	}

	return responseList, nil
}
