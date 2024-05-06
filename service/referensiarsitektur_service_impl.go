package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ReferensiArsitekturServiceImpl struct {
	ReferensiArsitekturRepository repository.ReferensiArsitekturRepository
	DB							  *sql.DB
	Validate 					  *validator.Validate
}

func NewReferensiArsitekturService(referensiarsitekturRepository repository.ReferensiArsitekturRepository,  DB *sql.DB, validate *validator.Validate) ReferensiArsitekturService{
	return &ReferensiArsitekturServiceImpl{
		ReferensiArsitekturRepository: referensiarsitekturRepository,
		DB: DB,
		Validate: validate,
	}
}

func(service *ReferensiArsitekturServiceImpl) Insert(ctx context.Context, request web.ReferensiArsitekturCreateRequest) web.ReferensiArsitekturResponse{
	err :=service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiarsitektur := domain.ReferensiArsitektur{
		Kode_referensi: request.Kode_referensi,
		Nama_referensi: request.Nama_referensi,
		Level_referensi: request.Level_referensi,
		Jenis_referensi: request.Jenis_referensi,
		Created_at: request.Created_at,
	}
	
	referensiarsitektur = service.ReferensiArsitekturRepository.Insert(ctx, tx, referensiarsitektur)
	return helper.ToReferensiArsitekturResponse(referensiarsitektur)
}

func(service *ReferensiArsitekturServiceImpl)Update(ctx context.Context, request web.ReferensiArsitekturUpdateRequest) web.ReferensiArsitekturResponse{
	err :=service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiarsitektur, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, request.Id)	
	helper.PanicIfError(err)
	
	referensiarsitektur.Kode_referensi= request.Kode_referensi
	referensiarsitektur.Nama_referensi= request.Nama_referensi
	referensiarsitektur.Level_referensi= request.Level_referensi
	referensiarsitektur.Jenis_referensi= request.Jenis_referensi
	referensiarsitektur.Updated_at= request.Updated_at
	
	referensiarsitektur = service.ReferensiArsitekturRepository.Update(ctx, tx, referensiarsitektur)
	return helper.ToReferensiArsitekturResponse(referensiarsitektur)
}

func(service *ReferensiArsitekturServiceImpl)Delete(ctx context.Context, referensiarsitekturId int){
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	referensiarsitektur,err := service.ReferensiArsitekturRepository.FindById(ctx, tx, referensiarsitekturId)
	helper.PanicIfError(err)

	service.ReferensiArsitekturRepository.Delete(ctx, tx, referensiarsitektur)

}

func(service *ReferensiArsitekturServiceImpl)FindAll(ctx context.Context)[]web.ReferensiArsitekturResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	reference := service.ReferensiArsitekturRepository.FindAll(ctx,tx)
	return helper.ToReferenceResponses(reference)
}