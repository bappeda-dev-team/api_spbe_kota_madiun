package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
)

type OpdServiceImpl struct {
	OpdRepository          repository.OpdRepository
	UrusanRepository       repository.UrusanRepository
	BidangUrusanRepository repository.BidangUrusanRepository
	DB                     *sql.DB
}

func NewOpdServiceImpl(opdRepository repository.OpdRepository, urusanRepository repository.UrusanRepository, bidangUrusanRepository repository.BidangUrusanRepository, DB *sql.DB) *OpdServiceImpl {
	return &OpdServiceImpl{
		OpdRepository:          opdRepository,
		UrusanRepository:       urusanRepository,
		BidangUrusanRepository: bidangUrusanRepository,
		DB:                     DB,
	}
}

func (service *OpdServiceImpl) FetchAllData(ctx context.Context) ([]web.Opd, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error memulai transaksi:", err)
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

	opds, err := service.OpdRepository.FetchAllData(ctx, tx)
	if err != nil {
		log.Println("Error mengambil dan menyimpan data API:", err)
		return nil, err
	}

	return opds, nil
}

func (service *OpdServiceImpl) FindAll(ctx context.Context, kodeOPD string) []web.Opd {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	opd := service.OpdRepository.FindAll(ctx, tx, kodeOPD)
	return helper.ToOpdResponses(opd)
}

func (service *OpdServiceImpl) FindKodeOpdUrusan(ctx context.Context, kodeOPD string) ([]web.OPD, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	opdUrusanBidangDomain, err := service.OpdRepository.FindKodeOpdUrusan(ctx, tx, kodeOPD)
	if err != nil {
		log.Println("Error mengambil data opd dan urusan bidang:", err)
		return nil, err
	}

	var opds []web.OPD
	opdMap := make(map[string]*web.OPD)

	for _, opdUrusanBidang := range opdUrusanBidangDomain {
		opd, err := service.OpdRepository.FindById(ctx, tx, opdUrusanBidang.KodeOpd.String)
		helper.PanicIfError(err)

		urusan, err := service.UrusanRepository.FindById(ctx, tx, opdUrusanBidang.KodeUrusan.String)
		helper.PanicIfError(err)

		bidangUrusan, err := service.BidangUrusanRepository.FindByBidangUrusan(ctx, tx, opdUrusanBidang.BidangUrusan.String)
		helper.PanicIfError(err)

		if _, exists := opdMap[opd.KodeOpd]; !exists {
			opdMap[opd.KodeOpd] = &web.OPD{
				KodeOpd: opd.KodeOpd,
				NamaOpd: opd.NamaOpd,
			}
		}

		opdEntry := opdMap[opd.KodeOpd]
		urusanExists := false

		for i, urusanEntry := range opdEntry.UrusanOPD {
			if urusanEntry.KodeUrusan == urusan.KodeUrusan {
				urusanExists = true
				opdEntry.UrusanOPD[i].BidangUrusanOPD = append(opdEntry.UrusanOPD[i].BidangUrusanOPD, web.BIDANGURUSANOPD{
					KodeBidangUrusan: bidangUrusan.KodeBidangUrusan,
					BidangUrusan:     bidangUrusan.BidangUrusan,
				})
				break
			}
		}

		if !urusanExists {
			opdEntry.UrusanOPD = append(opdEntry.UrusanOPD, web.URUSANOPD{
				Id:         urusan.Id,
				KodeUrusan: urusan.KodeUrusan,
				Urusan:     urusan.Urusan,
				BidangUrusanOPD: []web.BIDANGURUSANOPD{
					{
						Id:               bidangUrusan.ID,
						KodeBidangUrusan: bidangUrusan.KodeBidangUrusan,
						BidangUrusan:     bidangUrusan.BidangUrusan,
					},
				},
			})
		}
	}

	for _, opd := range opdMap {
		opds = append(opds, *opd)
	}

	return opds, nil
}
