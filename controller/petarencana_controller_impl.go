package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"fmt"
	"log"
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

// Tambahkan fungsi helper untuk memformat angka menjadi format Rupiah
func formatRupiah(amount float64) string {
	return fmt.Sprintf("Rp %.2f", amount)
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
		"No", "PROSES BISNIS / POHON KINERJA", "", "Jenis Pelayanan", "Data dan Informasi", "Aplikasi",
		"KEBUTUHAN SPBE", "", "", "", "", "", "", "",
		"RENCANA PELAKSANAAN", "", "", "", "", "", "", "", "", "", "", "",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%s5", string(rune('A'+i)))
		f.SetCellValue(sheetName, cell, header)
	}

	// Gabungkan sel untuk judul "PROSES BISNIS / POHON KINERJA"
	f.MergeCell(sheetName, "B5", "C5")

	// Set sub-judul untuk Proses Bisnis
	f.SetCellValue(sheetName, "B6", "Kode Proses Bisnis")
	f.SetCellValue(sheetName, "C6", "Nama Proses Bisnis")

	// Gabungkan sel untuk sub-judul Proses Bisnis
	f.MergeCell(sheetName, "B6", "B7")
	f.MergeCell(sheetName, "C6", "C7")

	// Gabungkan sel untuk Jenis Pelayanan, Data dan Informasi, dan Aplikasi
	f.MergeCell(sheetName, "D5", "D7")
	f.MergeCell(sheetName, "E5", "E7")
	f.MergeCell(sheetName, "F5", "F7")

	// Gabungkan sel untuk KEBUTUHAN SPBE
	f.MergeCell(sheetName, "G5", "N5")

	// Set sub-judul untuk KEBUTUHAN SPBE
	f.SetCellValue(sheetName, "G6", "Keterangan Gap")
	f.SetCellValue(sheetName, "H6", "Nama Domain")
	f.SetCellValue(sheetName, "I6", "Perangkat Daerah")
	f.MergeCell(sheetName, "I6", "J6")
	f.SetCellValue(sheetName, "K6", "Jenis Kebutuhan")
	f.SetCellValue(sheetName, "L6", "Kondisi Awal")

	// Merge Kondisi Awal
	f.MergeCell(sheetName, "L6", "N6")

	// Set tahun untuk Kondisi Awal
	f.SetCellValue(sheetName, "L7", "2022")
	f.SetCellValue(sheetName, "M7", "2023")
	f.SetCellValue(sheetName, "N7", "2024")

	// Gabungkan sel untuk RENCANA PELAKSANAAN
	f.MergeCell(sheetName, "O5", "Z5")

	// Set sub-judul untuk RENCANA PELAKSANAAN
	f.SetCellValue(sheetName, "O6", "Sasaran Kinerja")
	f.MergeCell(sheetName, "O6", "P6")
	f.SetCellValue(sheetName, "O7", "Kode Sasaran")
	f.SetCellValue(sheetName, "P7", "Sasaran Kinerja")
	f.SetCellValue(sheetName, "Q6", "Anggaran Sasaran")
	f.SetCellValue(sheetName, "R6", "Pelaksana Sasaran")
	f.SetCellValue(sheetName, "S6", "Kode Sub Kegiatan")
	f.SetCellValue(sheetName, "T6", "Sub Kegiatan")
	f.SetCellValue(sheetName, "U6", "Tahun Pelaksanaan")
	f.MergeCell(sheetName, "U6", "Z6")
	f.SetCellValue(sheetName, "U7", "2025")
	f.SetCellValue(sheetName, "V7", "2026")
	f.SetCellValue(sheetName, "W7", "2027")
	f.SetCellValue(sheetName, "X7", "2028")
	f.SetCellValue(sheetName, "Y7", "2029")

	// Gabungkan sel untuk kolom RENCANA PELAKSANAAN dari baris 6 ke 7
	f.MergeCell(sheetName, "Q6", "Q7")
	f.MergeCell(sheetName, "R6", "R7")
	f.MergeCell(sheetName, "S6", "S7")
	f.MergeCell(sheetName, "T6", "T7")

	// Gabungkan sel untuk kolom KEBUTUHAN SPBE dari baris 6 ke 7
	f.MergeCell(sheetName, "G6", "G7")
	f.MergeCell(sheetName, "H6", "H7")
	f.MergeCell(sheetName, "I6", "J7")
	f.MergeCell(sheetName, "K6", "K7")

	// Gabungkan sel untuk No
	f.MergeCell(sheetName, "A5", "A7")

	// Isi data
	row := 8 // Mulai dari baris 8 (A8)
	for _, data := range petarencanaResponse {
		startRow := row
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), row-7)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), data.KodeProsesBisnis)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), data.NamaProsesBisnis)

		// Gabungkan sel untuk Kode Proses Bisnis dan Nama Proses Bisnis
		f.MergeCell(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row+1))
		f.MergeCell(sheetName, fmt.Sprintf("C%d", row), fmt.Sprintf("C%d", row+1))

		// Jenis Pelayanan
		layanan := layananToString(data.Layanan)
		layananItems := strings.Split(layanan, ", ")
		for i, item := range layananItems {
			f.SetCellValue(sheetName, fmt.Sprintf("D%d", row+i), item)
		}

		// Data dan Informasi
		dataDanInformasi := dataDanInformasiToString(data.DataDanInformasi)
		dataDanInformasiItems := strings.Split(dataDanInformasi, ", ")
		for i, item := range dataDanInformasiItems {
			f.SetCellValue(sheetName, fmt.Sprintf("E%d", row+i), item)
		}

		// Aplikasi
		aplikasi := aplikasiToString(data.Aplikasi)
		aplikasiItems := strings.Split(aplikasi, ", ")
		for i, item := range aplikasiItems {
			f.SetCellValue(sheetName, fmt.Sprintf("F%d", row+i), item)
		}

		maxRows := max(len(layananItems), len(dataDanInformasiItems), len(aplikasiItems), 1)

		// KEBUTUHAN SPBE
		for i, keterangan := range data.Keterangan {
			currentRow := row + i
			f.SetCellValue(sheetName, fmt.Sprintf("G%d", currentRow), keterangan.KeteranganGap)
			f.SetCellValue(sheetName, fmt.Sprintf("H%d", currentRow), keterangan.NamaDomain)
			f.SetCellValue(sheetName, fmt.Sprintf("I%d", currentRow), keterangan.IndikatorPj)
			f.SetCellValue(sheetName, fmt.Sprintf("J%d", currentRow), keterangan.PenanggungJawab.NamaOpd)

			for j, jk := range keterangan.JenisKebutuhan {
				currentRowJK := currentRow + j
				f.SetCellValue(sheetName, fmt.Sprintf("K%d", currentRowJK), jk.Kebutuhan)

				for _, ka := range jk.KondisiAwal {
					switch ka.Tahun {
					case 2022:
						f.SetCellValue(sheetName, fmt.Sprintf("L%d", currentRowJK), ka.Keterangan)
					case 2023:
						f.SetCellValue(sheetName, fmt.Sprintf("M%d", currentRowJK), ka.Keterangan)
					case 2024:
						f.SetCellValue(sheetName, fmt.Sprintf("N%d", currentRowJK), ka.Keterangan)
					}
				}
			}

			for k, rencanaPelaksanaan := range keterangan.RencanaPelaksanaan {
				currentRowRP := currentRow + k
				f.SetCellValue(sheetName, fmt.Sprintf("O%d", currentRowRP), rencanaPelaksanaan.SasaranKinerja.KodeSasaran)
				f.SetCellValue(sheetName, fmt.Sprintf("P%d", currentRowRP), rencanaPelaksanaan.SasaranKinerja.SasaranKinerjaPegawai)

				// Konversi AnggaranSasaran dari string ke float64
				anggaranFloat, err := strconv.ParseFloat(rencanaPelaksanaan.SasaranKinerja.AnggaranSasaran, 64)
				if err != nil {
					// Tangani error jika konversi gagal
					log.Printf("Error mengonversi anggaran: %v", err)
					anggaranFloat = 0 // Atau nilai default lainnya
				}

				// Format anggaran menjadi Rupiah
				anggaranRupiah := formatRupiah(anggaranFloat)
				f.SetCellValue(sheetName, fmt.Sprintf("Q%d", currentRowRP), anggaranRupiah)

				f.SetCellValue(sheetName, fmt.Sprintf("R%d", currentRowRP), rencanaPelaksanaan.SasaranKinerja.PelaksanaSasaran)
				f.SetCellValue(sheetName, fmt.Sprintf("S%d", currentRowRP), rencanaPelaksanaan.SasaranKinerja.KodeSubKegiatan)
				f.SetCellValue(sheetName, fmt.Sprintf("T%d", currentRowRP), rencanaPelaksanaan.SasaranKinerja.SubKegiatan)

				// Isi tahun pelaksanaan dengan simbol centang
				for _, tahun := range rencanaPelaksanaan.TahunPelaksanaan {
					col := ""
					switch tahun.Tahun {
					case 2025:
						col = "U"
					case 2026:
						col = "V"
					case 2027:
						col = "W"
					case 2028:
						col = "X"
					case 2029:
						col = "Y"
					}
					if col != "" {
						f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, currentRowRP), "✓")
					}
				}
			}

			maxRows = max(maxRows, len(keterangan.JenisKebutuhan), len(keterangan.RencanaPelaksanaan))
		}

		// Gabungkan sel untuk kolom yang tidak berubah
		if maxRows > 1 {
			f.MergeCell(sheetName, fmt.Sprintf("A%d", startRow), fmt.Sprintf("A%d", startRow+maxRows-1))
			f.MergeCell(sheetName, fmt.Sprintf("B%d", startRow), fmt.Sprintf("B%d", startRow+maxRows-1))
			f.MergeCell(sheetName, fmt.Sprintf("C%d", startRow), fmt.Sprintf("C%d", startRow+maxRows-1))
		}

		row += maxRows
	}

	// Set lebar kolom
	for i := 0; i < len(headers); i++ {
		col := string(rune('A' + i))
		f.SetColWidth(sheetName, col, col, 20)
	}

	// Set style untuk header
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetCellStyle(sheetName, "A5", "Z7", headerStyle)

	// Setelah mengisi semua data, atur format sel untuk kolom anggaran
	anggaranStyle, _ := f.NewStyle(&excelize.Style{
		NumFmt: 163, // Format kustom untuk Rupiah
	})
	f.SetColStyle(sheetName, "Q", anggaranStyle)

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

// Fungsi helper untuk mendapatkan nilai maksimum
func max(numbers ...int) int {
	maxNum := numbers[0]
	for _, num := range numbers {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
