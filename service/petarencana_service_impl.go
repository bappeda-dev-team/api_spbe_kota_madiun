package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"log"
)

type PetarencanaServiceImpl struct {
	PetarencanaRepository        repository.PetarencanaRepository
	KebutuhanspbeRepository      repository.KebutuhanSPBERepository
	RencanaPelaksanaanRepository repository.RencanaPelaksanaanRepository
	SasaranKinerjaRepository     repository.SasaranKinerjaPegawaiRepository
	OpdRepository                repository.OpdRepository
	DB                           *sql.DB
}

func NewPetarencanaServiceImpl(petarencanaRepository repository.PetarencanaRepository, kebutuhanspbeRepository repository.KebutuhanSPBERepository, rencanaPelaksanaanRepository repository.RencanaPelaksanaanRepository, sasaranKinerjaRepository repository.SasaranKinerjaPegawaiRepository, opdRepository repository.OpdRepository, DB *sql.DB) *PetarencanaServiceImpl {
	return &PetarencanaServiceImpl{
		PetarencanaRepository:        petarencanaRepository,
		KebutuhanspbeRepository:      kebutuhanspbeRepository,
		RencanaPelaksanaanRepository: rencanaPelaksanaanRepository,
		SasaranKinerjaRepository:     sasaranKinerjaRepository,
		OpdRepository:                opdRepository,
		DB:                           DB,
	}
}

