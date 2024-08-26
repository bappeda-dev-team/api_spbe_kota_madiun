package service

import (
	"api_spbe_kota_madiun/exception"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
	"strconv"
)

type PohonKinerjaServiceImpl struct {
	PohonKinerjaRepository repository.PohonKinerjaRepository
	DB                     *sql.DB
}

func NewPohonKinerjaServiceImpl(pohonkinerjaRepository repository.PohonKinerjaRepository, DB *sql.DB) *PohonKinerjaServiceImpl {
	return &PohonKinerjaServiceImpl{
		PohonKinerjaRepository: pohonkinerjaRepository,
		DB:                     DB,
	}
}

func (service *PohonKinerjaServiceImpl) FindById(ctx context.Context, pohonId int) web.PohonKinerjaRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pohon, err := service.PohonKinerjaRepository.FindById(ctx, tx, pohonId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPohonKinerjaResponse(pohon)
}

func (service *PohonKinerjaServiceImpl) FindAll(ctx context.Context, kodeOpd string, tahun int) []web.PohonKinerjaRespons {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pohon := service.PohonKinerjaRepository.FindAll(ctx, tx, kodeOpd, tahun)
	return helper.ToPohonResponses(pohon)
}

func (service *PohonKinerjaServiceImpl) InsertApi(ctx context.Context) (web.PohonKinerjaApi, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.PohonKinerjaApi{}, err
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
	result, err := service.PohonKinerjaRepository.InsertApi(ctx, tx)
	if err != nil {
		log.Println("Error fetching and inserting API data:", err)
		return web.PohonKinerjaApi{}, err
	}

	return result, nil
}

func (service *PohonKinerjaServiceImpl) FindByOperational(ctx context.Context, pohonId int) web.PohonKinerjaHierarchyResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	strategic, tactical, operational, err := service.PohonKinerjaRepository.FindByOperational(ctx, tx, pohonId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	response := web.PohonKinerjaHierarchyResponse{
		Strategic:   []web.StrategicResponse{},
		Tactical:    []web.TacticalResponse{},
		Operational: []web.OperationalResponse{},
	}

	if strategic.ID != 0 {
		response.Strategic = append(response.Strategic, web.StrategicResponse{
			IDStrategic: strategic.ID,
			Level:       strategic.LevelPohon,
			NamaPohon:   strategic.NamaPohon,
		})
	}

	if len(tactical) > 0 {
		parentID, _ := strconv.Atoi(tactical[0].Parent)
		response.Tactical = append(response.Tactical, web.TacticalResponse{
			IDTactical: tactical[0].ID,
			Parent:     parentID,
			Level:      tactical[0].LevelPohon,
			NamaPohon:  tactical[0].NamaPohon,
		})
	}

	if len(operational) > 0 {
		parentID, _ := strconv.Atoi(operational[0].Parent)
		response.Operational = append(response.Operational, web.OperationalResponse{
			IDOperational: operational[0].ID,
			Parent:        parentID,
			Level:         operational[0].LevelPohon,
			NamaPohon:     operational[0].NamaPohon,
		})
	}

	return response
}
