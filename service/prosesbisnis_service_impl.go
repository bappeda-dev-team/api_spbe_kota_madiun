package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"errors"
	"log"
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

func NewProsesBisnisServiceImpl(referensiarsitekturRepository repository.ReferensiArsitekturRepository, sasarankotaRepository repository.SasaranKotaRepository, bidangUrusanRepository repository.BidangUrusanRepository, prosesbisnisRepository repository.ProsesBisnisRepository, pohonkinerjaRepository repository.PohonKinerjaRepository, DB *sql.DB, validate *validator.Validate) *ProsesBisnisServiceImpl {
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
	helper.PanicIfError(err)
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

		if prosesBisnis.StrategicId.Valid {
			rabLevel4Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.StrategicId.Int32))
			helper.PanicIfError(err)
			rabLevel4 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel4Data.ID,
				NamaPohon:  rabLevel4Data.NamaPohon,
				LevelPohon: rabLevel4Data.LevelPohon,
			}
		}

		if prosesBisnis.TacticalId.Valid {
			rabLevel5Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.TacticalId.Int32))
			helper.PanicIfError(err)
			rabLevel5 = &web.ProsBisPohonKinerjaRespons{
				ID:         rabLevel5Data.ID,
				NamaPohon:  rabLevel5Data.NamaPohon,
				LevelPohon: rabLevel5Data.LevelPohon,
			}
		}

		if prosesBisnis.OperationalId.Valid {
			rabLevel6Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.OperationalId.Int32))
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

func (service *ProsesBisnisServiceImpl) FindById(ctx context.Context, prosesbisnisId int, kodeOPD string) (web.ProsesBisnisRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.ProsesBisnisRespons{}, err
	}
	defer helper.CommitOrRollback(tx)

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, prosesbisnisId)
	if err != nil {
		return web.ProsesBisnisRespons{}, err
	}

	if kodeOPD != "" && prosesBisnis.KodeOPD != kodeOPD {
		return web.ProsesBisnisRespons{}, errors.New("proses bisnis tidak ditemukan untuk OPD ini")
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

	if prosesBisnis.StrategicId.Valid {
		rabLevel4Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.StrategicId.Int32))
		helper.PanicIfError(err)
		rabLevel4 = &web.ProsBisPohonKinerjaRespons{
			ID:         rabLevel4Data.ID,
			NamaPohon:  rabLevel4Data.NamaPohon,
			LevelPohon: rabLevel4Data.LevelPohon,
		}
	}

	if prosesBisnis.TacticalId.Valid {
		rabLevel5Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.TacticalId.Int32))
		helper.PanicIfError(err)
		rabLevel5 = &web.ProsBisPohonKinerjaRespons{
			ID:         rabLevel5Data.ID,
			NamaPohon:  rabLevel5Data.NamaPohon,
			LevelPohon: rabLevel5Data.LevelPohon,
		}
	}

	if prosesBisnis.OperationalId.Valid {
		rabLevel6Data, err := service.PohonKinerjaRepository.FindById(ctx, tx, int(prosesBisnis.OperationalId.Int32))
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
	kodeprosbis := helper.GenerateRandomKode()

	var namaProsesBisnis string
	if request.RabLevel6ID != nil {
		pohonKinerja, err := service.PohonKinerjaRepository.FindById(ctx, tx, *request.RabLevel6ID)
		if err != nil {
			helper.PanicIfError(err)
		}
		namaProsesBisnis = pohonKinerja.NamaPohon
	} else {
		namaProsesBisnis = request.NamaProsesBisnis
	}

	prosesBisnis := domain.ProsesBisnis{
		NamaProsesBisnis: namaProsesBisnis,

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
		prosesBisnis.StrategicId = sql.NullInt32{
			Int32: int32(*request.RabLevel4ID),
			Valid: true,
		}
	}

	if request.RabLevel5ID != nil {
		prosesBisnis.TacticalId = sql.NullInt32{
			Int32: int32(*request.RabLevel5ID),
			Valid: true,
		}
	}

	if request.RabLevel6ID != nil {
		prosesBisnis.OperationalId = sql.NullInt32{
			Int32: int32(*request.RabLevel6ID),
			Valid: true,
		}
	}

	prosesBisnis = service.ProsesBisnisRepository.Insert(ctx, tx, prosesBisnis)
	return helper.ToProsesBisnisResponse(prosesBisnis)
}

