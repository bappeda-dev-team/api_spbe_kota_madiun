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

type PohonKinerjaServiceImpl struct {
	PohonKinerjaRepository repository.PohonKinerjaRepository
	DB                     *sql.DB
}

func NewPohonKinerjaService(pohonkinerjaRepository repository.PohonKinerjaRepository, DB *sql.DB) PohonKinerjaService {
	return &PohonKinerjaServiceImpl{
		PohonKinerjaRepository: pohonkinerjaRepository,
		DB:                     DB,
	}
}

func (service *PohonKinerjaServiceImpl) FindById(ctx context.Context, pohonId int) web.PohonKinerjaRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pohon, err := service.PohonKinerjaRepository.FindById(ctx, tx, pohonId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPohonKinerjaResponse(pohon)
}

func (service *PohonKinerjaServiceImpl) FindAll(ctx context.Context) []web.PohonKinerjaRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pohon := service.PohonKinerjaRepository.FindAll(ctx, tx)
	return helper.ToPohonResponses(pohon)
}

func (service *PohonKinerjaServiceImpl)InsertApi(ctx context.Context) (web.PohonKinerjaApi, error){
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.PohonKinerjaApi{}, err
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
	result, err := service.PohonKinerjaRepository.InsertApi(ctx, tx)
	if err != nil {
		log.Println("Error fetching and inserting API data:", err)
		return web.SasaranKotaApi{}, err
	}

	return result, nil
}