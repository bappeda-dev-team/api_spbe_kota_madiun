package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type RencanaPelaksanaanServiceImpl struct {
	rencanaPelaksanaanRepository repository.RencanaPelaksanaanRepository
	sasaranKinerjaRepository     repository.SasaranKinerjaPegawaiRepository
	kebutuhanSPBERepository      repository.KebutuhanSPBERepository
	DB                           *sql.DB
}

func NewRencanaPelaksanaanServiceImpl(rencanaPelaksanaanRepository repository.RencanaPelaksanaanRepository, sasaranKinerjaRepository repository.SasaranKinerjaPegawaiRepository, kebutuhanSPBERepository repository.KebutuhanSPBERepository, DB *sql.DB) *RencanaPelaksanaanServiceImpl {
	return &RencanaPelaksanaanServiceImpl{
		rencanaPelaksanaanRepository: rencanaPelaksanaanRepository,
		sasaranKinerjaRepository:     sasaranKinerjaRepository,
		kebutuhanSPBERepository:      kebutuhanSPBERepository,
		DB:                           DB,
	}
}

func (service *RencanaPelaksanaanServiceImpl) Create(ctx context.Context, request web.RencanaPelaksanaanCreateRequest) (web.RencanaPelaksanaanResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kebutuhanSPBE, err := service.kebutuhanSPBERepository.FindById(ctx, tx, request.IdKebutuhan)
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, fmt.Errorf("kebutuhan SPBE tidak ditemukan: %v", err)
	}

	if !kebutuhanSPBE.PenanggungJawab.Valid || kebutuhanSPBE.PenanggungJawab.String != request.KodeOpd {
		return web.RencanaPelaksanaanResponse{}, fmt.Errorf("tidak diizinkan membuat rencana pelaksanaan untuk kebutuhan SPBE ini")
	}

	rencanaPelaksanaan := domain.RencanaPelaksanaanPegawai{
		IdKebutuhan:      request.IdKebutuhan,
		KodeOpd:          request.KodeOpd,
		IdSasaranKinerja: request.IdSasaranKinerja,
		IndikatorPD:      request.IndikatorPD,
		PerangkatDaerah:  request.PerangkatDaerah,
		TahunPelaksanaan: make([]domain.TahunPelaksanaan, len(request.TahunPelaksanaan)),
	}

	for i, tahun := range request.TahunPelaksanaan {
		rencanaPelaksanaan.TahunPelaksanaan[i] = domain.TahunPelaksanaan{Tahun: tahun.Tahun}
	}

	rencanaPelaksanaan, err = service.rencanaPelaksanaanRepository.Insert(ctx, tx, rencanaPelaksanaan)
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}

	return helper.ToRencanaPelaksanaanResponse(rencanaPelaksanaan), nil
}

func (service *RencanaPelaksanaanServiceImpl) Update(ctx context.Context, request web.RencanaPelaksanaanUpdateRequest) (web.RencanaPelaksanaanResponse, error) {
	log.Printf("Memulai proses update untuk ID: %d, KodeOpd: %s", request.Id, request.KodeOpd)
	tx, err := service.DB.Begin()
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rencanaPelaksanaan, err := service.rencanaPelaksanaanRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}

	kebutuhanSPBE, err := service.kebutuhanSPBERepository.FindById(ctx, tx, request.IdKebutuhan)
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, fmt.Errorf("kebutuhan SPBE tidak ditemukan: %v", err)
	}

	log.Printf("Data kebutuhanSPBE: %+v", kebutuhanSPBE)
	log.Printf("PenanggungJawab: %v, KodeOpd: %s", kebutuhanSPBE.PenanggungJawab, request.KodeOpd)
	if !kebutuhanSPBE.PenanggungJawab.Valid || kebutuhanSPBE.PenanggungJawab.String != request.KodeOpd {
		log.Printf("Gagal validasi PenanggungJawab")
		return web.RencanaPelaksanaanResponse{}, fmt.Errorf("tidak diizinkan edit data pada rencana pelaksanaan untuk kebutuhan SPBE ini")
	}

	rencanaPelaksanaan.KodeOpd = request.KodeOpd
	rencanaPelaksanaan.IdKebutuhan = request.IdKebutuhan
	rencanaPelaksanaan.IdSasaranKinerja = request.IdSasaranKinerja
	rencanaPelaksanaan.IndikatorPD = request.IndikatorPD
	rencanaPelaksanaan.PerangkatDaerah = request.PerangkatDaerah

	rencanaPelaksanaan.TahunPelaksanaan = []domain.TahunPelaksanaan{}
	for _, tp := range request.TahunPelaksanaan {
		tahunPelaksanaan := domain.TahunPelaksanaan{
			IdRencanaPelaksana: tp.IdRencanaPelaksana,
			Tahun:              tp.Tahun,
		}
		rencanaPelaksanaan.TahunPelaksanaan = append(rencanaPelaksanaan.TahunPelaksanaan, tahunPelaksanaan)
	}

	rencanaPelaksanaan, err = service.rencanaPelaksanaanRepository.Update(ctx, tx, rencanaPelaksanaan)
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}

	log.Printf("Update berhasil untuk ID: %d", rencanaPelaksanaan.Id)
	return helper.ToRencanaPelaksanaanResponse(rencanaPelaksanaan), nil
}

