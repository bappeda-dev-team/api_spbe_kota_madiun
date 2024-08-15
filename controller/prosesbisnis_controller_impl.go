package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"log"

	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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
