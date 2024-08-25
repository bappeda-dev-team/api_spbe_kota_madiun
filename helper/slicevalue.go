package helper

import "api_spbe_kota_madiun/model/web"

func ContainsLayanan(layanans []web.GapLayanan, layanan web.GapLayanan) bool {
	for _, l := range layanans {
		if l.NamaLayanan.Valid == layanan.NamaLayanan.Valid &&
			(l.NamaLayanan.Valid && l.NamaLayanan.String == layanan.NamaLayanan.String || !l.NamaLayanan.Valid) {
			return true
		}
	}
	return false
}

func ContainData(data []web.GapDataDanInformasi, datainfor web.GapDataDanInformasi) bool {
	for _, i := range data {
		if i.NamaData.Valid == datainfor.NamaData.Valid &&
			(i.NamaData.Valid && i.NamaData.String == datainfor.NamaData.String || !i.NamaData.Valid) {
			return true
		}
	}
	return false
}

func ContainAplikasi(aplikasi []web.GapAplikasi, app web.GapAplikasi) bool {
	for _, a := range aplikasi {
		if a.NamaAplikasi.Valid == app.NamaAplikasi.Valid &&
			(a.NamaAplikasi.Valid && a.NamaAplikasi.String == app.NamaAplikasi.String || !a.NamaAplikasi.Valid) {
			return true
		}
	}
	return false
}

func ContainKeterangan(keterangan []web.GapKeterangan, ket web.GapKeterangan) bool {
	for _, k := range keterangan {
		if k.Keterangan.Valid == ket.Keterangan.Valid &&
			(k.Keterangan.Valid && k.Keterangan.String == ket.Keterangan.String || !k.Keterangan.Valid) {
			return true
		}
	}
	return false
}