func (service *RencanaPelaksanaanServiceImpl) Delete(ctx context.Context, rencanaId int, kodeOpd string, role string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	rencanaPelaksanaan, err := service.rencanaPelaksanaanRepository.FindById(ctx, tx, rencanaId)
	helper.PanicIfError(err)

	if role == "admin_kota" {
		// Admin kota dapat menghapus semua data tanpa memeriksa kode OPD
		service.rencanaPelaksanaanRepository.Delete(ctx, tx, rencanaId)
	} else if role == "admin_opd" || role == "asn" {
		// Admin OPD dan ASN hanya dapat menghapus data berdasarkan kode OPD mereka
		if rencanaPelaksanaan.KodeOpd != kodeOpd {
			return errors.New("rencana pelaksanaan tidak ditemukan untuk OPD ini")
		}
		service.rencanaPelaksanaanRepository.Delete(ctx, tx, rencanaId)
	} else {
		return errors.New("role tidak memiliki izin untuk menghapus rencana pelaksanaan")
	}

	return nil
}

func (service *RencanaPelaksanaanServiceImpl) FindById(ctx context.Context, rencanaId int, kodeOpd string) (web.RencanaPelaksanaanResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rencanaPelaksanaan, err := service.rencanaPelaksanaanRepository.FindById(ctx, tx, rencanaId)
	if err != nil {
		return web.RencanaPelaksanaanResponse{}, err
	}

	if kodeOpd != "" && rencanaPelaksanaan.KodeOpd != kodeOpd {
		return web.RencanaPelaksanaanResponse{}, errors.New("rencana pelaksanaan tidak ditemukan untuk OPD ini")
	}

	// Mengambil data tahun pelaksanaan
	tahunPelaksanaan, err := service.rencanaPelaksanaanRepository.FindIdTahunPelaksanaan(ctx, tx, rencanaPelaksanaan.Id)
	helper.PanicIfError(err)

	// Mengonversi domain.TahunPelaksanaan ke web.TahunPelaksanaanResponse
	var tahunPelaksanaanResponses []web.TahunPelaksanaanResponse
	for _, tahun := range tahunPelaksanaan {
		tahunPelaksanaanResponses = append(tahunPelaksanaanResponses, web.TahunPelaksanaanResponse{
			Id:    tahun.Id,
			Tahun: tahun.Tahun,
		})
	}

	sasaranKinerja, err := service.sasaranKinerjaRepository.FindById(ctx, tx, rencanaPelaksanaan.IdSasaranKinerja)
	helper.PanicIfError(err)

	return web.RencanaPelaksanaanResponse{
		Id:          rencanaPelaksanaan.Id,
		KodeOpd:     rencanaPelaksanaan.KodeOpd,
		IdKebutuhan: rencanaPelaksanaan.IdKebutuhan,
		SasaranKinerja: web.SasaranKinerjaPegawaiResponse{
			Id:                    sasaranKinerja.Id,
			SasaranKinerjaPegawai: sasaranKinerja.SasaranKinerjaPegawai,
			AnggaranSasaran:       sasaranKinerja.AnggaranSasaran,
			KodeSubKegiatan:       sasaranKinerja.KodeSubKegiatan,
			SubKegiatan:           sasaranKinerja.SubKegiatan,
		},
		IndikatorPD:      rencanaPelaksanaan.IndikatorPD,
		PerangkatDaerah:  rencanaPelaksanaan.PerangkatDaerah,
		TahunPelaksanaan: tahunPelaksanaanResponses,
	}, nil
}

func (service *RencanaPelaksanaanServiceImpl) FindAll(ctx context.Context, kodeOpd string, kebutuhanId int) ([]web.RencanaPelaksanaanResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	rencanaPelaksanaanList, err := service.rencanaPelaksanaanRepository.FindAll(ctx, tx, kodeOpd, kebutuhanId)
	if err != nil {
		return nil, err
	}

	var responses []web.RencanaPelaksanaanResponse
	for _, rencana := range rencanaPelaksanaanList {

		// Mengambil data tahun pelaksanaan
		tahunPelaksanaan, err := service.rencanaPelaksanaanRepository.FindIdTahunPelaksanaan(ctx, tx, rencana.Id)
		helper.PanicIfError(err)

		// Mengonversi domain.TahunPelaksanaan ke web.TahunPelaksanaanResponse
		var tahunPelaksanaanResponses []web.TahunPelaksanaanResponse
		for _, tahun := range tahunPelaksanaan {
			tahunPelaksanaanResponses = append(tahunPelaksanaanResponses, web.TahunPelaksanaanResponse{
				Id:    tahun.Id,
				Tahun: tahun.Tahun,
			})
		}

		sasaranKinerja, err := service.sasaranKinerjaRepository.FindById(ctx, tx, rencana.IdSasaranKinerja)
		helper.PanicIfError(err)

		responses = append(responses, web.RencanaPelaksanaanResponse{
			Id:          rencana.Id,
			KodeOpd:     rencana.KodeOpd,
			IdKebutuhan: rencana.IdKebutuhan,
			SasaranKinerja: web.SasaranKinerjaPegawaiResponse{
				Id:                    sasaranKinerja.Id,
				SasaranKinerjaPegawai: sasaranKinerja.SasaranKinerjaPegawai,
				PelaksanaSasaran:      sasaranKinerja.PelaksanaSasaran,
				AnggaranSasaran:       sasaranKinerja.AnggaranSasaran,
				KodeSubKegiatan:       sasaranKinerja.KodeSubKegiatan,
				SubKegiatan:           sasaranKinerja.SubKegiatan,
			},
			IndikatorPD:      rencana.IndikatorPD,
			PerangkatDaerah:  rencana.PerangkatDaerah,
			TahunPelaksanaan: tahunPelaksanaanResponses,
		})
	}

	return responses, nil
}
