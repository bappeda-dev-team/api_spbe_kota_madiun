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

func NewProsesBisnisController(prosbisService service.ProsesBisnisService) ProsesBisnisController {
	return &ProsesBisnisControllerImpl{
		ProsesBisnisService: prosbisService,
	}
}

func (controller *ProsesBisnisControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := params.ByName("kodeOPD")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error

	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		helper.PanicIfError(err)
	}

	prosesBisnisResponse, err := controller.ProsesBisnisService.GetProsesBisnis(request.Context(), kodeOPD, tahun)
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
		Status: "Success get proses bisnis by kode opd and year",
		Data:   prosesBisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProsesBisnisControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisId := params.ByName("prosesbisnisId")
	id, err := strconv.Atoi(prosesbisnisId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid ID",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	prosesbisnisResponse, err := controller.ProsesBisnisService.FindById(request.Context(), id)
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
		Status: "Success get proses bisnis by id",
		Data:   prosesbisnisResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ProsesBisnisControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisInsertRequest := web.ProsesBisnisCreateRequest{}
	helper.ReadFromRequestBody(request, &prosesbisnisInsertRequest)

	prosesbisnisResponse := controller.ProsesBisnisService.Insert(request.Context(), prosesbisnisInsertRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create proses bisnis",
		Data:   prosesbisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ProsesBisnisControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisUpdateRequest := web.ProsesBisnisUpdateRequest{}
	helper.ReadFromRequestBody(request, &prosesbisnisUpdateRequest)

	prosesbisnisId := params.ByName("prosesbisnisId")
	id, err := strconv.Atoi(prosesbisnisId)
	helper.PanicIfError(err)

	prosesbisnisUpdateRequest.Id = id

	prosesbisnisResponse := controller.ProsesBisnisService.Update(request.Context(), prosesbisnisUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update proses bisnis",
		Data:   prosesbisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ProsesBisnisControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prosesbisnisId := params.ByName("prosesbisnisId")
	id, err := strconv.Atoi(prosesbisnisId)
	helper.PanicIfError(err)

	controller.ProsesBisnisService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete proses bisnis",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProsesBisnisControllerImpl) GetProsesBisnisGrouped(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := params.ByName("kodeOPD")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error

	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		helper.PanicIfError(err)
	}

	prosesBisnisResponse, err := controller.ProsesBisnisService.GetProsesBisnisGrouped(request.Context(), kodeOPD, tahun)
	if err != nil {
		log.Printf("Error getting Gap ESPBE: %v", err)
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
		Status: "Success get GAP ESPBE",
		Data:   prosesBisnisResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProsesBisnisControllerImpl) GetProsesBisnisNoGap(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := params.ByName("kodeOPD")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error

	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		helper.PanicIfError(err)
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