func (service *ProsesBisnisServiceImpl) Update(ctx context.Context, request web.ProsesBisnisUpdateRequest) web.ProsesBisnisRespons {
	err := service.Validate.Struct(request)
	if err != nil {
		log.Printf("Validation error: %v", err)
		helper.PanicIfError(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		helper.PanicIfError(err)
	}
	defer helper.CommitOrRollback(tx)

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		log.Printf("Error finding Proses Bisnis: %v", err)
		helper.PanicIfError(err)
	}

	if request.RabLevel6ID != 0 {
		pohonKinerja, err := service.PohonKinerjaRepository.FindById(ctx, tx, request.RabLevel6ID)
		if err != nil {
			if err.Error() == "pohon Kinerja is not found" {
				log.Printf("Pohon Kinerja dengan ID %d tidak ditemukan", request.RabLevel6ID)
			} else {
				log.Printf("Error fetching Pohon Kinerja: %v", err)
			}
		} else {
			prosesBisnis.NamaProsesBisnis = pohonKinerja.NamaPohon
		}
	}

	prosesBisnis.SasaranKotaId = sql.NullInt32{Int32: int32(request.SasaranKotaId), Valid: request.SasaranKotaId != 0}
	if prosesBisnis.KodeProsesBisnis == "" {
		prosesBisnis.KodeProsesBisnis = helper.GenerateRandomKode()
	}
	prosesBisnis.KodeOPD = request.KodeOPD
	prosesBisnis.BidangUrusanId = sql.NullInt32{Int32: int32(request.BidangUrusanId), Valid: request.BidangUrusanId != 0}
	prosesBisnis.RabLevel1ID = sql.NullInt32{Int32: int32(request.RabLevel1ID), Valid: request.RabLevel1ID != 0}
	prosesBisnis.RabLevel2ID = sql.NullInt32{Int32: int32(request.RabLevel2ID), Valid: request.RabLevel2ID != 0}
	prosesBisnis.RabLevel3ID = sql.NullInt32{Int32: int32(request.RabLevel3ID), Valid: request.RabLevel3ID != 0}
	prosesBisnis.StrategicId = sql.NullInt32{Int32: int32(request.RabLevel4ID), Valid: request.RabLevel4ID != 0}
	prosesBisnis.TacticalId = sql.NullInt32{Int32: int32(request.RabLevel5ID), Valid: request.RabLevel5ID != 0}
	prosesBisnis.OperationalId = sql.NullInt32{Int32: int32(request.RabLevel6ID), Valid: request.RabLevel6ID != 0}
	prosesBisnis.Tahun = request.Tahun
	prosesBisnis.UpdatedAt = request.UpdatedAt

	prosesBisnis = service.ProsesBisnisRepository.Update(ctx, tx, prosesBisnis)
	return helper.ToProsesBisnisResponse(prosesBisnis)
}

func (service *ProsesBisnisServiceImpl) Delete(ctx context.Context, prosesbisnisId int, kodeOPD string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, prosesbisnisId)
	helper.PanicIfError(err)

	if prosesBisnis.KodeOPD != kodeOPD {
		panic(errors.New("proses bisnis tidak ditemukan untuk OPD ini"))
	}

	service.ProsesBisnisRepository.Delete(ctx, tx, prosesBisnis)

	return nil
}

