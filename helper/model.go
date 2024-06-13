package helper

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"database/sql"
)

func ToReferensiArsitekturResponse(referensiarsitektur domain.ReferensiArsitektur) web.ReferensiArsitekturResponse {
	return web.ReferensiArsitekturResponse{
		Id:              referensiarsitektur.IdReferensi,
		Kode_referensi:  referensiarsitektur.Kode_referensi,
		Nama_referensi:  referensiarsitektur.Nama_referensi,
		Level_referensi: referensiarsitektur.Level_referensi,
		Jenis_referensi: referensiarsitektur.Jenis_referensi,
		Created_at:      referensiarsitektur.Created_at,
		Updated_at:      referensiarsitektur.Updated_at,
	}
}

func ToReferenceResponses(reference []domain.ReferensiArsitektur) []web.ReferensiArsitekturResponse {
	var referenceResponses []web.ReferensiArsitekturResponse
	for _, referensiarsitektur := range reference {
		referenceResponses = append(referenceResponses, ToReferensiArsitekturResponse(referensiarsitektur))
	}
	return referenceResponses
}

func ToProsesBisnisResponse(prosesbisnis domain.ProsesBisnis) web.ProsesBisnisRespons {
	return web.ProsesBisnisRespons{
		ID:               prosesbisnis.ID,
		KodeOPD:          prosesbisnis.KodeOPD,
		NamaProsesBisnis: prosesbisnis.NamaProsesBisnis,
		SasaranKota: &web.ProsbisSasaranKotaRespons{
			ID: nullInt32ToInt(prosesbisnis.SasaranKotaId),
		},
		KodeProsesBisnis: prosesbisnis.KodeProsesBisnis,
		BidangUrusan: &web.ProsBisBidangUrusanRespons{
			Id: nullInt32ToInt(prosesbisnis.BidangUrusanId),
		},
		RabLevel1: &web.ProsBisReferensiArsitekturRespons{
			Id: nullInt32ToInt(prosesbisnis.RabLevel1ID),
		},
		RabLevel2: &web.ProsBisReferensiArsitekturRespons{
			Id: nullInt32ToInt(prosesbisnis.RabLevel2ID),
		},
		RabLevel3: &web.ProsBisReferensiArsitekturRespons{
			Id: nullInt32ToInt(prosesbisnis.RabLevel3ID),
		},
		RabLevel4: &web.ProsBisPohonKinerjaRespons{
			ID: nullInt32ToInt(prosesbisnis.RabLevel4ID),
		},
		RabLevel5: &web.ProsBisPohonKinerjaRespons{
			ID: nullInt32ToInt(prosesbisnis.RabLevel5ID),
		},
		RabLevel6: &web.ProsBisPohonKinerjaRespons{
			ID: nullInt32ToInt(prosesbisnis.RabLevel6ID),
		},
		Tahun: prosesbisnis.Tahun,
	}
}

func ToSasaranKotaResponse(sasaran domain.SasaranKota) web.SasaranKotaRespons {
	return web.SasaranKotaRespons{
		ID:           sasaran.ID,
		Sasaran:      sasaran.Sasaran,
		TujuanKota:   sasaran.TujuanKota,
		StrategiKota: sasaran.StrategiKota,
		Tahun:        sasaran.Tahun,
		CreatedAt:    sasaran.CreatedAt,
		UpdatedAt:    sasaran.UpdatedAt,
	}
}

func ToPohonKinerjaResponse(pohon domain.PohonKinerja) web.PohonKinerjaRespons {
	return web.PohonKinerjaRespons{
		ID:         pohon.ID,
		NamaPohon:  pohon.NamaPohon,
		JenisPohon: pohon.JenisPohon,
		LevelPohon: pohon.LevelPohon,
		KodeOpd:    pohon.KodeOpd,
		Tahun:      pohon.Tahun,
		CreatedAt:  pohon.CreatedAt,
		UpdatedAt:  pohon.UpdatedAt,
	}
}

func ToBidangUrusanResponse(bidang domain.BidangUrusan) web.BidangUrusanRespons {
	return web.BidangUrusanRespons{
		Id:               bidang.ID,
		KodeBidangUrusan: bidang.KodeBidangUrusan,
		BidangUrusan:     bidang.BidangUrusan,
	}
}

// get all
func ToSasaranResponses(sasaran []domain.SasaranKota) []web.SasaranKotaRespons {
	var sasaranKotaResponses []web.SasaranKotaRespons
	for _, sasaranKota := range sasaran {
		sasaranKotaResponses = append(sasaranKotaResponses, ToSasaranKotaResponse(sasaranKota))
	}
	return sasaranKotaResponses
}

func ToBidangResponses(bidang []domain.BidangUrusan) []web.BidangUrusanRespons {
	var bidangUrusanRespons []web.BidangUrusanRespons
	for _, bidangUrusan := range bidang {
		bidangUrusanRespons = append(bidangUrusanRespons, ToBidangUrusanResponse(bidangUrusan))
	}
	return bidangUrusanRespons
}

func ToPohonResponses(pohon []domain.PohonKinerja) []web.PohonKinerjaRespons {
	var pohonRespons []web.PohonKinerjaRespons
	for _, pohonKinerja := range pohon {
		pohonRespons = append(pohonRespons, ToPohonKinerjaResponse(pohonKinerja))
	}
	return pohonRespons
}

func nullInt32ToInt(n sql.NullInt32) int {
	if n.Valid {
		return int(n.Int32)
	}
	return 0
}
