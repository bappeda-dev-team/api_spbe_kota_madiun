package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type ProsesBisnisServiceImpl struct{
	ProsesBisnisRepository repository.ProsesBisnisRepository
	DB							  *sql.DB
}

func NewProsesBisnisService(prosesbisnisRepository repository.ProsesBisnisRepository, DB *sql.DB) ProsesBisnisService{
	return &ProsesBisnisServiceImpl{
		ProsesBisnisRepository : prosesbisnisRepository,
		DB: DB,
	}
}

func(service *ProsesBisnisServiceImpl)FindAll(ctx context.Context)[]web.ProsesBinsisRespons{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	prosbis:= service.ProsesBisnisRepository.FindAll(ctx,tx)
	return helper.ToProsesBisnisResponses(prosbis)
}

