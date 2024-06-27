package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type BidangUrusanServiceImpl struct {
	BidangUrusanRepository repository.BidangUrusanRepository
	DB                     *sql.DB
}

func NewBidangUrusanService(bidangurusanRepository repository.BidangUrusanRepository, DB *sql.DB) BidangUrusanService {
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
