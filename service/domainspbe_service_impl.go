package service

import (
	"api_spbe_kota_madiun/exception"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type DomainSPBEServiceImpl struct {
	DomainSPBERepository repository.DomainSPBERepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewDomainSPBEService(domainspbeRepository repository.DomainSPBERepository, DB *sql.DB, validate *validator.Validate) DomainSPBEService {
	return &DomainSPBEServiceImpl{
		DomainSPBERepository: domainspbeRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *DomainSPBEServiceImpl) Insert(ctx context.Context, request web.DomainSPBECreateRequest) web.DomainSPBEResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Membuat kode domain acak menggunakan fungsi helper
	randomKodeDomain := helper.GenerateRandomKodeDomain()

	domainspbe := domain.DomainSPBE{
		NamaDomain: request.NamaDomain,
		KodeDomain: randomKodeDomain,
		Tahun:      request.Tahun,
	}

	domainspbe = service.DomainSPBERepository.Insert(ctx, tx, domainspbe)
	return helper.ToDomainSPBEResponse(domainspbe)
}

func (service *DomainSPBEServiceImpl) Update(ctx context.Context, request web.DomainSPBEUpdateRequest) web.DomainSPBEResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	domainspbe, err := service.DomainSPBERepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	domainspbe.NamaDomain = request.NamaDomain
	if request.KodeDomain == "" {
		domainspbe.KodeDomain = helper.GenerateRandomKodeDomain()
	} else {
		domainspbe.KodeDomain = request.KodeDomain
	}
	domainspbe.Tahun = request.Tahun

	domainspbe = service.DomainSPBERepository.Update(ctx, tx, domainspbe)
	return helper.ToDomainSPBEResponse(domainspbe)
}

func (service *DomainSPBEServiceImpl) Delete(ctx context.Context, domainspbeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	domainspbe, err := service.DomainSPBERepository.FindById(ctx, tx, domainspbeId)
	helper.PanicIfError(err)

	service.DomainSPBERepository.Delete(ctx, tx, domainspbe)
}

func (service *DomainSPBEServiceImpl) FindAll(ctx context.Context) []web.DomainSPBEResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	domains := service.DomainSPBERepository.FindAll(ctx, tx)
	return helper.ToDomainSPBEResponses(domains)
}

func (service *DomainSPBEServiceImpl) FindById(ctx context.Context, domainspbeId int) web.DomainSPBEResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	domainspbe, err := service.DomainSPBERepository.FindById(ctx, tx, domainspbeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToDomainSPBEResponse(domainspbe)
}
