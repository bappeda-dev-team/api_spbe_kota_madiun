package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KebutuhanSPBEControllerImpl struct {
	KebutuhanSPBEService service.KebutuhanSPBEService
}

func NewKebutuhanSPBEController(kebutuhanSPBEService service.KebutuhanSPBEService) KebutuhanSPBEController {
	return &KebutuhanSPBEControllerImpl{
		KebutuhanSPBEService: kebutuhanSPBEService,
	}
}

func (controller *KebutuhanSPBEControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBECreateRequest := web.KebutuhanSPBECreateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBECreateRequest)

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.Create(request.Context(), kebutuhanSPBECreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBEUpdateRequest := web.KebutuhanSPBEUpdateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBEUpdateRequest)

	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	kebutuhanSPBEUpdateRequest.ID = id

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.Update(request.Context(), kebutuhanSPBEUpdateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	err = controller.KebutuhanSPBEService.Delete(request.Context(), id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.FindById(request.Context(), id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) FindByKodeOpdAndTahun(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOpd := params.ByName("kodeOpd")
	tahunStr := params.ByName("tahun")

	var tahun int
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			webResponse := web.WebResponse{
				Code:   400,
				Status: "BAD REQUEST",
				Data:   "Tahun harus berupa angka",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
	}

	kebutuhanSPBEResponses, err := controller.KebutuhanSPBEService.FindByKodeOpdAndTahun(request.Context(), kodeOpd, tahun)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
