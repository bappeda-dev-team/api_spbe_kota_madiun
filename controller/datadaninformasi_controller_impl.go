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

type DataDanInformasiControllerImpl struct {
	datadaninformasiService service.DataDanInformasiService
}

func NewDataDanInformasiControllerImpl(datadaninformasiService service.DataDanInformasiService) *DataDanInformasiControllerImpl {
	return &DataDanInformasiControllerImpl{
		datadaninformasiService: datadaninformasiService,
	}
}

func (controller *DataDanInformasiControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	datainformasiResponse, err := controller.datadaninformasiService.FindByKodeOpd(request.Context(), kodeOPD, tahun)
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
		Status: "Success get data dan informasi",
		Data:   datainformasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DataDanInformasiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("dataId")
	id, err := strconv.Atoi(dataId)
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

	dataResponse, err := controller.datadaninformasiService.FindById(request.Context(), id, kodeOPD)
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
		Data:   dataResponse,
	})
}

func (controller *DataDanInformasiControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	DataDanInformasiCreateRequest := web.DataDanInformasiCreateRequest{}
	helper.ReadFromRequestBody(request, &DataDanInformasiCreateRequest)

	role := request.Context().Value("roles").(string)
	if role == "admin_opd" || role == "asn" {
		kodeOPD := request.Context().Value("kode_opd").(string)
		DataDanInformasiCreateRequest.KodeOPD = kodeOPD
	}

	dataResponse := controller.datadaninformasiService.Insert(request.Context(), DataDanInformasiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create data dan informasi",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DataDanInformasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := request.Context().Value("kode_opd").(string)
	role := request.Context().Value("roles").(string)

	DataDanInformasiUpdateRequest := web.DataDanInformasiUpdateRequest{}
	helper.ReadFromRequestBody(request, &DataDanInformasiUpdateRequest)

	layananspbeId := params.ByName("dataId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	DataDanInformasiUpdateRequest.Id = id

	if role == "asn" || role == "admin_opd" {
		existingData, err := controller.datadaninformasiService.FindById(request.Context(), id, kodeOPD)
		if err != nil || existingData.KodeOPD != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk memperbarui proses bisnis ini",
			})
			return
		}
		DataDanInformasiUpdateRequest.KodeOPD = kodeOPD
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

	dataResponse := controller.datadaninformasiService.Update(request.Context(), DataDanInformasiUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update data dan informasi",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DataDanInformasiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("dataId")
	id, err := strconv.Atoi(dataId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID data informasi tidak valid",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	kodeOPD, _ := request.Context().Value("kode_opd").(string)
	role, _ := request.Context().Value("roles").(string)

	err = controller.datadaninformasiService.Delete(request.Context(), id, kodeOPD, role)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   "Tidak dapat menghapus data dan informasi",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete data informasi",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DataDanInformasiControllerImpl) ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
	dataResponse, err := controller.datadaninformasiService.FindByKodeOpd(context.Background(), kodeOPD, tahun)
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
	f.SetCellValue(sheet, "B1", "Nama Data")
	f.MergeCell(sheet, "C1", "C2")
	f.SetCellValue(sheet, "C1", "Uraian Data")
	f.MergeCell(sheet, "D1", "D2")
	f.SetCellValue(sheet, "D1", "Sifat Data")
	f.MergeCell(sheet, "E1", "E2")
	f.SetCellValue(sheet, "E1", "Jenis Data")
	f.MergeCell(sheet, "F1", "F2")
	f.SetCellValue(sheet, "F1", "Validitas Data")
	f.MergeCell(sheet, "G1", "G2")
	f.SetCellValue(sheet, "G1", "Produsen Data")
	f.MergeCell(sheet, "H1", "H2")
	f.SetCellValue(sheet, "H1", "Penanggung Jawab Data")
	f.MergeCell(sheet, "I1", "I2")
	f.SetCellValue(sheet, "I1", "Informasi Terkait Input")
	f.MergeCell(sheet, "J1", "J2")
	f.SetCellValue(sheet, "J1", "Informasi Terkait Output")
	f.MergeCell(sheet, "K1", "K2")
	f.SetCellValue(sheet, "K1", "Interoperabilitas")
	f.MergeCell(sheet, "L1", "L2")
	f.SetCellValue(sheet, "L1", "Keterangan")
	f.MergeCell(sheet, "M1", "M2")
	f.SetCellValue(sheet, "M1", "Kode OPD")
	f.MergeCell(sheet, "N1", "N2")
	f.SetCellValue(sheet, "N1", "Tahun")
	f.MergeCell(sheet, "O1", "O2")
	f.SetCellValue(sheet, "O1", "RAD Level 1")
	f.MergeCell(sheet, "P1", "P2")
	f.SetCellValue(sheet, "P1", "RAD Level 2")
	f.MergeCell(sheet, "Q1", "Q2")
	f.SetCellValue(sheet, "Q1", "RAD Level 3")
	f.MergeCell(sheet, "R1", "R2")
	f.SetCellValue(sheet, "R1", "RAD Level 4")
	f.MergeCell(sheet, "S1", "S2")
	f.SetCellValue(sheet, "S1", "Strategic")
	f.MergeCell(sheet, "T1", "T2")
	f.SetCellValue(sheet, "T1", "Tactical")
	f.MergeCell(sheet, "U1", "U2")
	f.SetCellValue(sheet, "U1", "Operational")

	// Buat style header
	headerStyle, err := helper.CreateHeaderStyle(f)
	if err != nil {
		log.Printf("Error creating header style: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Terapkan style header ke sel-sel header
	headerCells := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1", "J1", "K1", "L1", "M1", "N1", "O1", "P1", "Q1", "R1", "S1", "T1"}
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
		{"Q1", "Q2"},
		{"R1", "R2"},
		{"S1", "S2"},
		{"T1", "T2"},
		{"U1", "U2"},
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
	for i, data := range dataResponse {
		row := i + 3
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), i+1)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), data.NamaData)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), data.UraianData)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), data.SifatData)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), data.JenisData)
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), data.ValiditasData)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), data.ProdusenData)
		f.SetCellValue(sheet, "H"+strconv.Itoa(row), data.PjData)
		f.SetCellValue(sheet, "I"+strconv.Itoa(row), data.InformasiTerkaitInput)
		f.SetCellValue(sheet, "J"+strconv.Itoa(row), data.InformasiTerkaitOutput)
		f.SetCellValue(sheet, "K"+strconv.Itoa(row), data.Interoprabilitas)
		f.SetCellValue(sheet, "L"+strconv.Itoa(row), data.Keterangan)
		f.SetCellValue(sheet, "M"+strconv.Itoa(row), data.KodeOPD)
		f.SetCellValue(sheet, "N"+strconv.Itoa(row), data.Tahun)
		if data.RadLevel1id != nil {
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), data.RadLevel1id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), "")
		}

		if data.RadLevel2id != nil {
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), data.RadLevel2id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), "")
		}

		if data.RadLevel3id != nil {
			f.SetCellValue(sheet, "Q"+strconv.Itoa(row), data.RadLevel3id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "Q"+strconv.Itoa(row), "")
		}

		if data.RadLevel4id != nil {
			f.SetCellValue(sheet, "R"+strconv.Itoa(row), data.RadLevel4id.Nama_referensi)
		} else {
			f.SetCellValue(sheet, "R"+strconv.Itoa(row), "")
		}

		if data.StrategicId != nil {
			f.SetCellValue(sheet, "S"+strconv.Itoa(row), data.StrategicId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "S"+strconv.Itoa(row), "")
		}

		if data.TacticalId != nil {
			f.SetCellValue(sheet, "T"+strconv.Itoa(row), data.TacticalId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "T"+strconv.Itoa(row), "")
		}

		if data.OperationalId != nil {
			f.SetCellValue(sheet, "U"+strconv.Itoa(row), data.OperationalId.NamaPohon)
		} else {
			f.SetCellValue(sheet, "U"+strconv.Itoa(row), "")
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
			"S" + strconv.Itoa(row),
			"T" + strconv.Itoa(row),
			"U" + strconv.Itoa(row),
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
		"N": 11,
		"O": 32,
		"P": 32,
		"Q": 32,
		"R": 32,
		"S": 87,
		"T": 87,
		"U": 101,
	}
	err = helper.SetColumnWidths(f, sheet, columnWidths)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Tentukan nama file berdasarkan role dan kode OPD
	var filename string
	if kodeOPD != "" {
		filename = "Domain_Data_dan_Informasi_" + kodeOPD + ".xlsx"
	} else {
		filename = "Domain_Data_dan_Informasi_All.xlsx"
	}

	// Gunakan helper untuk mengirim file Excel
	err = helper.SendExcelFile(writer, f, filename)
	if err != nil {
		// Log kesalahan telah ditangani oleh helper
		return
	}
}