func (service *PetarencanaServiceImpl) FindAll(ctx context.Context, kodeOpd string, tahun int) ([]web.PetarencanaRespons, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("Gagal memulai transaksi: %v", err)
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	petarencanas, err := service.PetarencanaRepository.FindAll(ctx, tx, kodeOpd, tahun)
	if err != nil {
		log.Printf("Gagal menemukan petarencana: %v", err)
		return nil, err
	}

	petarencanaMap := make(map[int]web.PetarencanaRespons)
	kebutuhanMap := make(map[int]web.PjKebutuhanSPBEResponse)

	for _, petarencana := range petarencanas {
		petarencanaResponse, exists := petarencanaMap[petarencana.ID]
		if !exists {
			petarencanaResponse = web.PetarencanaRespons{
				GeneratedID:      petarencana.GeneratedID,
				ID:               petarencana.ID,
				KodeOpd:          petarencana.KodeOpd,
				Tahun:            petarencana.Tahun,
				NamaProsesBisnis: petarencana.NamaProsesBisnis,
				KodeProsesBisnis: petarencana.KodeProsesBisnis,
				Layanan:          []web.RencanaLayanan{},
				DataDanInformasi: []web.RencanaDataDanInformasi{},
				Aplikasi:         []web.RencanaAplikasi{},
				Keterangan:       []web.PjKebutuhanSPBEResponse{},
			}
		}

		if len(petarencana.Keterangan) > 0 && petarencana.Keterangan[0].IdKeterangan.Valid {
			idKeterangan := int(petarencana.Keterangan[0].IdKeterangan.Int32)
			_, exists := kebutuhanMap[idKeterangan]

			if !exists {
				// Proses kebutuhan SPBE baru
				kebutuhanSPBE, err := service.KebutuhanspbeRepository.FindIdForPetarencana(ctx, tx, idKeterangan, kodeOpd, kodeOpd)
				if err != nil {
					log.Printf("Error saat mencari kebutuhan SPBE: %v", err)
					return nil, err
				}
				if kebutuhanSPBE.ID != 0 {
					jenisKebutuhan, err := service.KebutuhanspbeRepository.FindJenisKebutuhanByKebutuhanId(ctx, tx, kebutuhanSPBE.ID)
					if err != nil {
						log.Printf("Gagal menemukan jenis kebutuhan: %v", err)
						return nil, err
					}

					var jenisKebutuhanResponses []web.JenisKebutuhanResponse
					for _, jk := range jenisKebutuhan {
						kondisiAwal, err := service.KebutuhanspbeRepository.FindKondisiAwalByJenisKebutuhanId(ctx, tx, jk.Id)
						if err != nil {
							log.Printf("Gagal menemukan kondisi awal: %v", err)
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

					rencanaPelaksanaan, err := service.RencanaPelaksanaanRepository.FindByKebutuhanId(ctx, tx, kebutuhanSPBE.ID)
					if err != nil {
						log.Printf("Gagal menemukan rencana pelaksanaan: %v", err)
						return nil, err
					}

					opdPetarencana, err := service.OpdRepository.FindById(ctx, tx, kebutuhanSPBE.PenanggungJawab.String)
					helper.PanicIfError(err)

					var rencanaPelaksanaanResponses []web.RencanaPelaksanaanResponse
					for _, rp := range rencanaPelaksanaan {
						if rp.KodeOpd == kebutuhanSPBE.PenanggungJawab.String {
							tahunPelaksanaan, err := service.RencanaPelaksanaanRepository.FindIdTahunPelaksanaan(ctx, tx, rp.Id)
							helper.PanicIfError(err)
							var tahunPelaksanaanResponses []web.TahunPelaksanaanResponse
							for _, tp := range tahunPelaksanaan {
								tahunPelaksanaanResponses = append(tahunPelaksanaanResponses, web.TahunPelaksanaanResponse{
									Id:    tp.Id,
									Tahun: tp.Tahun,
								})
							}

							sasaranKinerja, err := service.SasaranKinerjaRepository.FindById(ctx, tx, rp.IdSasaranKinerja)
							helper.PanicIfError(err)

							rencanaPelaksanaanResponses = append(rencanaPelaksanaanResponses, web.RencanaPelaksanaanResponse{
								Id:          rp.Id,
								KodeOpd:     rp.KodeOpd,
								IdKebutuhan: rp.IdKebutuhan,
								SasaranKinerja: web.SasaranKinerjaPegawaiResponse{
									Id:                    sasaranKinerja.Id,
									KodeOpd:               sasaranKinerja.KodeOpd,
									KodeSasaran:           sasaranKinerja.KodeSasaran,
									Tahun:                 sasaranKinerja.Tahun,
									SasaranKinerjaPegawai: sasaranKinerja.SasaranKinerjaPegawai,
									AnggaranSasaran:       sasaranKinerja.AnggaranSasaran,
									PelaksanaSasaran:      sasaranKinerja.PelaksanaSasaran,
									KodeSubKegiatan:       sasaranKinerja.KodeSubKegiatan,
									SubKegiatan:           sasaranKinerja.SubKegiatan,
								},
								TahunPelaksanaan: tahunPelaksanaanResponses,
							})
						}
					}

					opdKebutuhan, err := service.OpdRepository.FindById(ctx, tx, kebutuhanSPBE.KodeOpd)
					helper.PanicIfError(err)

					kebutuhanSpbeRespons := web.PjKebutuhanSPBEResponse{
						ID:            kebutuhanSPBE.ID,
						KeteranganGap: kebutuhanSPBE.Keterangan,
						KodeOpd:       kebutuhanSPBE.KodeOpd,
						PerangkatDaerah: web.OpdRespons{
							KodeOpd: opdKebutuhan.KodeOpd,
							NamaOpd: opdKebutuhan.NamaOpd,
						},
						Tahun:          kebutuhanSPBE.Tahun,
						NamaDomain:     kebutuhanSPBE.NamaDomain.String,
						ProsesBisnis:   web.ProsesBisnisResponse{ID: kebutuhanSPBE.IdProsesbisnis, NamaProsesBisnis: petarencana.NamaProsesBisnis},
						JenisKebutuhan: jenisKebutuhanResponses,
						IndikatorPj:    kebutuhanSPBE.IndikatorPj.String,
						PenanggungJawab: web.OpdRespons{
							KodeOpd: opdPetarencana.KodeOpd,
							NamaOpd: opdPetarencana.NamaOpd},
						RencanaPelaksanaan: rencanaPelaksanaanResponses,
					}

					kebutuhanMap[idKeterangan] = kebutuhanSpbeRespons
					petarencanaResponse.Keterangan = append(petarencanaResponse.Keterangan, kebutuhanSpbeRespons)
				}
			}
		}

		if petarencana.Layanan != nil {
			for _, l := range petarencana.Layanan {
				var webNullString web.NullString

				if l.NamaLayanan.Valid {
					webNullString.String = l.NamaLayanan.String
					webNullString.Valid = true
				}

				tempLayanan := web.RencanaLayanan{NamaLayanan: webNullString}
				if !helper.ContainsLayananPetarencana(petarencanaResponse.Layanan, tempLayanan) {
					petarencanaResponse.Layanan = append(petarencanaResponse.Layanan, tempLayanan)
				}
			}
		}

		if petarencana.DataDanInformasi != nil {
			for _, d := range petarencana.DataDanInformasi {
				var webNullString web.NullString

				if d.NamaData.Valid {
					webNullString.String = d.NamaData.String
					webNullString.Valid = true
				}

				tempData := web.RencanaDataDanInformasi{NamaData: webNullString}
				if !helper.ContainDataPetarencana(petarencanaResponse.DataDanInformasi, tempData) {
					petarencanaResponse.DataDanInformasi = append(petarencanaResponse.DataDanInformasi, tempData)
				}
			}
		}

		if petarencana.Aplikasi != nil {
			for _, a := range petarencana.Aplikasi {
				var webNullString web.NullString

				if a.NamaAplikasi.Valid {
					webNullString.String = a.NamaAplikasi.String
					webNullString.Valid = true
				}

				tempAplikasi := web.RencanaAplikasi{NamaAplikasi: webNullString}
				if !helper.ContainAplikasiPetarencana(petarencanaResponse.Aplikasi, tempAplikasi) {
					petarencanaResponse.Aplikasi = append(petarencanaResponse.Aplikasi, tempAplikasi)
				}
			}
		}

		petarencanaMap[petarencana.ID] = petarencanaResponse
	}

	var petarencanaResponses []web.PetarencanaRespons
	for _, response := range petarencanaMap {
		petarencanaResponses = append(petarencanaResponses, response)
	}

	return petarencanaResponses, nil
}
