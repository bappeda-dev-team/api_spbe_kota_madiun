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

type ProsesBisnisControllerImpl struct {
	ProsesBisnisService service.ProsesBisnisService
}

func NewProsesBisnisControllerImpl(prosbisService service.ProsesBisnisService) *ProsesBisnisControllerImpl {
	return &ProsesBisnisControllerImpl{
		ProsesBisnisService: prosbisService,
	}
}

func (controller *ProsesBisnisControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	prosesBisnisResponse, err := controller.ProsesBisnisService.GetProsesBisnis(request.Context(), kodeOPD, tahun)
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
		Status: "Berhasil mendapatkan proses bisnis",
		Data:   prosesBisnisResponse,
	})
}

func (controller *ProsesBisnisControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisId := params.ByName("prosesbisnisId")
	id, err := strconv.Atoi(prosesbisnisId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID tidak valid",
		})
		return
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := ""
	if role != "admin_kota" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	prosesbisnisResponse, err := controller.ProsesBisnisService.FindById(request.Context(), id, kodeOPD)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil mendapatkan proses bisnis berdasarkan ID",
		Data:   prosesbisnisResponse,
	})
}

func (controller *ProsesBisnisControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisInsertRequest := web.ProsesBisnisCreateRequest{}
	helper.ReadFromRequestBody(request, &prosesbisnisInsertRequest)

	// Ambil kode OPD dari context yang telah ditambahkan oleh middleware
	kodeOPD, ok := request.Context().Value("kode_opd").(string)
	if !ok {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   "Kode OPD tidak ditemukan",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Tambahkan kode OPD ke request
	prosesbisnisInsertRequest.KodeOPD = kodeOPD

	prosesbisnisResponse := controller.ProsesBisnisService.Insert(request.Context(), prosesbisnisInsertRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil membuat proses bisnis",
		Data:   prosesbisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProsesBisnisControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	if role != "asn" {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya pengguna ASN yang dapat memperbarui proses bisnis",
		})
		return
	}

	//read request body
	prosesbisnisUpdateRequest := web.ProsesBisnisUpdateRequest{}
	helper.ReadFromRequestBody(request, &prosesbisnisUpdateRequest)
	prosesbisnisId, _ := strconv.Atoi(params.ByName("prosesbisnisId"))
	prosesbisnisUpdateRequest.Id = prosesbisnisId

	// cek == kode opd
	existingProsesBisnis, err := controller.ProsesBisnisService.FindById(request.Context(), prosesbisnisId, kodeOPD)
	if err != nil || existingProsesBisnis.KodeOPD != kodeOPD {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Anda tidak memiliki akses untuk memperbarui proses bisnis ini",
		})
		return
	}

	prosesbisnisUpdateRequest.KodeOPD = kodeOPD

	prosesbisnisResponse := controller.ProsesBisnisService.Update(request.Context(), prosesbisnisUpdateRequest)

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil memperbarui proses bisnis",
		Data:   prosesbisnisResponse,
	})
}

