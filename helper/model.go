package helper

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
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
		SasaranKota:      prosesbisnis.SasaranKota,
		KodeProsesBisnis: prosesbisnis.KodeProsesBisnis,
		BidangUrusan:     prosesbisnis.BidangUrusan,
		RabLevel1: web.ReferensiArsitekturResponse{
			Id: prosesbisnis.RabLevel1ID,
		},
		RabLevel2: web.ReferensiArsitekturResponse{
			Id: prosesbisnis.RabLevel2ID,
		},
		RabLevel3: web.ReferensiArsitekturResponse{
			Id: prosesbisnis.RabLevel3ID,
		},
		Tahun: prosesbisnis.Tahun,
	}
}

// func ToProsesBisnisResponses(prosbis []domain.ProsesBisnis) []web.ProsesBinsisRespons {
// 	var prosbisResponse []web.ProsesBinsisRespons
// 	for _, prosesbisnis := range prosbis {
// 		prosbisResponse = append(prosbisResponse, ToProsesBisnisResponse(prosesbisnis))
// 	}
// 	return prosbisResponse
// }
