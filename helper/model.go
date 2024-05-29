package helper

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
)

func ToReferensiArsitekturResponse(referensiarsitektur domain.ReferensiArsitektur) web.ReferensiArsitekturResponse{
	return web.ReferensiArsitekturResponse{
		Id: referensiarsitektur.IdReferensi,
		Kode_referensi: referensiarsitektur.Kode_referensi,
		Nama_referensi: referensiarsitektur.Nama_referensi,
		Level_referensi: referensiarsitektur.Level_referensi,
		Jenis_referensi: referensiarsitektur.Jenis_referensi,
		Created_at: referensiarsitektur.Created_at,
		Updated_at: referensiarsitektur.Updated_at,
	}
}

func ToReferenceResponses(reference []domain.ReferensiArsitektur) []web.ReferensiArsitekturResponse {
	var referenceResponses []web.ReferensiArsitekturResponse
	for _, referensiarsitektur := range reference {
		referenceResponses = append(referenceResponses, ToReferensiArsitekturResponse(referensiarsitektur))
	}
	return referenceResponses
}

// func ToProsesBisnisResponse(prosesbisnis domain.ProsesBisnis)web.ProsesBinsisRespons{
// 	return web.ProsesBinsisRespons{
// 		Id: prosesbisnis.Id,
// 		// Nama_proses_bisnis: prosesbisnis.Nama_proses_bisnis,
// 		// Sasaran_kota: prosesbisnis.Sasaran_kota,
// 		// Kode_proses_bisnis: prosesbisnis.Kode_proses_bisnis,
// 		// Kode_opd: prosesbisnis.Kode_opd,
// 		// Bidang_urusan: prosesbisnis.Bidang_urusan,
// 		// Rab_level_1: prosesbisnis.Rab_level_1,
// 		// Rab_level_2: prosesbisnis.Rab_level_2,
// 		// Rab_level_3: prosesbisnis.Rab_level_3,
// 		// Rab_level_4: prosesbisnis.Rab_level_4,
// 		// Rab_level_5: prosesbisnis.Rab_level_5,
// 		// Rab_level_6: prosesbisnis.Rab_level_6,
// 	}
// }

// func ToProsesBisnisResponses(prosbis []domain.ProsesBisnis) []web.ProsesBinsisRespons {
// 	var prosbisResponse []web.ProsesBinsisRespons
// 	for _, prosesbisnis := range prosbis {
// 		prosbisResponse = append(prosbisResponse, ToProsesBisnisResponse(prosesbisnis))
// 	}
// 	return prosbisResponse
// }