func (controller *ProsesBisnisControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisId := params.ByName("prosesbisnisId")
	id, err := strconv.Atoi(prosesbisnisId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID proses bisnis tidak valid",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Ambil kode OPD dari context
	kodeOPD, ok := request.Context().Value("kode_opd").(string)
	if !ok {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   "Kode OPD tidak ditemukan",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err = controller.ProsesBisnisService.Delete(request.Context(), id, kodeOPD)
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

func (controller *ProsesBisnisControllerImpl) GetProsesBisnisGrouped(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	prosesBisnisResponse, err := controller.ProsesBisnisService.GetProsesBisnisGrouped(request.Context(), kodeOPD, tahun)
	if err != nil {
		log.Printf("Error mendapatkan Gap ESPBE: %v", err)
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Kesalahan Internal Server",
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan GAP ESPBE",
		Data:   prosesBisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProsesBisnisControllerImpl) GetProsesBisnisNoGap(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	prosesBisnisResponse, err := controller.ProsesBisnisService.GetProsesBisnisNoGap(request.Context(), kodeOPD, tahun)
	if err != nil {
		log.Printf("Error mendapatkan Proses Bisnis No Gap: %v", err)
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Kesalahan Internal Server",
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan Proses Bisnis No Gap",
		Data:   prosesBisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProsesBisnisControllerImpl) ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	// Ambil kode_opd dari query atau context (tergantung user role)
	role := request.Context().Value("roles").(string)
	kodeOPD := ""
	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	// Panggil service untuk mendapatkan data yang difilter
	prosesBisnisResponse, err := controller.ProsesBisnisService.GetProsesBisnis(context.Background(), kodeOPD, tahun)
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
	f.SetCellValue(sheet, "B1", "Nama Proses Bisnis")
	f.MergeCell(sheet, "C1", "C2")
	f.SetCellValue(sheet, "C1", "Kode Proses Bisnis")
	f.MergeCell(sheet, "D1", "D2")
	f.SetCellValue(sheet, "D1", "Kode OPD")
	f.MergeCell(sheet, "E1", "E2")
	f.SetCellValue(sheet, "E1", "Tahun")
	f.MergeCell(sheet, "F1", "F2")
	f.SetCellValue(sheet, "F1", "Bidang Urusan")
	f.MergeCell(sheet, "G1", "G2")
	f.SetCellValue(sheet, "G1", "RAB Level 1")
	f.MergeCell(sheet, "H1", "H2")
	f.SetCellValue(sheet, "H1", "RAB Level 2")
	f.MergeCell(sheet, "I1", "I2")
	f.SetCellValue(sheet, "I1", "RAB Level 3")
	f.MergeCell(sheet, "J1", "J2")
	f.SetCellValue(sheet, "J1", "Strategic")
	f.MergeCell(sheet, "K1", "K2")
	f.SetCellValue(sheet, "K1", "Tactical")
	f.MergeCell(sheet, "L1", "L2")
	f.SetCellValue(sheet, "L1", "Operational")

	// Buat style header
	headerStyle, err := helper.CreateHeaderStyle(f)
	if err != nil {
		log.Printf("Error creating header style: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Terapkan style header ke sel-sel header
	headerCells := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1", "J1", "K1", "L1"}
	err = helper.ApplyHeaderStyle(f, sheet, headerCells, headerStyle)
	if err != nil {
		log.Printf("Error applying header style: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Terapkan style header ke area gabungan
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
	for i, pb := range prosesBisnisResponse {
		row := i + 3
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), i+1)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), pb.NamaProsesBisnis)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), pb.KodeProsesBisnis)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), pb.KodeOPD)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), pb.Tahun)

		// Tambahkan bidang urusan jika ada
		if pb.BidangUrusan != nil {
			f.SetCellValue(sheet, "F"+strconv.Itoa(row), pb.BidangUrusan.BidangUrusan)
		} else {
			f.SetCellValue(sheet, "F"+strconv.Itoa(row), "")
		}

		// Tambahkan RAB Level 1, 2, 3, 4, 5, 6 jika ada
		if pb.RabLevel1 != nil {
			f.SetCellValue(sheet, "G"+strconv.Itoa(row), pb.RabLevel1.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "G"+strconv.Itoa(row), "")
		}

		if pb.RabLevel2 != nil {
			f.SetCellValue(sheet, "H"+strconv.Itoa(row), pb.RabLevel2.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "H"+strconv.Itoa(row), "")
		}

		if pb.RabLevel3 != nil {
			f.SetCellValue(sheet, "I"+strconv.Itoa(row), pb.RabLevel3.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "I"+strconv.Itoa(row), "")
		}

		if pb.RabLevel4 != nil {
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), pb.RabLevel4.NamaPohon)
		} else {
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), "")
		}

		if pb.RabLevel5 != nil {
			f.SetCellValue(sheet, "K"+strconv.Itoa(row), pb.RabLevel5.NamaPohon)
		} else {
			f.SetCellValue(sheet, "K"+strconv.Itoa(row), "")
		}

		if pb.RabLevel6 != nil {
			f.SetCellValue(sheet, "L"+strconv.Itoa(row), pb.RabLevel6.NamaPohon)
		} else {
			f.SetCellValue(sheet, "L"+strconv.Itoa(row), "")
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
		"B": 40,
		"C": 20,
		"D": 20,
		"E": 10,
		"F": 41,
		"G": 32,
		"H": 32,
		"I": 32,
		"J": 87,
		"K": 87,
		"L": 101,
	}
	err = helper.SetColumnWidths(f, sheet, columnWidths)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Tentukan nama file berdasarkan role dan kode OPD
	var filename string
	if kodeOPD != "" {
		filename = "Domain_Prosesbisnis_" + kodeOPD + ".xlsx"
	} else {
		filename = "Domain_Proses_Bisnis_All.xlsx"
	}

	// Gunakan helper untuk mengirim file Excel
	err = helper.SendExcelFile(writer, f, filename)
	if err != nil {
		// Log kesalahan telah ditangani oleh helper
		return
	}
}
