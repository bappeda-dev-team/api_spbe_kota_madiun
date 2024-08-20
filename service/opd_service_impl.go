package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
)

type OpdServiceImpl struct {
	OpdRepository repository.OpdRepository
	DB            *sql.DB
}

func NewOpdServiceImpl(opdRepository repository.OpdRepository, DB *sql.DB) *OpdServiceImpl {
	return &OpdServiceImpl{
		OpdRepository: opdRepository,
		DB:            DB,
	}
}

func (service *OpdServiceImpl) FetchKodeOpd(ctx context.Context) (web.Opd, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.Opd{}, err
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
	result, err := service.OpdRepository.FetchKodeOpd(ctx, tx)
	if err != nil {
		log.Println("Error fetching and inserting API data:", err)
		return web.Opd{}, err
	}

	return result, nil
}

func (service *OpdServiceImpl) FindAll(ctx context.Context, kodeOPD string) []web.Opd {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	opd := service.OpdRepository.FindAll(ctx, tx, kodeOPD)
	return helper.ToOpdResponses(opd)
}
