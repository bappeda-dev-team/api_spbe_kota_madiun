package service

import (
	"api_spbe_kota_madiun/exception"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"errors"
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

func (service *PohonKinerjaServiceImpl) InsertApi(ctx context.Context, kodeOPD string, tahun string) (web.PohonKinerjaApi, error) {

	tx, err := service.DB.Begin()
	if err != nil {
		return web.PohonKinerjaApi{}, err
	}
	defer tx.Rollback()

	result, err := service.PohonKinerjaRepository.InsertApi(ctx, tx, kodeOPD, tahun)
	if err != nil {
		return web.PohonKinerjaApi{}, err
	}

	err = tx.Commit()
	if err != nil {
		return web.PohonKinerjaApi{}, err
	}

	return result, nil

	// tx, err := service.DB.BeginTx(ctx, nil)
	// if err != nil {
	// 	log.Println("Error starting transaction:", err)
	// 	return web.PohonKinerjaApi{}, err
	// }
	// defer func() {
	// 	if p := recover(); p != nil {
	// 		tx.Rollback()
	// 		panic(p)
	// 	} else if err != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		err = tx.Commit()
	// 	}
	// }()
	// result, err := service.PohonKinerjaRepository.InsertApi(ctx, tx, kodeOPD, tahun)
	// if err != nil {
	// 	log.Println("Error fetching and inserting API data:", err)
	// 	return web.PohonKinerjaApi{}, err
	// }

	// return result, nil
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
			ID:        strategic.ID,
			Level:     strategic.LevelPohon,
			NamaPohon: strategic.NamaPohon,
		})
	}

	if len(tactical) > 0 {
		parentID, _ := strconv.Atoi(tactical[0].Parent)
		response.Tactical = append(response.Tactical, web.TacticalResponse{
			ID:        tactical[0].ID,
			ParentID:  parentID,
			Level:     tactical[0].LevelPohon,
			NamaPohon: tactical[0].NamaPohon,
		})
	}

	if len(operational) > 0 {
		parentID, _ := strconv.Atoi(operational[0].Parent)
		response.Operational = append(response.Operational, web.OperationalResponse{
			ID:        operational[0].ID,
			ParentID:  parentID,
			Level:     operational[0].LevelPohon,
			NamaPohon: operational[0].NamaPohon,
		})
	}

	return response
}

