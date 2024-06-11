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
	SasaranKotaRepository repository.SasaranKotaRepository
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	PohonKinerjaRepository        repository.PohonKinerjaRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewProsesBisnisService(referensiarsitekturRepository repository.ReferensiArsitekturRepository, sasarankotaRepository repository.SasaranKotaRepository, prosesbisnisRepository repository.ProsesBisnisRepository, pohonkinerjaRepository repository.PohonKinerjaRepository, DB *sql.DB, validate *validator.Validate) ProsesBisnisService {
	return &ProsesBisnisServiceImpl{
		ProsesBisnisRepository:        prosesbisnisRepository,
		SasaranKotaRepository:		 sasarankotaRepository,
		ReferensiArsitekturRepository: referensiarsitekturRepository,
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
		sasaranKota ,err := service.SasaranKotaRepository.FindById(ctx, tx, prosesBisnis.SasaranKotaId)
		helper.PanicIfError(err)

		rabLevel1, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel1ID)
		helper.PanicIfError(err)

		rabLevel2, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel2ID)
		helper.PanicIfError(err)

		rabLevel3, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel3ID)
		helper.PanicIfError(err)

		var rabLevel4, rabLevel5, rabLevel6 *web.PohonKinerjaRespons

		if prosesBisnis.RabLevel4ID.Valid {
			rabLevel4Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel4ID.Int64))
			helper.PanicIfError(err)
			rabLevel4 = &web.PohonKinerjaRespons{
				ID:         rabLevel4Data.ID,
				NamaPohon:  rabLevel4Data.NamaPohon,
				JenisPohon: rabLevel4Data.JenisPohon,
				LevelPohon: rabLevel4Data.LevelPohon,
				KodeOpd:    rabLevel4Data.KodeOpd,
				Tahun:      rabLevel4Data.Tahun,
				CreatedAt:  rabLevel4Data.CreatedAt,
				UpdatedAt:  rabLevel4Data.UpdatedAt,
			}
		}

		if prosesBisnis.RabLevel5ID.Valid {
			rabLevel5Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel5ID.Int64))
			helper.PanicIfError(err)
			rabLevel5 = &web.PohonKinerjaRespons{
				ID:         rabLevel5Data.ID,
				NamaPohon:  rabLevel5Data.NamaPohon,
				JenisPohon: rabLevel5Data.JenisPohon,
				LevelPohon: rabLevel5Data.LevelPohon,
				KodeOpd:    rabLevel5Data.KodeOpd,
				Tahun:      rabLevel5Data.Tahun,
				CreatedAt:  rabLevel5Data.CreatedAt,
				UpdatedAt:  rabLevel5Data.UpdatedAt,
			}
		}

		if prosesBisnis.RabLevel6ID.Valid {
			rabLevel6Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.RabLevel6ID.Int64))
			helper.PanicIfError(err)
			rabLevel6 = &web.PohonKinerjaRespons{
				ID:         rabLevel6Data.ID,
				NamaPohon:  rabLevel6Data.NamaPohon,
				JenisPohon: rabLevel6Data.JenisPohon,
				LevelPohon: rabLevel6Data.LevelPohon,
				KodeOpd:    rabLevel6Data.KodeOpd,
				Tahun:      rabLevel6Data.Tahun,
				CreatedAt:  rabLevel6Data.CreatedAt,
				UpdatedAt:  rabLevel6Data.UpdatedAt,
			}
		}

		response := web.ProsesBisnisRespons{
			ID:               prosesBisnis.ID,
			NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
			SasaranKota:      web.SasaranKotaRespons{
				ID: sasaranKota.ID,
				Sasaran: sasaranKota.Sasaran,
				TujuanKota: sasaranKota.TujuanKota,
				StrategiKota: sasaranKota.StrategiKota,
				CreatedAt: sasaranKota.CreatedAt,
				UpdatedAt: sasaranKota.UpdatedAt,
			},
			KodeProsesBisnis: prosesBisnis.KodeProsesBisnis,
			KodeOPD:          prosesBisnis.KodeOPD,
			BidangUrusan:     prosesBisnis.BidangUrusan,
			RabLevel1: web.ReferensiArsitekturResponse{
				Id:              rabLevel1.IdReferensi,
				Kode_referensi:  rabLevel1.Kode_referensi,
				Nama_referensi:  rabLevel1.Nama_referensi,
				Level_referensi: rabLevel1.Level_referensi,
				Jenis_referensi: rabLevel1.Jenis_referensi,
				Tahun:           rabLevel1.Tahun,
				Created_at:      rabLevel1.Created_at,
				Updated_at:      rabLevel1.Updated_at,
			},
			RabLevel2: web.ReferensiArsitekturResponse{
				Id:              rabLevel2.IdReferensi,
				Kode_referensi:  rabLevel2.Kode_referensi,
				Nama_referensi:  rabLevel2.Nama_referensi,
				Level_referensi: rabLevel2.Level_referensi,
				Jenis_referensi: rabLevel2.Jenis_referensi,
				Tahun:           rabLevel2.Tahun,
				Created_at:      rabLevel2.Created_at,
				Updated_at:      rabLevel2.Updated_at,
			},
			RabLevel3: web.ReferensiArsitekturResponse{
				Id:              rabLevel3.IdReferensi,
				Kode_referensi:  rabLevel3.Kode_referensi,
				Nama_referensi:  rabLevel3.Nama_referensi,
				Level_referensi: rabLevel3.Level_referensi,
				Jenis_referensi: rabLevel3.Jenis_referensi,
				Tahun:           rabLevel3.Tahun,
				Created_at:      rabLevel3.Created_at,
				Updated_at:      rabLevel3.Updated_at,
			},
			RabLevel4: rabLevel4,
			RabLevel5: rabLevel5,
			RabLevel6: rabLevel6,
			Tahun:     prosesBisnis.Tahun,
			CreatedAt: prosesBisnis.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: prosesBisnis.UpdatedAt.Format("2006-01-02 15:04:05"),
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

	sasaranKota ,err := service.SasaranKotaRepository.FindById(ctx, tx, prosesBisnis.SasaranKotaId)
	helper.PanicIfError(err)

	rabLevel1, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel1ID)
	helper.PanicIfError(err)

	rabLevel2, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel2ID)
	helper.PanicIfError(err)

	rabLevel3, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel3ID)
	helper.PanicIfError(err)


	response := web.ProsesBisnisRespons{
		ID:               prosesBisnis.ID,
		NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
		SasaranKota:      web.SasaranKotaRespons{
			ID: sasaranKota.ID,
			Sasaran: sasaranKota.Sasaran,
			TujuanKota: sasaranKota.TujuanKota,
			StrategiKota: sasaranKota.StrategiKota,
			CreatedAt: sasaranKota.CreatedAt,
			UpdatedAt: sasaranKota.UpdatedAt,
		},
		KodeProsesBisnis: prosesBisnis.KodeProsesBisnis,
		KodeOPD:          prosesBisnis.KodeOPD,
		BidangUrusan:     prosesBisnis.BidangUrusan,
		RabLevel1: web.ReferensiArsitekturResponse{
			Id:              rabLevel1.IdReferensi,
			Kode_referensi:  rabLevel1.Kode_referensi,
			Nama_referensi:  rabLevel1.Nama_referensi,
			Level_referensi: rabLevel1.Level_referensi,
			Jenis_referensi: rabLevel1.Jenis_referensi,
			Tahun:           rabLevel1.Tahun,
			Created_at:      rabLevel1.Created_at,
			Updated_at:      rabLevel1.Updated_at,
		},
		RabLevel2: web.ReferensiArsitekturResponse{
			Id:              rabLevel2.IdReferensi,
			Kode_referensi:  rabLevel2.Kode_referensi,
			Nama_referensi:  rabLevel2.Nama_referensi,
			Level_referensi: rabLevel2.Level_referensi,
			Jenis_referensi: rabLevel2.Jenis_referensi,
			Tahun:           rabLevel2.Tahun,
			Created_at:      rabLevel2.Created_at,
			Updated_at:      rabLevel2.Updated_at,
		},
		RabLevel3: web.ReferensiArsitekturResponse{
			Id:              rabLevel3.IdReferensi,
			Kode_referensi:  rabLevel3.Kode_referensi,
			Nama_referensi:  rabLevel3.Nama_referensi,
			Level_referensi: rabLevel3.Level_referensi,
			Jenis_referensi: rabLevel3.Jenis_referensi,
			Tahun:           rabLevel3.Tahun,
			Created_at:      rabLevel3.Created_at,
			Updated_at:      rabLevel3.Updated_at,
		},
		Tahun:     prosesBisnis.Tahun,
		CreatedAt: prosesBisnis.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: prosesBisnis.UpdatedAt.Format("2006-01-02 15:04:05"),
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
        SasaranKotaId: 	request.SasaranKotaId,
        KodeProsesBisnis: kodeprosbis,
        KodeOPD:          request.KodeOPD,
        BidangUrusan:     request.BidangUrusan,
        RabLevel1ID:      request.RabLevel1ID,
        RabLevel2ID:      request.RabLevel2ID,
        RabLevel3ID:      request.RabLevel3ID,
		RabLevel4ID: sql.NullInt64{
            Int64: int64(0),
            Valid: false,
        },
        RabLevel5ID: sql.NullInt64{
            Int64: int64(0),
            Valid: false,
        },
        RabLevel6ID: sql.NullInt64{
            Int64: int64(0),
            Valid: false,
        },
        Tahun:     request.Tahun,
        CreatedAt: currentTime,
    }
	    // Mengisi nilai jika tidak null
		if request.RabLevel4ID != nil {
			prosesBisnis.RabLevel4ID = sql.NullInt64{
				Int64: int64(*request.RabLevel4ID),
				Valid: true,
			}
		}

		if request.RabLevel5ID != nil {
			prosesBisnis.RabLevel5ID = sql.NullInt64{
				Int64: int64(*request.RabLevel5ID),
				Valid: true,
			}
		}
		
		if request.RabLevel6ID != nil {
			prosesBisnis.RabLevel6ID = sql.NullInt64{
				Int64: int64(*request.RabLevel6ID),
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
	prosesBisnis.SasaranKotaId = request.SasaranKotaId
	if prosesBisnis.KodeProsesBisnis == "" {
		prosesBisnis.KodeProsesBisnis = helper.GenerateRandomKodeProsesBisnis()
	}
	prosesBisnis.KodeOPD = request.KodeOPD
	prosesBisnis.BidangUrusan = request.BidangUrusan
	prosesBisnis.RabLevel1ID = request.RabLevel1ID
	prosesBisnis.RabLevel2ID = request.RabLevel2ID
	prosesBisnis.RabLevel3ID = request.RabLevel3ID
	prosesBisnis.RabLevel4ID = sql.NullInt64{Int64: int64(request.RabLevel4ID), Valid: request.RabLevel4ID != 0}
	prosesBisnis.RabLevel5ID = sql.NullInt64{Int64: int64(request.RabLevel5ID), Valid: request.RabLevel5ID != 0}
	prosesBisnis.RabLevel6ID = sql.NullInt64{Int64: int64(request.RabLevel6ID), Valid: request.RabLevel6ID != 0}
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
