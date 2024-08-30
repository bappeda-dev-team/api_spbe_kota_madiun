package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"errors"
	"time"
)

type KebutuhanSPBEServiceImpl struct {
	KebutuhanSPBERepository repository.KebutuhanSPBERepository
	ProsesBisnisRepository  repository.ProsesBisnisRepository
	DB                      *sql.DB
}

func NewKebutuhanSPBEServiceImpl(kebutuhanSPBERepository repository.KebutuhanSPBERepository, prosesBisnisRepository repository.ProsesBisnisRepository, DB *sql.DB) *KebutuhanSPBEServiceImpl {
	return &KebutuhanSPBEServiceImpl{
		KebutuhanSPBERepository: kebutuhanSPBERepository,
		ProsesBisnisRepository:  prosesBisnisRepository,
		DB:                      DB,
	}
}

func (service *KebutuhanSPBEServiceImpl) Create(ctx context.Context, request web.KebutuhanSPBECreateRequest) (web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Menggunakan tahun sekarang jika tidak ada input
	tahun := time.Now().Year()
	if request.Tahun != 0 {
		tahun = request.Tahun
	}

	// Konversi NamaDomain ke pointer, null jika kosong
	namaDomain := sql.NullString{
		String: request.NamaDomain,
		Valid:  request.NamaDomain != "",
	}

	kebutuhanSPBE := domain.KebutuhanSPBE{
		Keterangan:     request.KeteranganGap,
		KodeOpd:        request.KodeOpd,
		Tahun:          tahun,
		NamaDomain:     namaDomain,
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
	kebutuhanSPBE.NamaDomain = sql.NullString{
		String: request.NamaDomain,
		Valid:  request.NamaDomain != "",
	}
	kebutuhanSPBE.IdProsesbisnis = request.IdProsesbisnis

	for _, jk := range request.JenisKebutuhan {
		jenisKebutuhan := domain.JenisKebutuhan{
			Id:          jk.ID,
			KebutuhanId: jk.JenisKebutuhanId,
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

func (service *KebutuhanSPBEServiceImpl) UpdateKeterangan(ctx context.Context, request web.KebutuhanSPBEKeteranganUpdateRequest) (web.KebutuhanSPBEKeteranganResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEKeteranganResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE, err := service.KebutuhanSPBERepository.FindById(ctx, tx, request.ID)
	if err != nil {
		return web.KebutuhanSPBEKeteranganResponse{}, err
	}

	kebutuhanSPBE.Keterangan = request.Keterangan
	kebutuhanSPBE, err = service.KebutuhanSPBERepository.UpdateKeterangan(ctx, tx, kebutuhanSPBE)
	helper.PanicIfError(err)

	return helper.ToKebutuhanSPBEKeteranganResponse(kebutuhanSPBE), nil
}

func (service *KebutuhanSPBEServiceImpl) UpdatePenanggungJawab(ctx context.Context, request web.KebutuhanSPBEPjUpdateRequest) (web.KebutuhanSPBEPjResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEPjResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE, err := service.KebutuhanSPBERepository.FindById(ctx, tx, request.ID)
	if err != nil {
		return web.KebutuhanSPBEPjResponse{}, err
	}

	kebutuhanSPBE.IndikatorPj = sql.NullString{
		String: request.IndikatorPj,
		Valid:  request.IndikatorPj != "",
	}

	// Jika indikator internal, gunakan kode_opd dari context
	if request.IndikatorPj == "internal" {
		kodeOPD, ok := ctx.Value("kode_opd").(string)
		if !ok {
			return web.KebutuhanSPBEPjResponse{}, errors.New("kode OPD tidak ditemukan dalam context")
		}
		kebutuhanSPBE.PenanggungJawab = sql.NullString{
			String: kodeOPD,
			Valid:  true,
		}
	} else {
		// Jika bukan internal, gunakan penanggung jawab dari request
		kebutuhanSPBE.PenanggungJawab = sql.NullString{
			String: request.PenanggungJawab,
			Valid:  request.PenanggungJawab != "",
		}
	}

	kebutuhanSPBE, err = service.KebutuhanSPBERepository.UpdatePenanggungJawab(ctx, tx, kebutuhanSPBE)
	if err != nil {
		return web.KebutuhanSPBEPjResponse{}, err
	}

	return helper.ToKebutuhanSPBEPjResponse(kebutuhanSPBE), nil
}

func (service *KebutuhanSPBEServiceImpl) Delete(ctx context.Context, kebutuhanSPBEId int, kodeOPD string, role string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kebutuhanspbe, err := service.KebutuhanSPBERepository.FindById(ctx, tx, kebutuhanSPBEId)
	helper.PanicIfError(err)

	if role == "admin_kota" {
		// Admin kota dapat menghapus semua data tanpa memeriksa kode OPD
		service.KebutuhanSPBERepository.Delete(ctx, tx, kebutuhanSPBEId)
	} else if role == "admin_opd" || role == "asn" {
		// Admin OPD dan ASN hanya dapat menghapus data berdasarkan kode OPD mereka
		if kebutuhanspbe.KodeOpd != kodeOPD {
			return errors.New("kebutuhan spbe tidak ditemukan untuk OPD ini")
		}
		service.KebutuhanSPBERepository.Delete(ctx, tx, kebutuhanSPBEId)
	} else {
		return errors.New("role tidak memiliki izin untuk menghapus kebutuhan spbe")
	}

	return nil
}

// fix by id
func (service *KebutuhanSPBEServiceImpl) FindById(ctx context.Context, kebutuhanSPBEId int, kodeOpd string) (web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE, err := service.KebutuhanSPBERepository.FindById(ctx, tx, kebutuhanSPBEId)
	if err != nil {
		return web.KebutuhanSPBEResponse{}, err
	}

	if kodeOpd != "" && kebutuhanSPBE.KodeOpd != kodeOpd {
		return web.KebutuhanSPBEResponse{}, errors.New("kebutuhan spbe tidak ditemukan untuk OPD ini")
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

	prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, kebutuhanSPBE.IdProsesbisnis)
	helper.PanicIfError(err)

	response := web.KebutuhanSPBEResponse{
		ID:            kebutuhanSPBE.ID,
		KeteranganGap: kebutuhanSPBE.Keterangan,
		KodeOpd:       kebutuhanSPBE.KodeOpd,
		Tahun:         kebutuhanSPBE.Tahun,
		NamaDomain:    kebutuhanSPBE.NamaDomain.String,
		ProsesBisnis: web.ProsesBisnisResponse{
			ID:               prosesBisnis.ID,
			NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
		},
		JenisKebutuhan:  jenisKebutuhanResponses,
		IndikatorPj:     kebutuhanSPBE.IndikatorPj.String,
		PenanggungJawab: kebutuhanSPBE.PenanggungJawab.String,
	}

	return response, nil
}

func (service *KebutuhanSPBEServiceImpl) FindByKodeOpdAndTahun(ctx context.Context, kodeOpd string, tahun int, prosesbisnis int) ([]web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBEs, err := service.KebutuhanSPBERepository.FindByKodeOpdAndTahun(ctx, tx, kodeOpd, tahun, prosesbisnis)
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

		prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, kebutuhanSPBE.IdProsesbisnis)
		helper.PanicIfError(err)

		responses = append(responses, web.KebutuhanSPBEResponse{
			ID:            kebutuhanSPBE.ID,
			KeteranganGap: kebutuhanSPBE.Keterangan,
			KodeOpd:       kebutuhanSPBE.KodeOpd,
			Tahun:         kebutuhanSPBE.Tahun,
			NamaDomain:    kebutuhanSPBE.NamaDomain.String,
			ProsesBisnis: web.ProsesBisnisResponse{
				ID:               prosesBisnis.ID,
				NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
			},
			JenisKebutuhan:  jenisKebutuhanResponses,
			IndikatorPj:     kebutuhanSPBE.IndikatorPj.String,
			PenanggungJawab: kebutuhanSPBE.PenanggungJawab.String,
		})
	}

	return responses, nil
}

func (service *KebutuhanSPBEServiceImpl) FindDataPemenuhanKebutuhan(ctx context.Context, kodeOpd string, tahun int, prosesbisnis int) ([]web.KebutuhanSPBEResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBEs, err := service.KebutuhanSPBERepository.FindByKodeOpdAndTahun(ctx, tx, kodeOpd, tahun, prosesbisnis)
	if err != nil {
		return nil, err
	}

	var responses []web.KebutuhanSPBEResponse
	for _, kebutuhanSPBE := range kebutuhanSPBEs {
		jenisKebutuhan, err := service.KebutuhanSPBERepository.FindJenisKebutuhanByKebutuhanId(ctx, tx, kebutuhanSPBE.ID)
		if err != nil {
			return nil, err
		}

		// Jika tidak ada jenis kebutuhan, lanjutkan ke iterasi berikutnya
		if len(jenisKebutuhan) == 0 {
			continue
		}

		var jenisKebutuhanResponses []web.JenisKebutuhanResponse
		for _, jk := range jenisKebutuhan {
			kondisiAwal, err := service.KebutuhanSPBERepository.FindKondisiAwalByJenisKebutuhanId(ctx, tx, jk.Id)
			if err != nil {
				return nil, err
			}

			var kondisiAwalResponses []web.KondisiAwalResponse
			for _, ka := range kondisiAwal {
				if ka.Keterangan != "" {
					kondisiAwalResponses = append(kondisiAwalResponses, web.KondisiAwalResponse{
						Id:               ka.Id,
						JenisKebutuhanId: ka.JenisKebutuhanId,
						Keterangan:       ka.Keterangan,
						Tahun:            ka.Tahun,
					})
				}
			}

			if len(kondisiAwalResponses) > 0 {
				jenisKebutuhanResponses = append(jenisKebutuhanResponses, web.JenisKebutuhanResponse{
					Id:          jk.Id,
					KebutuhanId: jk.KebutuhanId,
					Kebutuhan:   jk.Kebutuhan,
					KondisiAwal: kondisiAwalResponses,
				})
			}
		}

		// Jika tidak ada jenis kebutuhan yang valid, lanjutkan ke iterasi berikutnya
		if len(jenisKebutuhanResponses) == 0 {
			continue
		}

		prosesBisnis, err := service.ProsesBisnisRepository.FindById(ctx, tx, kebutuhanSPBE.IdProsesbisnis)
		helper.PanicIfError(err)

		response := web.KebutuhanSPBEResponse{
			ID:            kebutuhanSPBE.ID,
			KeteranganGap: kebutuhanSPBE.Keterangan,
			KodeOpd:       kebutuhanSPBE.KodeOpd,
			Tahun:         kebutuhanSPBE.Tahun,
			ProsesBisnis: web.ProsesBisnisResponse{
				ID:               prosesBisnis.ID,
				NamaProsesBisnis: prosesBisnis.NamaProsesBisnis,
			},
			JenisKebutuhan:  jenisKebutuhanResponses,
			IndikatorPj:     kebutuhanSPBE.IndikatorPj.String,
			PenanggungJawab: kebutuhanSPBE.PenanggungJawab.String,
		}

		// Hanya tambahkan NamaDomain jika valid
		if kebutuhanSPBE.NamaDomain.Valid && kebutuhanSPBE.NamaDomain.String != "" {
			response.NamaDomain = kebutuhanSPBE.NamaDomain.String
		}

		responses = append(responses, response)
	}

	return responses, nil
}
