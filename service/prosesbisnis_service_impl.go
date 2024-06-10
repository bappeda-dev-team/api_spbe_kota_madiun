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
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	PohonKinerjaRepository        repository.PohonKinerjaRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewProsesBisnisService(referensiarsitekturRepository repository.ReferensiArsitekturRepository, prosesbisnisRepository repository.ProsesBisnisRepository, pohonkinerjaRepository repository.PohonKinerjaRepository, DB *sql.DB, validate *validator.Validate) ProsesBisnisService {
	return &ProsesBisnisServiceImpl{
		ProsesBisnisRepository:        prosesbisnisRepository,
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
		rabLevel1, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel1ID)
		helper.PanicIfError(err)

		rabLevel2, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel2ID)
		helper.PanicIfError(err)

		rabLevel3, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel3ID)
		helper.PanicIfError(err)

		rabLevel4, err := service.PohonKinerjaRepository.FindById(ctx, tx, prosesBisnis.RabLevel4ID)
		helper.PanicIfError(err)

		rabLevel5, err := service.PohonKinerjaRepository.FindById(ctx, tx, prosesBisnis.RabLevel5ID)
		helper.PanicIfError(err)

		rabLevel6, err := service.PohonKinerjaRepository.FindById(ctx, tx, prosesBisnis.RabLevel6ID)
		helper.PanicIfError(err)

		response := web.ProsesBisnisRespons{
			ID:               prosesBisnis.ID,
			NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
			SasaranKota:      prosesBisnis.SasaranKota,
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
			RabLevel4: web.PohonKinerjaRespons{
				ID:         rabLevel4.ID,
				NamaPohon:  rabLevel4.NamaPohon,
				JenisPohon: rabLevel4.JenisPohon,
				LevelPohon: rabLevel4.LevelPohon,
				KodeOpd:    rabLevel4.KodeOpd,
				Tahun:      rabLevel4.Tahun,
				CreatedAt:  rabLevel4.CreatedAt,
				UpdatedAt:  rabLevel4.UpdatedAt,
			},
			RabLevel5: web.PohonKinerjaRespons{
				ID:         rabLevel5.ID,
				NamaPohon:  rabLevel5.NamaPohon,
				JenisPohon: rabLevel5.JenisPohon,
				LevelPohon: rabLevel5.LevelPohon,
				KodeOpd:    rabLevel5.KodeOpd,
				Tahun:      rabLevel5.Tahun,
				CreatedAt:  rabLevel5.CreatedAt,
				UpdatedAt:  rabLevel5.UpdatedAt,
			},
			RabLevel6: web.PohonKinerjaRespons{
				ID:         rabLevel6.ID,
				NamaPohon:  rabLevel6.NamaPohon,
				JenisPohon: rabLevel6.JenisPohon,
				LevelPohon: rabLevel6.LevelPohon,
				KodeOpd:    rabLevel6.KodeOpd,
				Tahun:      rabLevel6.Tahun,
				CreatedAt:  rabLevel6.CreatedAt,
				UpdatedAt:  rabLevel6.UpdatedAt,
			},
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

	rabLevel1, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel1ID)
	helper.PanicIfError(err)

	rabLevel2, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel2ID)
	helper.PanicIfError(err)

	rabLevel3, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel3ID)
	helper.PanicIfError(err)

	response := web.ProsesBisnisRespons{
		ID:               prosesBisnis.ID,
		NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
		SasaranKota:      prosesBisnis.SasaranKota,
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

	prosesBisnis := domain.ProsesBisnis{
		NamaProsesBisnis: request.NamaProsesBisnis,
		SasaranKota:      request.SasaranKota,
		KodeProsesBisnis: request.KodeProsesBisnis,
		KodeOPD:          request.KodeOPD,
		BidangUrusan:     request.BidangUrusan,
		RabLevel1ID:      request.RabLevel1ID,
		RabLevel2ID:      request.RabLevel2ID,
		RabLevel3ID:      request.RabLevel3ID,
		Tahun:            request.Tahun,
		CreatedAt:        currentTime,
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

	prosesBisnis.NamaProsesBisnis = request.KodeProsesBisnis
	prosesBisnis.SasaranKota = request.SasaranKota
	prosesBisnis.KodeProsesBisnis = request.KodeProsesBisnis
	prosesBisnis.KodeOPD = request.KodeOPD
	prosesBisnis.BidangUrusan = request.BidangUrusan
	prosesBisnis.RabLevel1ID = request.RabLevel1ID
	prosesBisnis.RabLevel2ID = request.RabLevel2ID
	prosesBisnis.RabLevel3ID = request.RabLevel3ID
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
