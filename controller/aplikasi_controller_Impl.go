package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/xuri/excelize/v2"
)

type AplikasiControllerImpl struct {
	aplikasiService service.AplikasiService
}

func NewAplikasiControllerImpl(aplikasiService service.AplikasiService) *AplikasiControllerImpl {
	return &AplikasiControllerImpl{
		aplikasiService: aplikasiService,
	}
}

func (controller *AplikasiControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
	aplikasiResponse, err := controller.aplikasiService.FindByKodeOpd(request.Context(), kodeOPD, tahun)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success get data aplikasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aplikasiId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(aplikasiId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid ID",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := ""
	if role != "admin_kota" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	aplikasiResponse, err := controller.aplikasiService.FindById(request.Context(), id, kodeOPD)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success get aplikasi by id",
		Data:   aplikasiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *AplikasiControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aplikasiCreateRequest := web.AplikasiCreateRequest{}
	helper.ReadFromRequestBody(request, &aplikasiCreateRequest)

	role := request.Context().Value("roles").(string)
	if role == "admin_opd" || role == "asn" {
		kodeOPD := request.Context().Value("kode_opd").(string)
		aplikasiCreateRequest.KodeOPD = kodeOPD
	}

	aplikasiResponse := controller.aplikasiService.Insert(request.Context(), aplikasiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create data aplikasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	AplikasiUpdateRequest := web.AplikasiUpdateRequest{}
	helper.ReadFromRequestBody(request, &AplikasiUpdateRequest)

	aplikasiId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(aplikasiId)
	helper.PanicIfError(err)

	AplikasiUpdateRequest.Id = id

	if role == "asn" || role == "admin_opd" {
		existingAplikasi, err := controller.aplikasiService.FindById(request.Context(), id, kodeOPD)
		if err != nil || existingAplikasi.KodeOPD != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk memperbarui aplikasi ini",
			})
			return
		}
		AplikasiUpdateRequest.KodeOPD = kodeOPD
	} else if role == "admin_kota" {
		// untuk admin_kota, tidak ada cek kode OPD
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Role tidak diizinkan untuk memperbarui proses bisnis",
		})
		return
	}

	aplikasiResponse := controller.aplikasiService.Update(request.Context(), AplikasiUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update data aplikasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *AplikasiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(dataId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID proses bisnis tidak valid",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	kodeOPD, _ := request.Context().Value("kode_opd").(string)
	role, _ := request.Context().Value("roles").(string)

	err = controller.aplikasiService.Delete(request.Context(), id, kodeOPD, role)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil menghapus proses bisnis",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil query parameter 'tahun'
	tahunStr := request.URL.Query().Get("tahun")
	tahun := 0
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			http.Error(writer, "Format tahun tidak valid", http.StatusBadRequest)
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

	// Panggil service untuk mendapatkan data yang difilter
	aplikasiResponse, err := controller.aplikasiService.FindByKodeOpd(context.Background(), kodeOPD, tahun)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Buat file Excel dan isi dengan data
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.MergeCell(sheet, "A1", "A2")
	f.SetCellValue(sheet, "A1", "No.")
	f.MergeCell(sheet, "B1", "B2")
	f.SetCellValue(sheet, "B1", "Nama Aplikasi")
	f.MergeCell(sheet, "C1", "C2")
	f.SetCellValue(sheet, "C1", "Fungsi Aplikasi")
	f.MergeCell(sheet, "D1", "D2")
	f.SetCellValue(sheet, "D1", "Jenis Aplikasi")
	f.MergeCell(sheet, "E1", "E2")
	f.SetCellValue(sheet, "E1", "Produsen Aplikasi")
	f.MergeCell(sheet, "F1", "F2")
	f.SetCellValue(sheet, "F1", "Penanggung Jawab Data")
	f.MergeCell(sheet, "G1", "G2")
	f.SetCellValue(sheet, "G1", "Informasi Terkait Input")
	f.MergeCell(sheet, "H1", "H2")
	f.SetCellValue(sheet, "H1", "Informasi Terkait Output")
	f.MergeCell(sheet, "I1", "I2")
	f.SetCellValue(sheet, "I1", "Interoperabilitas")
	f.MergeCell(sheet, "J1", "J2")
	f.SetCellValue(sheet, "J1", "Keterangan")
	f.MergeCell(sheet, "K1", "K2")
	f.SetCellValue(sheet, "K1", "Tahun")
	f.MergeCell(sheet, "L1", "L2")
	f.SetCellValue(sheet, "L1", "Kode OPD")
	f.MergeCell(sheet, "M1", "M2")
	f.SetCellValue(sheet, "M1", "RAA Level 1")
	f.MergeCell(sheet, "N1", "N2")
	f.SetCellValue(sheet, "N1", "RAA Level 2")
	f.MergeCell(sheet, "O1", "O2")
	f.SetCellValue(sheet, "O1", "RAA Level 3")
	f.MergeCell(sheet, "P1", "P2")
	f.SetCellValue(sheet, "P1", "Strategic")
	f.MergeCell(sheet, "Q1", "Q2")
	f.SetCellValue(sheet, "Q1", "Tactical")
	f.MergeCell(sheet, "R1", "R2")
	f.SetCellValue(sheet, "R1", "Operational")

	headerStyle, err := helper.CreateHeaderStyle(f)
	if err != nil {
		log.Printf("Error creating header style: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	headerCells := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1", "J1", "K1", "L1", "M1", "N1", "O1", "P1", "Q1", "R1"}
	err = helper.ApplyHeaderStyle(f, sheet, headerCells, headerStyle)
	if err != nil {
		log.Printf("Error applying header style: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	mergedCells := []struct {
		startCell string
		endCell   string
	}{
		{"A1", "A2"},
		{"B1", "B2"},
		{"C1", "C2"},
		{"D1", "D2"},
		{"E1", "E2"},
		{"F1", "F2"},
		{"G1", "G2"},
		{"H1", "H2"},
		{"I1", "I2"},
		{"J1", "J2"},
		{"K1", "K2"},
		{"L1", "L2"},
		{"M1", "M2"},
		{"N1", "N2"},
		{"O1", "O2"},
		{"P1", "P2"},
		{"Q1", "Q2"},
		{"R1", "R2"},
	}

	for _, rangeCells := range mergedCells {
		err = f.SetCellStyle(sheet, rangeCells.startCell, rangeCells.endCell, headerStyle)
		if err != nil {
			log.Printf("Error applying header style: %v", err)
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Buat style untuk body
	bodyStyle, err := helper.CreateBodyStyle(f)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Tambahkan data ke Excel dengan style body
	for i, aplikasi := range aplikasiResponse {
		row := i + 3
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), i+1)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), aplikasi.NamaAplikasi)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), aplikasi.FungsiAplikasi)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), aplikasi.JenisAplikasi)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), aplikasi.ProdusenAplikasi)
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), aplikasi.PjAplikasi)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), aplikasi.InformasiTerkaitInput)
		f.SetCellValue(sheet, "H"+strconv.Itoa(row), aplikasi.InformasiTerkaitOutput)
		f.SetCellValue(sheet, "I"+strconv.Itoa(row), aplikasi.Interoprabilitas)
		if aplikasi.Keterangan != nil {
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), aplikasi.Keterangan)
		} else {
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), "")
		}
		f.SetCellValue(sheet, "K"+strconv.Itoa(row), aplikasi.Tahun)
		f.SetCellValue(sheet, "L"+strconv.Itoa(row), aplikasi.KodeOPD)
		if aplikasi.RaaLevel1id != nil {
			f.SetCellValue(sheet, "M"+strconv.Itoa(row), aplikasi.RaaLevel1id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "M"+strconv.Itoa(row), "")
		}

		if aplikasi.RaaLevel2id != nil {
			f.SetCellValue(sheet, "N"+strconv.Itoa(row), aplikasi.RaaLevel2id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "N"+strconv.Itoa(row), "")
		}

		if aplikasi.RaaLevel3id != nil {
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), aplikasi.RaaLevel3id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), "")
		}

		if aplikasi.StrategicId != nil {
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), aplikasi.StrategicId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), "")
		}

		if aplikasi.TacticalId != nil {
			f.SetCellValue(sheet, "Q"+strconv.Itoa(row), aplikasi.TacticalId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "Q"+strconv.Itoa(row), "")
		}

		if aplikasi.OperationalId != nil {
			f.SetCellValue(sheet, "R"+strconv.Itoa(row), aplikasi.OperationalId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "R"+strconv.Itoa(row), "")
		}

		// Terapkan style body pada baris data
		bodyCells := []string{
			"A" + strconv.Itoa(row),
			"B" + strconv.Itoa(row),
			"C" + strconv.Itoa(row),
			"D" + strconv.Itoa(row),
			"E" + strconv.Itoa(row),
			"F" + strconv.Itoa(row),
			"G" + strconv.Itoa(row),
			"H" + strconv.Itoa(row),
			"I" + strconv.Itoa(row),
			"J" + strconv.Itoa(row),
			"K" + strconv.Itoa(row),
			"L" + strconv.Itoa(row),
			"M" + strconv.Itoa(row),
			"N" + strconv.Itoa(row),
			"O" + strconv.Itoa(row),
			"P" + strconv.Itoa(row),
			"Q" + strconv.Itoa(row),
			"R" + strconv.Itoa(row),
		}
		err = helper.ApplyBodyStyle(f, sheet, bodyCells, bodyStyle)
		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Mengatur lebar kolom
	columnWidths := map[string]float64{
		"A": 10,
		"B": 32,
		"C": 32,
		"D": 32,
		"E": 32,
		"F": 32,
		"G": 32,
		"H": 32,
		"I": 32,
		"J": 32,
		"K": 32,
		"L": 32,
		"M": 32,
		"N": 32,
		"O": 32,
		"P": 87,
		"Q": 87,
		"R": 101,
	}
	err = helper.SetColumnWidths(f, sheet, columnWidths)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Tentukan nama file berdasarkan role dan kode OPD
	var filename string
	if kodeOPD != "" {
		filename = "Aplikasi_" + kodeOPD + ".xlsx"
	} else {
		filename = "Aplikasi_All.xlsx"
	}

	// Gunakan helper untuk mengirim file Excel
	err = helper.SendExcelFile(writer, f, filename)
	if err != nil {
		// Log kesalahan telah ditangani oleh helper
		return
	}
}
