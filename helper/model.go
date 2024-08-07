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

func ToDomainSPBEResponse(domainSPBE domain.DomainSPBE) web.DomainSPBEResponse {
	return web.DomainSPBEResponse{
		Id:         domainSPBE.Id,
		NamaDomain: domainSPBE.NamaDomain,
		KodeDomain: domainSPBE.KodeDomain,
		Tahun:      domainSPBE.Tahun,
	}
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
			ID: nullInt32ToInt(prosesbisnis.StrategicId),
		},
		RabLevel5: &web.ProsBisPohonKinerjaRespons{
			ID: nullInt32ToInt(prosesbisnis.TacticalId),
		},
		RabLevel6: &web.ProsBisPohonKinerjaRespons{
			ID: nullInt32ToInt(prosesbisnis.OperationalId),
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

func ToLayananSpbeRespons(layananSpbe domain.LayananSPBE) web.LayananSpbeRespons {
	return web.LayananSpbeRespons{
		Id:          layananSpbe.Id,
		NamaLayanan: layananSpbe.NamaLayanan,
		KodeLayanan: layananSpbe.KodeLayanan,
		TujuanLayananId: &web.LayananspbePohonRespons{
			ID: nullInt32ToInt(layananSpbe.TujuanLayananId),
		},
		FungsiLayanan:     layananSpbe.FungsiLayanan,
		Tahun:             layananSpbe.Tahun,
		KodeOPD:           layananSpbe.KodeOPD,
		KementrianTerkait: layananSpbe.KementrianTerkait,
		MetodeLayanan:     layananSpbe.MetodeLayanan,
		RalLevel1id: &web.LayananSpbeReferensiArsitekturRespons{
			Id: nullInt32ToInt(layananSpbe.RalLevel1id),
		},
		RalLevel2id: &web.LayananSpbeReferensiArsitekturRespons{
			Id: nullInt32ToInt(layananSpbe.RalLevel2id),
		},
		RalLevel3id: &web.LayananSpbeReferensiArsitekturRespons{
			Id: nullInt32ToInt(layananSpbe.RalLevel3id),
		},
		RalLevel4id: &web.LayananSpbeReferensiArsitekturRespons{
			Id: nullInt32ToInt(layananSpbe.RalLevel4id),
		},
		StrategicId: &web.LayananspbePohonRespons{
			ID: nullInt32ToInt(layananSpbe.StrategicId),
		},
		TacticalId: &web.LayananspbePohonRespons{
			ID: nullInt32ToInt(layananSpbe.TacticalId),
		},
		OperationalId: &web.LayananspbePohonRespons{
			ID: nullInt32ToInt(layananSpbe.OperationalId),
		},
	}
}

func ToDataDanInformasiRespons(data domain.DataDanInformasi) web.DataDanInformasiRespons {
	return web.DataDanInformasiRespons{
		Id:                     data.Id,
		NamaData:               data.NamaData,
		SifatData:              data.SifatData,
		JenisData:              data.JenisData,
		ProdusenData:           data.ProdusenData,
		UraianData:             data.UraianData,
		ValiditasData:          data.ValiditasData,
		PjData:                 data.PjData,
		KodeOPD:                data.KodeOPD,
		InformasiTerkaitInput:  data.InformasiTerkaitInput,
		InformasiTerkaitOutput: data.InformasiTerkaitOutput,
		Interoprabilitas:       data.Interoprabilitas,
		Tahun:                  data.Tahun,
		CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
		RadLevel1id: &web.DataDanInformasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(data.RadLevel1id),
		},
		RadLevel2id: &web.DataDanInformasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(data.RadLevel2id),
		},
		RadLevel3id: &web.DataDanInformasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(data.RadLevel3id),
		},
		RadLevel4id: &web.DataDanInformasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(data.RadLevel4id),
		},
		StrategicId: &web.DataDanInformasiPohonResponns{
			ID: nullInt32ToInt(data.StrategicId),
		},
		TacticalId: &web.DataDanInformasiPohonResponns{
			ID: nullInt32ToInt(data.TacticalId),
		},
		OperationalId: &web.DataDanInformasiPohonResponns{
			ID: nullInt32ToInt(data.OperationalId),
		},
	}
}

func ToAplikasiRespons(aplikasi domain.Aplikasi) web.AplikasiRespons {
	return web.AplikasiRespons{
		Id:                     aplikasi.Id,
		NamaAplikasi:           aplikasi.NamaAplikasi,
		FungsiAplikasi:         aplikasi.FungsiAplikasi,
		JenisAplikasi:          aplikasi.JenisAplikasi,
		ProdusenAplikasi:       aplikasi.ProdusenAplikasi,
		PjAplikasi:             aplikasi.PjAplikasi,
		KodeOPD:                aplikasi.KodeOPD,
		InformasiTerkaitInput:  aplikasi.InformasiTerkaitInput,
		InformasiTerkaitOutput: aplikasi.InformasiTerkaitOutput,
		Interoprabilitas:       aplikasi.Interoprabilitas,
		Tahun:                  aplikasi.Tahun,
		CreatedAt:              aplikasi.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:              aplikasi.UpdatedAt.Format("2006-01-02 15:04:05"),
		RaaLevel1id: &web.AplikasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(aplikasi.RaaLevel1id),
		},
		RaaLevel2id: &web.AplikasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(aplikasi.RaaLevel2id),
		},
		RaaLevel3id: &web.AplikasiReferensiArsitekturRespons{
			Id: nullInt32ToInt(aplikasi.RaaLevel3id),
		},
		StrategicId: &web.AplikasiPohonRespons{
			ID: nullInt32ToInt(aplikasi.StrategicId),
		},
		TacticalId: &web.AplikasiPohonRespons{
			ID: nullInt32ToInt(aplikasi.TacticalId),
		},
		OperationalId: &web.AplikasiPohonRespons{
			ID: nullInt32ToInt(aplikasi.OperationalId),
		},
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

