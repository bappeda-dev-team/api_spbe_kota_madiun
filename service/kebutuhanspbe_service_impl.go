package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type KebutuhanSPBEServiceImpl struct {
	KebutuhanSPBERepository repository.KebutuhanSPBERepository
	DB                      *sql.DB
}

func NewKebutuhanSPBEService(kebutuhanSPBERepository repository.KebutuhanSPBERepository, DB *sql.DB) KebutuhanSPBEService {
	return &KebutuhanSPBEServiceImpl{
		KebutuhanSPBERepository: kebutuhanSPBERepository,
		DB:                      DB,
	}
}

func (service *KebutuhanSPBEServiceImpl) Create(ctx context.Context, request web.KebutuhanSPBECreateRequest) (web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE := domain.KebutuhanSPBE{
		KodeOpd:        request.KodeOpd,
		Tahun:          request.Tahun,
		NamaDomain:     request.NamaDomain,
		IdProsesbisnis: request.IdProsesbisnis,
		JenisKebutuhan: []domain.JenisKebutuhan{},
	}

	for _, jk := range request.JenisKebutuhan {
		jenisKebutuhan := domain.JenisKebutuhan{
			Kebutuhan:   jk.Kebutuhan,
			KondisiAwal: []domain.KondisiAwal{},
		}
		for _, ka := range jk.KondisiAwal {
			jenisKebutuhan.KondisiAwal = append(jenisKebutuhan.KondisiAwal, domain.KondisiAwal{
				Keterangan: ka.Keterangan,
				Tahun:      ka.Tahun,
			})
		}
		kebutuhanSPBE.JenisKebutuhan = append(kebutuhanSPBE.JenisKebutuhan, jenisKebutuhan)
	}

	kebutuhanSPBE, err = service.KebutuhanSPBERepository.Insert(ctx, tx, kebutuhanSPBE)
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}

	return helper.ToKebutuhanSPBEResponse(kebutuhanSPBE), nil
}

func (service *KebutuhanSPBEServiceImpl) Update(ctx context.Context, request web.KebutuhanSPBEUpdateRequest) (web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE, err := service.KebutuhanSPBERepository.FindById(ctx, tx, request.ID)
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}

	kebutuhanSPBE.KodeOpd = request.KodeOpd
	kebutuhanSPBE.Tahun = request.Tahun
	kebutuhanSPBE.NamaDomain = request.NamaDomain
	kebutuhanSPBE.IdProsesbisnis = request.IdProsesbisnis

	for _, jk := range request.JenisKebutuhan {
		jenisKebutuhan := domain.JenisKebutuhan{
			Id:          jk.ID,
			Kebutuhan:   jk.Kebutuhan,
			KondisiAwal: []domain.KondisiAwal{},
		}

		for _, ka := range jk.KondisiAwal {
			jenisKebutuhan.KondisiAwal = append(jenisKebutuhan.KondisiAwal, domain.KondisiAwal{
				Id:               ka.ID,
				JenisKebutuhanId: ka.JenisKebutuhanId,
				Keterangan:       ka.Keterangan,
				Tahun:            ka.Tahun,
			})
		}
		kebutuhanSPBE.JenisKebutuhan = append(kebutuhanSPBE.JenisKebutuhan, jenisKebutuhan)
	}

	kebutuhanSPBE, err = service.KebutuhanSPBERepository.Update(ctx, tx, kebutuhanSPBE)
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}

	return helper.ToKebutuhanSPBEResponse(kebutuhanSPBE), nil
}

func (service *KebutuhanSPBEServiceImpl) Delete(ctx context.Context, kebutuhanSPBEId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.KebutuhanSPBERepository.Delete(ctx, tx, kebutuhanSPBEId)
	if err != nil {
		return err
	}

	return nil
}

