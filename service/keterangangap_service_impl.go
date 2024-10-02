package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type KeteranganGapServiceImpl struct {
	KeteranganGapRepository repository.KeteranganGapRepository
	DB                      *sql.DB
}

func NewKeteranganGapServiceImpl(keteranganGapRepository repository.KeteranganGapRepository, DB *sql.DB) *KeteranganGapServiceImpl {
	return &KeteranganGapServiceImpl{
		KeteranganGapRepository: keteranganGapRepository,
		DB:                      DB,
	}
}

func (service *KeteranganGapServiceImpl) Insert(ctx context.Context, request web.KeteranganGapCreateRequest) (web.KeteranganGapRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keteranganGap := domain.KeteranganGap{
		KodeOpd:        request.KodeOpd,
		IdProsesBisnis: request.IdProsesBisnis,
		KeteranganGap:  request.KeteranganGap,
	}

	result, err := service.KeteranganGapRepository.Insert(ctx, tx, keteranganGap)
	if err != nil {
		return web.KeteranganGapRespons{}, err
	}

	return helper.ToKeteranganGapResponse(result), nil
}

func (service *KeteranganGapServiceImpl) FindById(ctx context.Context, id int) (web.KeteranganGapRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.KeteranganGapRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.KeteranganGapRespons{}, err
	}

	return helper.ToKeteranganGapResponse(result), nil
}

func (service *KeteranganGapServiceImpl) Update(ctx context.Context, request web.KeteranganGapUpdateRequest) (web.KeteranganGapRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keteranganGap := domain.KeteranganGap{
		Id:             request.Id,
		KodeOpd:        request.KodeOpd,
		IdProsesBisnis: request.IdProsesBisnis,
		KeteranganGap:  request.KeteranganGap,
	}

	result, err := service.KeteranganGapRepository.Update(ctx, tx, keteranganGap)
	if err != nil {
		return web.KeteranganGapRespons{}, err
	}

	return helper.ToKeteranganGapResponse(result), nil
}

func (service *KeteranganGapServiceImpl) Delete(ctx context.Context, id int) (web.KeteranganGapRespons, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keteranganGap, err := service.KeteranganGapRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.KeteranganGapRespons{}, err
	}

	err = service.KeteranganGapRepository.Delete(ctx, tx, id)
	if err != nil {
		return web.KeteranganGapRespons{}, err
	}

	return helper.ToKeteranganGapResponse(keteranganGap), nil
}
