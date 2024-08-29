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

type LayananSpbeControllerImpl struct {
	layananSpbeService service.LayananSpbeService
}

func NewLayananSPBEControllerImpl(layananspbeService service.LayananSpbeService) *LayananSpbeControllerImpl {
	return &LayananSpbeControllerImpl{
		layananSpbeService: layananspbeService,
	}
}

func (controller *LayananSpbeControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	layananSpbeResponse, err := controller.layananSpbeService.FindByKodeOpd(request.Context(), kodeOPD, tahun)
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
		Status: "Success get layanan spbe",
		Data:   layananSpbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
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

	layananspbeResponse, err := controller.layananSpbeService.FindById(request.Context(), id, kodeOPD)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Mengirimkan respons dengan data proses bisnis
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success get layanan spbe by id",
		Data:   layananspbeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	LayananSpbeCreateRequest := web.LayananSpbeCreateRequest{}
	helper.ReadFromRequestBody(request, &LayananSpbeCreateRequest)

	role := request.Context().Value("roles").(string)
	if role == "admin_opd" || role == "asn" {
		kodeOPD := request.Context().Value("kode_opd").(string)
		LayananSpbeCreateRequest.KodeOPD = kodeOPD
	}

	layananspbeResponse := controller.layananSpbeService.Insert(request.Context(), LayananSpbeCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create layanan spbe",
		Data:   layananspbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *LayananSpbeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	LayananSpbeUpdateRequest := web.LayananSpbeUpdateRequest{}
	helper.ReadFromRequestBody(request, &LayananSpbeUpdateRequest)
	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)
	LayananSpbeUpdateRequest.Id = id

	if role == "asn" || role == "admin_opd" {
		existingLayananSpbe, err := controller.layananSpbeService.FindById(request.Context(), id, kodeOPD)
		if err != nil || existingLayananSpbe.KodeOPD != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk memperbarui proses bisnis ini",
			})
			return
		}
		LayananSpbeUpdateRequest.KodeOPD = kodeOPD
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

	layananspbeResponse := controller.layananSpbeService.Update(request.Context(), LayananSpbeUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update layanan spbe",
		Data:   layananspbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	kodeOPD, _ := request.Context().Value("kode_opd").(string)
	role, _ := request.Context().Value("roles").(string)

	err = controller.layananSpbeService.Delete(request.Context(), id, kodeOPD, role)
	if err != nil {
		if err.Error() == "layanan spbe tidak ditemukan untuk OPD ini" {
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk menghapus layanan spbe ini",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		panic(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete layanan spbe",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
	layananspbeResponse, err := controller.layananSpbeService.FindByKodeOpd(context.Background(), kodeOPD, tahun)
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
	f.SetCellValue(sheet, "B1", "Nama Layanan")
	f.MergeCell(sheet, "C1", "C2")
	f.SetCellValue(sheet, "C1", "Kode Layanan")
	f.MergeCell(sheet, "D1", "D2")
	f.SetCellValue(sheet, "D1", "Kode OPD")
	f.MergeCell(sheet, "E1", "E2")
	f.SetCellValue(sheet, "E1", "Tahun")
	f.MergeCell(sheet, "F1", "F2")
	f.SetCellValue(sheet, "F1", "Tujuan Layanan")
	f.MergeCell(sheet, "G1", "G2")
	f.SetCellValue(sheet, "G1", "Fungsi Layanan")
	f.MergeCell(sheet, "H1", "H2")
	f.SetCellValue(sheet, "H1", "Kementrian Terkait")
	f.MergeCell(sheet, "I1", "I2")
	f.SetCellValue(sheet, "I1", "Metode Layaan")
	f.MergeCell(sheet, "J1", "J2")
	f.SetCellValue(sheet, "J1", "RAL Level 1")
	f.MergeCell(sheet, "K1", "K2")
	f.SetCellValue(sheet, "K1", "RAL Level 2")
	f.MergeCell(sheet, "L1", "L2")
	f.SetCellValue(sheet, "L1", "RAL Level 3")
	f.MergeCell(sheet, "M1", "M2")
	f.SetCellValue(sheet, "M1", "RAL Level 4")
	f.MergeCell(sheet, "N1", "N2")
	f.SetCellValue(sheet, "N1", "Strategic")
	f.MergeCell(sheet, "O1", "O2")
	f.SetCellValue(sheet, "O1", "Tactical")
	f.MergeCell(sheet, "P1", "P2")
	f.SetCellValue(sheet, "P1", "Operational")

	// Buat style header
	headerStyle, err := helper.CreateHeaderStyle(f)
	if err != nil {
		log.Printf("Error creating header style: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Terapkan style header ke sel-sel header
	headerCells := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1", "J1", "K1", "L1", "M1", "N1", "O1", "P1"}
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
		{"M1", "M2"},
		{"N1", "N2"},
		{"O1", "O2"},
		{"P1", "P2"},
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
	for i, l := range layananspbeResponse {
		row := i + 3
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), i+1)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), l.NamaLayanan)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), l.KodeLayanan)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), l.KodeOPD)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), l.Tahun)

		// Tambahkan bidang urusan jika ada
		if l.TujuanLayananId != nil {
			f.SetCellValue(sheet, "F"+strconv.Itoa(row), l.TujuanLayananId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "F"+strconv.Itoa(row), "")
		}

		f.SetCellValue(sheet, "G"+strconv.Itoa(row), l.FungsiLayanan)
		f.SetCellValue(sheet, "H"+strconv.Itoa(row), l.KementrianTerkait)
		f.SetCellValue(sheet, "I"+strconv.Itoa(row), l.MetodeLayanan)

		if l.RalLevel1id != nil {
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), l.RalLevel1id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), "")
		}

		if l.RalLevel2id != nil {
			f.SetCellValue(sheet, "K"+strconv.Itoa(row), l.RalLevel2id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "K"+strconv.Itoa(row), "")
		}

		if l.RalLevel3id != nil {
			f.SetCellValue(sheet, "L"+strconv.Itoa(row), l.RalLevel3id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "L"+strconv.Itoa(row), "")
		}

		if l.RalLevel4id != nil {
			f.SetCellValue(sheet, "M"+strconv.Itoa(row), l.RalLevel4id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "M"+strconv.Itoa(row), "")
		}

		if l.StrategicId != nil {
			f.SetCellValue(sheet, "N"+strconv.Itoa(row), l.StrategicId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "N"+strconv.Itoa(row), "")
		}

		if l.TacticalId != nil {
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), l.TacticalId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), "")
		}

		if l.OperationalId != nil {
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), l.OperationalId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), "")
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
		"F": 50,
		"G": 32,
		"H": 32,
		"I": 32,
		"J": 32,
		"K": 32,
		"L": 32,
		"M": 32,
		"N": 87,
		"O": 87,
		"P": 101,
	}
	err = helper.SetColumnWidths(f, sheet, columnWidths)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Tentukan nama file berdasarkan role dan kode OPD
	var filename string
	if kodeOPD != "" {
		filename = "Domain_LayananSpbe_" + kodeOPD + ".xlsx"
	} else {
		filename = "Domain_LayananSPBE_All.xlsx"
	}

	// Gunakan helper untuk mengirim file Excel
	err = helper.SendExcelFile(writer, f, filename)
	if err != nil {
		// Log kesalahan telah ditangani oleh helper
		return
	}
}
