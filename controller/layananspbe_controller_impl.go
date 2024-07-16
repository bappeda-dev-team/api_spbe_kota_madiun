package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LayananSpbeControllerImpl struct {
	layananSpbeService service.LayananSpbeService
}

func NewLayananSPBEController(layananspbeService service.LayananSpbeService) LayananSpbeController {
	return &LayananSpbeControllerImpl{
		layananSpbeService: layananspbeService,
	}
}

func (controller *LayananSpbeControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := params.ByName("kodeOPD")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error

	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		helper.PanicIfError(err)
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

	layananspbeResponse, err := controller.layananSpbeService.FindById(request.Context(), id)
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

	layananspbeResponse := controller.layananSpbeService.Insert(request.Context(), LayananSpbeCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create layanan spbe",
		Data:   layananspbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *LayananSpbeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	LayananSpbeUpdateRequest := web.LayananSpbeUpdateRequest{}
	helper.ReadFromRequestBody(request, &LayananSpbeUpdateRequest)

	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	LayananSpbeUpdateRequest.Id = id

	layananspbeResponse := controller.layananSpbeService.Update(request.Context(), LayananSpbeUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update proses bisnis",
		Data:   layananspbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *LayananSpbeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	controller.layananSpbeService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete layanan spbe",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *LayananSpbeControllerImpl) FindByNull(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananspbeService, err := controller.layananSpbeService.FindByNull(request.Context())
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
		Status: "Success get null data layanan spbe",
		Data:   layananspbeService,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
