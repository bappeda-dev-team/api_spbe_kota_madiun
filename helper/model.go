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
