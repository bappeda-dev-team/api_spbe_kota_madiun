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

func (service *BidangUrusanServiceImpl) FindBidangUrusanOPD(ctx context.Context, kodeOPD string) ([]web.BidangUrusanRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return nil, err
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

	opdUrusanBidang, err := service.BidangUrusanRepository.FindBidangUrusanOPD(ctx, tx, kodeOPD)
	if err != nil {
		log.Println("Error finding Bidang Urusan OPD:", err)
		return nil, err
	}

	var result []web.BidangUrusanRespons
	for _, item := range opdUrusanBidang {
		bidangUrusan, err := service.BidangUrusanRepository.FindByBidangUrusan(ctx, tx, item.BidangUrusan.String)
		helper.PanicIfError(err)
		result = append(result, web.BidangUrusanRespons{
			Id:               bidangUrusan.ID, // Menggunakan ID dari FindByBidangUrusan
			KodeOpd:          item.KodeOpd.String,
			KodeBidangUrusan: bidangUrusan.KodeBidangUrusan,
			BidangUrusan:     bidangUrusan.BidangUrusan,
		})
	}

	return result, nil
}