func ToDomainSPBEResponses(domains []domain.DomainSPBE) []web.DomainSPBEResponse {
	var domainSPBEResponses []web.DomainSPBEResponse
	for _, domain := range domains {
		domainSPBEResponses = append(domainSPBEResponses, ToDomainSPBEResponse(domain))
	}
	return domainSPBEResponses
}

// func ToKebutuhanSPBERespons(kebutuhanSPBE domain.KebutuhanSPBE) web.KebutuhanSPBEResponse {
// 	var jenisKebutuhanResponses []web.JenisKebutuhanResponse
// 	for _, jk := range kebutuhanSPBE.JenisKebutuhan {
// 		var kondisiAwalResponses []web.KondisiAwalResponse
// 		for _, ka := range jk.KondisiAwal {
// 			kondisiAwalResponses = append(kondisiAwalResponses, web.KondisiAwalResponse{
// 				Id:               ka.Id,
// 				JenisKebutuhanId: ka.JenisKebutuhanId,
// 				Keterangan:       ka.Keterangan,
// 				Tahun:            ka.Tahun,
// 			})
// 		}
// 		jenisKebutuhanResponses = append(jenisKebutuhanResponses, web.JenisKebutuhanResponse{
// 			Id:          jk.Id,
// 			KebutuhanId: jk.KebutuhanId,
// 			Kebutuhan:   jk.Kebutuhan,
// 			KondisiAwal: kondisiAwalResponses,
// 		})
// 	}
// 	return web.KebutuhanSPBEResponse{
// 		ID:             kebutuhanSPBE.ID,
// 		KodeOpd:        kebutuhanSPBE.KodeOpd,
// 		Tahun:          kebutuhanSPBE.Tahun,
// 		NamaDomain:     kebutuhanSPBE.NamaDomain,
// 		IdProsesbisnis: kebutuhanSPBE.IdProsesbisnis,
// 		JenisKebutuhan: jenisKebutuhanResponses,
// 	}
// }

func ToKebutuhanSPBEResponse(kebutuhanSPBE domain.KebutuhanSPBE) web.KebutuhanSPBEResponse {
	return web.KebutuhanSPBEResponse{
		ID:         kebutuhanSPBE.ID,
		KodeOpd:    kebutuhanSPBE.KodeOpd,
		Tahun:      kebutuhanSPBE.Tahun,
		NamaDomain: kebutuhanSPBE.NamaDomain,
		ProsesBisnis: web.ProsesBisnisResponse{
			ID: kebutuhanSPBE.IdProsesbisnis,
		},
		JenisKebutuhan: ToJenisKebutuhanResponses(kebutuhanSPBE.JenisKebutuhan),
	}
}

func ToJenisKebutuhanResponses(jenisKebutuhan []domain.JenisKebutuhan) []web.JenisKebutuhanResponse {
	var jenisKebutuhanResponses []web.JenisKebutuhanResponse
	for _, jk := range jenisKebutuhan {
		jenisKebutuhanResponses = append(jenisKebutuhanResponses, ToJenisKebutuhanResponse(jk))
	}
	return jenisKebutuhanResponses
}

func ToJenisKebutuhanResponse(jenisKebutuhan domain.JenisKebutuhan) web.JenisKebutuhanResponse {
	return web.JenisKebutuhanResponse{
		Id:          jenisKebutuhan.Id,
		KebutuhanId: jenisKebutuhan.KebutuhanId,
		Kebutuhan:   jenisKebutuhan.Kebutuhan,
		KondisiAwal: ToKondisiAwalResponses(jenisKebutuhan.KondisiAwal),
	}
}

func ToKondisiAwalResponses(kondisiAwal []domain.KondisiAwal) []web.KondisiAwalResponse {
	var kondisiAwalResponses []web.KondisiAwalResponse
	for _, ka := range kondisiAwal {
		kondisiAwalResponses = append(kondisiAwalResponses, ToKondisiAwalResponse(ka))
	}
	return kondisiAwalResponses
}

func ToKondisiAwalResponse(kondisiAwal domain.KondisiAwal) web.KondisiAwalResponse {
	return web.KondisiAwalResponse{
		Id:               kondisiAwal.Id,
		JenisKebutuhanId: kondisiAwal.JenisKebutuhanId,
		Keterangan:       kondisiAwal.Keterangan,
		Tahun:            kondisiAwal.Tahun,
	}
}

func nullInt32ToInt(n sql.NullInt32) int {
	if n.Valid {
		return int(n.Int32)
	}
	return 0
}
