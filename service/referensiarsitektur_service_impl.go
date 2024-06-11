package service

import (
	"api_spbe_kota_madiun/exception"
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

type ReferensiArsitekturServiceImpl struct {
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewReferensiArsitekturService(referensiarsitekturRepository repository.ReferensiArsitekturRepository, DB *sql.DB, validate *validator.Validate) ReferensiArsitekturService {
	return &ReferensiArsitekturServiceImpl{
		ReferensiArsitekturRepository: referensiarsitekturRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func (service *ReferensiArsitekturServiceImpl) Insert(ctx context.Context, request web.ReferensiArsitekturCreateRequest) web.ReferensiArsitekturResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	currentTime := time.Now()

	referensiarsitektur := domain.ReferensiArsitektur{
		Kode_referensi:  request.Kode_referensi,
		Nama_referensi:  request.Nama_referensi,
		Level_referensi: request.Level_referensi,
		Jenis_referensi: request.Jenis_referensi,
		Created_at:      currentTime,
		Tahun:           request.Tahun,
	}

	referensiarsitektur = service.ReferensiArsitekturRepository.Insert(ctx, tx, referensiarsitektur)
	return helper.ToReferensiArsitekturResponse(referensiarsitektur)
}

func (service *ReferensiArsitekturServiceImpl) Update(ctx context.Context, request web.ReferensiArsitekturUpdateRequest) web.ReferensiArsitekturResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiarsitektur, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	referensiarsitektur.Kode_referensi = request.Kode_referensi
	referensiarsitektur.Nama_referensi = request.Nama_referensi
	referensiarsitektur.Level_referensi = request.Level_referensi
	referensiarsitektur.Jenis_referensi = request.Jenis_referensi
	referensiarsitektur.Updated_at = request.Updated_at
	referensiarsitektur.Tahun = request.Tahun

	referensiarsitektur = service.ReferensiArsitekturRepository.Update(ctx, tx, referensiarsitektur)
	return helper.ToReferensiArsitekturResponse(referensiarsitektur)
}

func (service *ReferensiArsitekturServiceImpl) Delete(ctx context.Context, referensiarsitekturId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiarsitektur, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, referensiarsitekturId)
	helper.PanicIfError(err)

	service.ReferensiArsitekturRepository.Delete(ctx, tx, referensiarsitektur)

}

func (service *ReferensiArsitekturServiceImpl) FindAll(ctx context.Context) []web.ReferensiArsitekturResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	reference := service.ReferensiArsitekturRepository.FindAll(ctx, tx)
	return helper.ToReferenceResponses(reference)
}

func (service *ReferensiArsitekturServiceImpl) GetDataHierarchy(ctx context.Context, kodeReferensi string) ([]web.ReferensiArsitekturResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiList, err := service.ReferensiArsitekturRepository.FindByKodeRef(ctx, tx, kodeReferensi)
	if err != nil {
		if err.Error() == "data not found" {
			log.Println("Service: Data not found for kodeReferensi:", kodeReferensi)
			return nil, err
		}
		return nil, err
	}

	if len(referensiList) == 0 {
		log.Println("Service: Data not found for kodeReferensi:", kodeReferensi)
		return nil, errors.New("data not found")
	}

	var responseList []web.ReferensiArsitekturResponse
	for _, referensi := range referensiList {
		response := web.ReferensiArsitekturResponse{
			Id:              referensi.IdReferensi,
			Kode_referensi:  referensi.Kode_referensi,
			Nama_referensi:  referensi.Nama_referensi,
			Level_referensi: referensi.Level_referensi,
			Jenis_referensi: referensi.Jenis_referensi,
			Created_at:      referensi.Created_at,
			Updated_at:      referensi.Updated_at,
			Tahun:           referensi.Tahun,
		}
		responseList = append(responseList, response)
	}

	return responseList, nil

}

func (service *ReferensiArsitekturServiceImpl) FindById(ctx context.Context, referensiarsitekturId int) web.ReferensiArsitekturResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiarsitektur, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, referensiarsitekturId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToReferensiArsitekturResponse(referensiarsitektur)
}