func (service *PohonKinerjaServiceImpl) GetHierarchy(ctx context.Context, id int) (web.PohonKinerjaHierarchyResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.PohonKinerjaHierarchyResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	item, err := service.PohonKinerjaRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PohonKinerjaHierarchyResponse{}, err
	}

	response := web.PohonKinerjaHierarchyResponse{
		Strategic:   []web.StrategicResponse{},
		Tactical:    []web.TacticalResponse{},
		Operational: []web.OperationalResponse{},
	}

	switch item.LevelPohon {
	case 4:
		response.Strategic = append(response.Strategic, web.StrategicResponse{
			ID:        item.ID,
			Level:     item.LevelPohon,
			NamaPohon: item.NamaPohon,
		})

		// Cari children (level 5)
		tacticalChildren, err := service.PohonKinerjaRepository.FindChildren(ctx, tx, id)
		if err != nil {
			return web.PohonKinerjaHierarchyResponse{}, err
		}
		for _, tacticalChild := range tacticalChildren {
			tacticalResponse := web.TacticalResponse{
				ID:        tacticalChild.ID,
				ParentID:  id,
				Level:     tacticalChild.LevelPohon,
				NamaPohon: tacticalChild.NamaPohon,
			}
			response.Tactical = append(response.Tactical, tacticalResponse)

			// Cari children dari level 5 (level 6)
			operationalChildren, err := service.PohonKinerjaRepository.FindChildren(ctx, tx, tacticalChild.ID)
			if err != nil {
				return web.PohonKinerjaHierarchyResponse{}, err
			}
			for _, operationalChild := range operationalChildren {
				response.Operational = append(response.Operational, web.OperationalResponse{
					ID:        operationalChild.ID,
					ParentID:  tacticalChild.ID,
					Level:     operationalChild.LevelPohon,
					NamaPohon: operationalChild.NamaPohon,
				})
			}
		}

	case 5:
		parent, err := service.PohonKinerjaRepository.FindById(ctx, tx, helper.ConvertStringToInt(item.Parent))
		if err == nil {
			response.Strategic = append(response.Strategic, web.StrategicResponse{
				ID:        parent.ID,
				Level:     parent.LevelPohon,
				NamaPohon: parent.NamaPohon,
			})
		}

		// Tambahkan Tactical
		response.Tactical = append(response.Tactical, web.TacticalResponse{
			ID:        item.ID,
			ParentID:  helper.ConvertStringToInt(item.Parent),
			Level:     item.LevelPohon,
			NamaPohon: item.NamaPohon,
		})

		// Cari children di level 6
		children, err := service.PohonKinerjaRepository.FindChildren(ctx, tx, item.ID)
		if err != nil {
			return web.PohonKinerjaHierarchyResponse{}, err
		}

		for _, child := range children {
			response.Operational = append(response.Operational, web.OperationalResponse{
				ID:        child.ID,
				ParentID:  item.ID,
				Level:     child.LevelPohon,
				NamaPohon: child.NamaPohon,
			})
		}

	case 6:
		// Ambil parent dari level 5 (tactical)
		parentID := helper.ConvertStringToInt(item.Parent)
		if parentID == 0 {
			// Jika parent kosong, return error atau handle sebagai tidak ditemukan
			return web.PohonKinerjaHierarchyResponse{}, errors.New("parent ID tidak ditemukan")
		}

		// Ambil Tactical Parent
		tacticalParent, err := service.PohonKinerjaRepository.FindById(ctx, tx, parentID)
		if err == nil {
			response.Tactical = append(response.Tactical, web.TacticalResponse{
				ID:        tacticalParent.ID,
				ParentID:  helper.ConvertStringToInt(tacticalParent.Parent),
				Level:     tacticalParent.LevelPohon,
				NamaPohon: tacticalParent.NamaPohon,
			})
		}

		// Ambil Strategic Parent jika parent dari tactical ada
		if tacticalParent.Parent != "" {
			strategicParentID := helper.ConvertStringToInt(tacticalParent.Parent)
			strategicParent, err := service.PohonKinerjaRepository.FindById(ctx, tx, strategicParentID)
			if err == nil {
				response.Strategic = append(response.Strategic, web.StrategicResponse{
					ID:        strategicParent.ID,
					Level:     strategicParent.LevelPohon,
					NamaPohon: strategicParent.NamaPohon,
				})
			}
		}

		// Tambahkan item level 6 ke dalam Operational
		response.Operational = append(response.Operational, web.OperationalResponse{
			ID:        item.ID,
			ParentID:  parentID, // Pastikan parent ID benar
			Level:     item.LevelPohon,
			NamaPohon: item.NamaPohon,
		})

		// 	response.Operational = append(response.Operational, web.OperationalResponse{
		// 		ID:        item.ID,
		// 		ParentID:  helper.ConvertStringToInt(item.Parent),
		// 		Level:     item.LevelPohon,
		// 		NamaPohon: item.NamaPohon,
		// 	})

		// 	tacticalParents, err := service.PohonKinerjaRepository.FindChildren(ctx, tx, helper.ConvertStringToInt(item.Parent))
		// 	if err == nil && len(tacticalParents) > 0 {
		// 		tacticalParent := tacticalParents[0]
		// 		response.Tactical = append(response.Tactical, web.TacticalResponse{
		// 			ID:        tacticalParent.ID,
		// 			ParentID:  helper.ConvertStringToInt(tacticalParent.Parent),
		// 			Level:     tacticalParent.LevelPohon,
		// 			NamaPohon: tacticalParent.NamaPohon,
		// 		})

		// 		strategicParents, err := service.PohonKinerjaRepository.FindChildren(ctx, tx, helper.ConvertStringToInt(tacticalParent.Parent))
		// 		if err == nil && len(strategicParents) > 0 {
		// 			strategicParent := strategicParents[0]
		// 			response.Strategic = append(response.Strategic, web.StrategicResponse{
		// 				ID:        strategicParent.ID,
		// 				Level:     strategicParent.LevelPohon,
		// 				NamaPohon: strategicParent.NamaPohon,
		// 			})
		// 		}
		// 	}
		// }
	}

	return response, nil
}
