package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
)

type ProsesBisnisServiceImpl struct {
    ProsesBisnisRepository          repository.ProsesBisnisRepository
    ReferensiArsitekturRepository   repository.ReferensiArsitekturRepository
    DB                              *sql.DB
}

func NewProsesBisnisService(referensiarsitekturRepository repository.ReferensiArsitekturRepository, prosesbisnisRepository repository.ProsesBisnisRepository, DB *sql.DB) ProsesBisnisService {
    return &ProsesBisnisServiceImpl{
        ProsesBisnisRepository:        prosesbisnisRepository,
        ReferensiArsitekturRepository: referensiarsitekturRepository,
        DB:                            DB,
    }
}

func (service *ProsesBisnisServiceImpl) GetProsesBisnis(ctx context.Context) ([]web.ProsesBisnisRespons, error) {
    tx, err := service.DB.Begin()
    if err != nil {
        return nil, err
    }
    defer helper.CommitOrRollback(tx)

    prosesBisnisList, err := service.ProsesBisnisRepository.FindAll(ctx, tx)
    if err != nil {
        return nil, err
    }

    var responses []web.ProsesBisnisRespons
    for _, prosesBisnis := range prosesBisnisList {
        rabLevel1, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel1ID)
        if err != nil {
            return nil, err
        }

        rabLevel2, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel2ID)
        if err != nil {
            return nil, err
        }

        rabLevel3, err := service.ReferensiArsitekturRepository.FindById(ctx, tx, prosesBisnis.RabLevel3ID)
        if err != nil {
            return nil, err
        }

        response := web.ProsesBisnisRespons{
            ID:               prosesBisnis.ID,
            NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
            SasaranKota:      prosesBisnis.SasaranKota,
            KodeProsesBisnis: prosesBisnis.KodeProsesBisnis,
            KodeOPD:          prosesBisnis.KodeOPD,
            BidangUrusan:      prosesBisnis.BidangUrusan,
            RabLevel1:       web.ReferensiArsitekturResponse{
				Id: rabLevel1.IdReferensi,
				Kode_referensi: rabLevel1.Kode_referensi,
				Nama_referensi: rabLevel1.Nama_referensi,
				Level_referensi: rabLevel1.Level_referensi,
				Jenis_referensi: rabLevel1.Jenis_referensi,
				Tahun: rabLevel1.Tahun,
				Created_at: rabLevel1.Created_at,
				Updated_at: rabLevel1.Updated_at,
			},
            RabLevel2:         web.ReferensiArsitekturResponse{
				Id: rabLevel2.IdReferensi,
				Kode_referensi: rabLevel2.Kode_referensi,
				Nama_referensi: rabLevel2.Nama_referensi,
				Level_referensi: rabLevel2.Level_referensi,
				Jenis_referensi: rabLevel2.Jenis_referensi,
				Tahun: rabLevel2.Tahun,
				Created_at: rabLevel2.Created_at,
				Updated_at: rabLevel2.Updated_at,
			},
            RabLevel3:         web.ReferensiArsitekturResponse{
				Id: rabLevel2.IdReferensi,
				Kode_referensi: rabLevel3.Kode_referensi,
				Nama_referensi: rabLevel3.Nama_referensi,
				Level_referensi: rabLevel3.Level_referensi,
				Jenis_referensi: rabLevel3.Jenis_referensi,
				Tahun: rabLevel3.Tahun,
				Created_at: rabLevel3.Created_at,
				Updated_at: rabLevel3.Updated_at,
			},
            Tahun:            prosesBisnis.Tahun,
            CreatedAt:        prosesBisnis.CreatedAt,
            UpdatedAt:        prosesBisnis.UpdatedAt,
        }
        responses = append(responses, response)
    }

    return responses, nil
}
