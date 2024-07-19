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

func NewDataDanInformasiController(datadaninformasiService service.DataDanInformasiService) DataDanInformasiController {
	return &DataDanInformasiControllerImpl{
		datadaninformasiService: datadaninformasiService,
	}
}

func (controller *DataDanInformasiControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := params.ByName("kodeOPD")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error

	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		helper.PanicIfError(err)
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
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid ID",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	dataResponse, err := controller.datadaninformasiService.FindById(request.Context(), id)
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
		Status: "Success get data by id",
		Data:   dataResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *DataDanInformasiControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	DataDanInformasiCreateRequest := web.DataDanInformasiCreateRequest{}
	helper.ReadFromRequestBody(request, &DataDanInformasiCreateRequest)

	dataResponse := controller.datadaninformasiService.Insert(request.Context(), DataDanInformasiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create data dan informasi",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DataDanInformasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	DataDanInformasiUpdateRequest := web.DataDanInformasiUpdateRequest{}
	helper.ReadFromRequestBody(request, &DataDanInformasiUpdateRequest)

	layananspbeId := params.ByName("dataId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	DataDanInformasiUpdateRequest.Id = id

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
	helper.PanicIfError(err)

	controller.datadaninformasiService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete data informasi",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
