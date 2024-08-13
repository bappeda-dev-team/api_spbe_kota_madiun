package service

import (
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
)

type UrusanServiceImpl struct {
	UrusanRepository repository.UrusanRepository
	DB               *sql.DB
}

func NewUrusanServiceImpl(urusanRepository repository.UrusanRepository, DB *sql.DB) *UrusanServiceImpl {
	return &UrusanServiceImpl{
		UrusanRepository: urusanRepository,
		DB:               DB,
	}
}

func (service *UrusanServiceImpl) FetchUrusan(ctx context.Context) (web.UrusanOPD, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.UrusanOPD{}, err
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
	result, err := service.UrusanRepository.FetchUrusan(ctx, tx)
	if err != nil {
		log.Println("Error fetching and inserting API data:", err)
		return web.UrusanOPD{}, err
	}

	return result, nil
}
