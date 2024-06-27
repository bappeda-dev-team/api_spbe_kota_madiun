package service

import (
	"api_spbe_kota_madiun/exception"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type SasaranKotaServiceImpl struct {
	SasaranKotaRepository repository.SasaranKotaRepository
	DB                    *sql.DB
}

func NewSasaranKotaService(sasarankotaRepository repository.SasaranKotaRepository, DB *sql.DB) SasaranKotaService {
	return &SasaranKotaServiceImpl{
		SasaranKotaRepository: sasarankotaRepository,
		DB:                    DB,
	}
}

func (service *SasaranKotaServiceImpl) FindById(ctx context.Context, sasarankotaId int) web.SasaranKotaRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sasaran, err := service.SasaranKotaRepository.FindById(ctx, tx, sasarankotaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSasaranKotaResponse(sasaran)
}

func (service *SasaranKotaServiceImpl) FindAll(ctx context.Context) []web.SasaranKotaRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sasaran := service.SasaranKotaRepository.FindAll(ctx, tx)
	return helper.ToSasaranResponses(sasaran)
}
