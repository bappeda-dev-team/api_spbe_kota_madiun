package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AplikasiControllerImpl struct {
	aplikasiService service.AplikasiService
}

func NewAplikasiController(aplikasiService service.AplikasiService) AplikasiController {
	return &AplikasiControllerImpl{
		aplikasiService: aplikasiService,
	}
}

func (controller *AplikasiControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOPD := params.ByName("kodeOPD")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error

	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		helper.PanicIfError(err)
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
		Status: "Success get data dan informasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("aplikasiId")
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

	aplikasiResponse, err := controller.aplikasiService.FindById(request.Context(), id)
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
	plikasiCreateRequest := web.AplikasiCreateRequest{}
	helper.ReadFromRequestBody(request, &plikasiCreateRequest)

	aplikasiResponse := controller.aplikasiService.Insert(request.Context(), plikasiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create data dan informasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	AplikasiUpdateRespons := web.AplikasiUpdateRespons{}
	helper.ReadFromRequestBody(request, &AplikasiUpdateRespons)

	layananspbeId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	AplikasiUpdateRespons.Id = id

	aplikasiResponse := controller.aplikasiService.Update(request.Context(), AplikasiUpdateRespons)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update data dan informasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *AplikasiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(dataId)
	helper.PanicIfError(err)

	controller.aplikasiService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete aplikasi",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