// fix by id
func (service *KebutuhanSPBEServiceImpl) FindById(ctx context.Context, kebutuhanSPBEId int) (web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE, err := service.KebutuhanSPBERepository.FindById(ctx, tx, kebutuhanSPBEId)
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}

	jenisKebutuhan, err := service.KebutuhanSPBERepository.FindJenisKebutuhanByKebutuhanId(ctx, tx, kebutuhanSPBE.ID)
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}

	var jenisKebutuhanResponses []web.JenisKebutuhanResponse
	for _, jk := range jenisKebutuhan {
		kondisiAwal, err := service.KebutuhanSPBERepository.FindKondisiAwalByJenisKebutuhanId(ctx, tx, jk.Id)
		if err != nil {
			return web.KebutuhanSPBEResponse{}, err
		}

		var kondisiAwalResponses []web.KondisiAwalResponse
		for _, ka := range kondisiAwal {
			kondisiAwalResponses = append(kondisiAwalResponses, web.KondisiAwalResponse{
				Id:               ka.Id,
				JenisKebutuhanId: ka.JenisKebutuhanId,
				Keterangan:       ka.Keterangan,
				Tahun:            ka.Tahun,
			})
		}

		jenisKebutuhanResponses = append(jenisKebutuhanResponses, web.JenisKebutuhanResponse{
			Id:          jk.Id,
			KebutuhanId: jk.KebutuhanId,
			Kebutuhan:   jk.Kebutuhan,
			KondisiAwal: kondisiAwalResponses,
		})
	}

	response := web.KebutuhanSPBEResponse{
		ID:             kebutuhanSPBE.ID,
		KodeOpd:        kebutuhanSPBE.KodeOpd,
		Tahun:          kebutuhanSPBE.Tahun,
		NamaDomain:     kebutuhanSPBE.NamaDomain,
		IdProsesbisnis: kebutuhanSPBE.IdProsesbisnis,
		JenisKebutuhan: jenisKebutuhanResponses,
	}

	return response, nil
}

func (service *KebutuhanSPBEServiceImpl) FindByKodeOpdAndTahun(ctx context.Context, kodeOpd string, tahun int) ([]web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBEs, err := service.KebutuhanSPBERepository.FindByKodeOpdAndTahun(ctx, tx, kodeOpd, tahun)
	if err != nil {
		return nil, err
	}

	var responses []web.KebutuhanSPBEResponse
	for _, kebutuhanSPBE := range kebutuhanSPBEs {
		jenisKebutuhan, err := service.KebutuhanSPBERepository.FindJenisKebutuhanByKebutuhanId(ctx, tx, kebutuhanSPBE.ID)
		if err != nil {
			return nil, err
		}

		var jenisKebutuhanResponses []web.JenisKebutuhanResponse
		for _, jk := range jenisKebutuhan {
			kondisiAwal, err := service.KebutuhanSPBERepository.FindKondisiAwalByJenisKebutuhanId(ctx, tx, jk.Id)
			if err != nil {
				return nil, err
			}

			var kondisiAwalResponses []web.KondisiAwalResponse
			for _, ka := range kondisiAwal {
				kondisiAwalResponses = append(kondisiAwalResponses, web.KondisiAwalResponse{
					Id:               ka.Id,
					JenisKebutuhanId: ka.JenisKebutuhanId,
					Keterangan:       ka.Keterangan,
					Tahun:            ka.Tahun,
				})
			}

			jenisKebutuhanResponses = append(jenisKebutuhanResponses, web.JenisKebutuhanResponse{
				Id:          jk.Id,
				KebutuhanId: jk.KebutuhanId,
				Kebutuhan:   jk.Kebutuhan,
				KondisiAwal: kondisiAwalResponses,
			})
		}

		responses = append(responses, web.KebutuhanSPBEResponse{
			ID:             kebutuhanSPBE.ID,
			KodeOpd:        kebutuhanSPBE.KodeOpd,
			Tahun:          kebutuhanSPBE.Tahun,
			NamaDomain:     kebutuhanSPBE.NamaDomain,
			IdProsesbisnis: kebutuhanSPBE.IdProsesbisnis,
			JenisKebutuhan: jenisKebutuhanResponses,
		})
	}

	return responses, nil
}
