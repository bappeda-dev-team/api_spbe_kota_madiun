package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
)

type BidangUrusanServiceImpl struct {
	BidangUrusanRepository repository.BidangUrusanRepository
	DB                     *sql.DB
}

func NewBidangUrusanServiceImpl(bidangurusanRepository repository.BidangUrusanRepository, DB *sql.DB) *BidangUrusanServiceImpl {
	return &BidangUrusanServiceImpl{
		BidangUrusanRepository: bidangurusanRepository,
		DB:                     DB,
	}
}

func (service *BidangUrusanServiceImpl) FindAll(ctx context.Context) []web.BidangUrusanRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bidang := service.BidangUrusanRepository.FindAll(ctx, tx)
	return helper.ToBidangResponses(bidang)
}

func (service *BidangUrusanServiceImpl) FetchUrusan(ctx context.Context) (web.BidangUrusanOPD, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.BidangUrusanOPD{}, err
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
	result, err := service.BidangUrusanRepository.FetchBidangUrusan(ctx, tx)
	if err != nil {
		log.Println("Error fetching and inserting API data:", err)
		return web.BidangUrusanOPD{}, err
	}

	return result, nil
}
