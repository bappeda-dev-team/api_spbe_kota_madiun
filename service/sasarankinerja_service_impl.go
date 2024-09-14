package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type SasaranKinerjaPegawaiServiceImpl struct {
	SasaranKinerjaRepository repository.SasaranKinerjaPegawaiRepository
	DB                       *sql.DB
}

func NewSasaranKinerjaPegawaiServiceImpl(sasaranKinerjaRepository repository.SasaranKinerjaPegawaiRepository, db *sql.DB) *SasaranKinerjaPegawaiServiceImpl {
	return &SasaranKinerjaPegawaiServiceImpl{
		SasaranKinerjaRepository: sasaranKinerjaRepository,
		DB:                       db,
	}
}

func (service *SasaranKinerjaPegawaiServiceImpl) FetchsasaranKinerja(ctx context.Context, kodeOPD string, tahun string) (web.ApiResponsSasaranKinerjaPegawai, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}
	defer tx.Rollback()

	result, err := service.SasaranKinerjaRepository.FetchsasaranKinerja(ctx, tx, kodeOPD, tahun)
	if err != nil {
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}

	err = tx.Commit()
	if err != nil {
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}

	return result, nil
}

func (service *SasaranKinerjaPegawaiServiceImpl) FindByKodeOpdAndTahun(ctx context.Context, kodeOPD string, tahun string) []web.SasaranKinerjaPegawaiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sasaran := service.SasaranKinerjaRepository.FindByKodeOpdAndTahun(ctx, tx, kodeOPD, tahun)
	return helper.ToSasaranKinerjaResponses(sasaran)
}

func (service *SasaranKinerjaPegawaiServiceImpl) FindById(ctx context.Context, sasaranKinerjaId int) (web.SasaranKinerjaPegawaiResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.SasaranKinerjaPegawaiResponse{}, err
	}
	defer tx.Rollback()

	sasaranKinerja, err := service.SasaranKinerjaRepository.FindById(ctx, tx, sasaranKinerjaId)
	if err != nil {
		return web.SasaranKinerjaPegawaiResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		return web.SasaranKinerjaPegawaiResponse{}, err
	}

	return helper.ToSasaranKinerjaPegawaiResponse(sasaranKinerja), nil
}
