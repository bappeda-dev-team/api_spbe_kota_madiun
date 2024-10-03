package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/xuri/excelize/v2"
)

type PetarencanaControllerImpl struct {
	PetarencanaService service.PetarencanaService
}

func NewPetarencanaControllerImpl(petarencanaService service.PetarencanaService) *PetarencanaControllerImpl {
	return &PetarencanaControllerImpl{PetarencanaService: petarencanaService}
}

func (controller *PetarencanaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunStr := request.URL.Query().Get("tahun")
	tahun := 0
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Format tahun tidak valid",
				Data:   nil,
			})
			return
		}
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	petarencanaResponse, err := controller.PetarencanaService.FindAll(request.Context(), kodeOPD, tahun)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil mendapatkan petarencana",
		Data:   petarencanaResponse,
	})

}

func (controller *PetarencanaControllerImpl) ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunStr := request.URL.Query().Get("tahun")
	tahun := 0
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Format tahun tidak valid",
				Data:   nil,
			})
			return
		}
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	petarencanaResponse, err := controller.PetarencanaService.FindAll(request.Context(), kodeOPD, tahun)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		})
		return
	}

	// Buat file Excel baru
	f := excelize.NewFile()

	// Buat sheet baru
	sheetName := "Petarencana"
	index, _ := f.NewSheet(sheetName)

	// Set judul kolom
	headers := []string{
		"No", "Kode OPD", "Tahun", "Nama Proses Bisnis", "Kode Proses Bisnis",
		"Layanan", "Data dan Informasi", "Aplikasi", "Keterangan Gap",
		"Nama Domain", "Indikator PJ", "Penanggung Jawab",
		"Jenis Kebutuhan", "Kondisi Awal",
		"Sasaran Kinerja Pegawai", "Anggaran Sasaran", "Pelaksana Sasaran",
		"Kode Sub Kegiatan", "Sub Kegiatan", "Tahun Pelaksanaan",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%s5", string(rune('A'+i)))
		f.SetCellValue(sheetName, cell, header)
	}

	// Isi data
	row := 6 // Mulai dari baris 6 (A6)
	for _, data := range petarencanaResponse {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), row-5)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), data.KodeOpd)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), data.Tahun)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), data.NamaProsesBisnis)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), data.KodeProsesBisnis)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), layananToString(data.Layanan))
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), dataDanInformasiToString(data.DataDanInformasi))
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), aplikasiToString(data.Aplikasi))

		keteranganGap := make([]string, 0)
		namaDomain := make([]string, 0)
		indikatorPj := make([]string, 0)
		penanggungJawab := make([]string, 0)
		jenisKebutuhan := make([]string, 0)
		kondisiAwal := make([]string, 0)
		sasaranKinerja := make([]string, 0)
		anggaranSasaran := make([]string, 0)
		pelaksanaSasaran := make([]string, 0)
		kodeSubKegiatan := make([]string, 0)
		subKegiatan := make([]string, 0)
		tahunPelaksanaan := make([]string, 0)

		for _, keterangan := range data.Keterangan {
			keteranganGap = append(keteranganGap, keterangan.KeteranganGap)
			namaDomain = append(namaDomain, keterangan.NamaDomain)
			indikatorPj = append(indikatorPj, keterangan.IndikatorPj)
			penanggungJawab = append(penanggungJawab, keterangan.PenanggungJawab.NamaOpd)
			for _, jk := range keterangan.JenisKebutuhan {
				jenisKebutuhan = append(jenisKebutuhan, jk.Kebutuhan)
				for _, ka := range jk.KondisiAwal {
					kondisiAwal = append(kondisiAwal, ka.Keterangan)
				}
			}
			for _, rencanaPelaksanaan := range keterangan.RencanaPelaksanaan {
				sasaranKinerja = append(sasaranKinerja, rencanaPelaksanaan.SasaranKinerja.SasaranKinerjaPegawai)
				anggaranSasaran = append(anggaranSasaran, rencanaPelaksanaan.SasaranKinerja.AnggaranSasaran)
				pelaksanaSasaran = append(pelaksanaSasaran, rencanaPelaksanaan.SasaranKinerja.PelaksanaSasaran)
				kodeSubKegiatan = append(kodeSubKegiatan, rencanaPelaksanaan.SasaranKinerja.KodeSubKegiatan)
				subKegiatan = append(subKegiatan, rencanaPelaksanaan.SasaranKinerja.SubKegiatan)
				if len(rencanaPelaksanaan.TahunPelaksanaan) > 0 {
					tahunPelaksanaan = append(tahunPelaksanaan, strconv.Itoa(rencanaPelaksanaan.TahunPelaksanaan[0].Tahun))
				}
			}
		}

		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), strings.Join(keteranganGap, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), strings.Join(namaDomain, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("K%d", row), strings.Join(indikatorPj, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), strings.Join(penanggungJawab, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), strings.Join(jenisKebutuhan, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), strings.Join(kondisiAwal, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("O%d", row), strings.Join(sasaranKinerja, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("P%d", row), strings.Join(anggaranSasaran, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("Q%d", row), strings.Join(pelaksanaSasaran, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("R%d", row), strings.Join(kodeSubKegiatan, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("S%d", row), strings.Join(subKegiatan, "\n"))
		f.SetCellValue(sheetName, fmt.Sprintf("T%d", row), strings.Join(tahunPelaksanaan, "\n"))
		row++
	}

	// Set lebar kolom
	for i := 0; i < len(headers); i++ {
		col := string(rune('A' + i))
		f.SetColWidth(sheetName, col, col, 20)
	}

	// Set sheet aktif
	f.SetActiveSheet(index)

	// Set header untuk download
	writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	writer.Header().Set("Content-Disposition", "attachment; filename=petarencana.xlsx")

	// Simpan file Excel
	if err := f.Write(writer); err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal mengekspor Excel",
			Data:   nil,
		})
		return
	}
}

// Fungsi helper untuk mengonversi sslice ke string
func layananToString(layanan []web.RencanaLayanan) string {
	var result []string
	for _, l := range layanan {
		if l.NamaLayanan.Valid {
			result = append(result, l.NamaLayanan.String)
		}
	}
	return strings.Join(result, ", ")
}

func dataDanInformasiToString(data []web.RencanaDataDanInformasi) string {
	var result []string
	for _, d := range data {
		if d.NamaData.Valid {
			result = append(result, d.NamaData.String)
		}
	}
	return strings.Join(result, ", ")
}

func aplikasiToString(aplikasi []web.RencanaAplikasi) string {
	var result []string
	for _, a := range aplikasi {
		if a.NamaAplikasi.Valid {
			result = append(result, a.NamaAplikasi.String)
		}
	}
	return strings.Join(result, ", ")
}