func (service *ProsesBisnisServiceImpl) GetProsesBisnisGrouped(ctx context.Context, kodeOpd string, tahun int) ([]web.GapProsesBisnis, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prosesBisnisList, err := service.ProsesBisnisRepository.GapProsesBisnis(ctx, tx, kodeOpd, tahun)
	helper.PanicIfError(err)

	var webProsesBisnisList []web.GapProsesBisnis
	idMap := make(map[int]*web.GapProsesBisnis)

	for _, pb := range prosesBisnisList {
		if _, exists := idMap[pb.ID]; !exists {
			idMap[pb.ID] = &web.GapProsesBisnis{
				ID:               pb.ID,
				KodeOpd:          pb.KodeOpd,
				Tahun:            pb.Tahun,
				NamaProsesBisnis: pb.NamaProsesBisnis,
				KodeProsesBisnis: pb.KodeProsesBisnis,
				Layanans:         []web.GapLayanan{},
				DataDanInformasi: []web.GapDataDanInformasi{},
				Aplikasi:         []web.GapAplikasi{},
			}
		}

		webPb := idMap[pb.ID]

		if pb.Layanan != nil {
			for _, l := range pb.Layanan {
				var webNullString web.NullString

				if l.NamaLayanan.Valid {
					webNullString.String = l.NamaLayanan.String
					webNullString.Valid = true
				}

				tempLayanan := web.GapLayanan{NamaLayanan: webNullString}
				if !helper.ContainsLayanan(webPb.Layanans, tempLayanan) {
					webPb.Layanans = append(webPb.Layanans, tempLayanan)
				}
			}
		}

		if pb.DataDanInformasi != nil {
			for _, d := range pb.DataDanInformasi {
				var webNullString web.NullString

				if d.NamaData.Valid {
					webNullString.String = d.NamaData.String
					webNullString.Valid = true
				}

				tempData := web.GapDataDanInformasi{NamaData: webNullString}
				if !helper.ContainData(webPb.DataDanInformasi, tempData) {
					webPb.DataDanInformasi = append(webPb.DataDanInformasi, tempData)
				}
			}
		}

		if pb.Aplikasi != nil {
			for _, a := range pb.Aplikasi {
				var webNullString web.NullString

				if a.NamaAplikasi.Valid {
					webNullString.String = a.NamaAplikasi.String
					webNullString.Valid = true
				}

				temAplikasi := web.GapAplikasi{NamaAplikasi: webNullString}
				if !helper.ContainAplikasi(webPb.Aplikasi, temAplikasi) {
					webPb.Aplikasi = append(webPb.Aplikasi, temAplikasi)
				}
			}
		}
	}

	for _, webPb := range idMap {
		webProsesBisnisList = append(webProsesBisnisList, *webPb)
	}

	return webProsesBisnisList, nil
}

func (service *ProsesBisnisServiceImpl) GetProsesBisnisNoGap(ctx context.Context, kodeOPD string, tahun int) ([]web.GapProsesBisnis, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	prosesBisnisList, err := service.ProsesBisnisRepository.NoGapProsesBisnis(ctx, tx, kodeOPD, tahun)
	if err != nil {
		return nil, err
	}

	var webProsesBisnisList []web.GapProsesBisnis
	idMap := make(map[int]*web.GapProsesBisnis)

	for _, pb := range prosesBisnisList {
		webPb, exists := idMap[pb.ID]
		if !exists {
			webPb = &web.GapProsesBisnis{
				ID:               pb.ID,
				KodeOpd:          pb.KodeOpd,
				Tahun:            pb.Tahun,
				NamaProsesBisnis: pb.NamaProsesBisnis,
				KodeProsesBisnis: pb.KodeProsesBisnis,
			}
			idMap[pb.ID] = webPb
		}

		if pb.Layanan != nil {
			for _, l := range pb.Layanan {
				var webNullString web.NullString
				if l.NamaLayanan.Valid {
					webNullString.String = l.NamaLayanan.String
					webNullString.Valid = true
				}
				tempLayanan := web.GapLayanan{NamaLayanan: webNullString}
				if !helper.ContainsLayanan(webPb.Layanans, tempLayanan) {
					webPb.Layanans = append(webPb.Layanans, tempLayanan)
				}
			}
		}

		if pb.DataDanInformasi != nil {
			for _, d := range pb.DataDanInformasi {
				var webNullString web.NullString
				if d.NamaData.Valid {
					webNullString.String = d.NamaData.String
					webNullString.Valid = true
				}
				tempData := web.GapDataDanInformasi{NamaData: webNullString}
				if !helper.ContainData(webPb.DataDanInformasi, tempData) {
					webPb.DataDanInformasi = append(webPb.DataDanInformasi, tempData)
				}
			}
		}

		if pb.Aplikasi != nil {
			for _, a := range pb.Aplikasi {
				var webNullString web.NullString
				if a.NamaAplikasi.Valid {
					webNullString.String = a.NamaAplikasi.String
					webNullString.Valid = true
				}
				tempAplikasi := web.GapAplikasi{NamaAplikasi: webNullString}
				if !helper.ContainAplikasi(webPb.Aplikasi, tempAplikasi) {
					webPb.Aplikasi = append(webPb.Aplikasi, tempAplikasi)
				}
			}
		}
	}

	for _, webPb := range idMap {
		webProsesBisnisList = append(webProsesBisnisList, *webPb)
	}

	return webProsesBisnisList, nil
}
