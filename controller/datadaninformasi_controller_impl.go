package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

	DataDanInformasiCreateRequest.KodeOPD = kodeOPD

	dataResponse := controller.datadaninformasiService.Insert(request.Context(), DataDanInformasiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create data dan informasi",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DataDanInformasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	if role != "asn" {
		webResponse := web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya pengguna ASN yang dapat memperbarui data dan informasi",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	DataDanInformasiUpdateRequest := web.DataDanInformasiUpdateRequest{}
	helper.ReadFromRequestBody(request, &DataDanInformasiUpdateRequest)

	layananspbeId := params.ByName("dataId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	DataDanInformasiUpdateRequest.Id = id

	existingData, err := controller.datadaninformasiService.FindById(request.Context(), id, kodeOPD)
	if err != nil || existingData.KodeOPD != kodeOPD {
		webResponse := web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Anda tidak memiliki akses untuk memperbarui data ini",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	DataDanInformasiUpdateRequest.KodeOPD = kodeOPD

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

	err = controller.datadaninformasiService.Delete(request.Context(), id, kodeOPD)
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
