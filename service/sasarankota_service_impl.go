package service

import (
	"api_spbe_kota_madiun/exception"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
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

func (service *SasaranKotaServiceImpl) InsertApi(ctx context.Context) (web.SasaranKotaApi, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.SasaranKotaApi{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	result, err := service.SasaranKotaRepository.InsertApi(ctx, tx)
	if err != nil {
		log.Println("Error fetching and inserting API data:", err)
		return web.SasaranKotaApi{}, err
	}

	return result, nil
}
