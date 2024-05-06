package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ReferensiArsitekturControllerImpl struct {
	ReferensiArsitekturService service.ReferensiArsitekturService
}

func NewReferensiarstitekturController(referenceService service.ReferensiArsitekturService)ReferensiArsitekturController{
	return &ReferensiArsitekturControllerImpl{
		ReferensiArsitekturService: referenceService,
	}
}

func (controller *ReferensiArsitekturControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	referensiarsitekturInsertRequest := web.ReferensiArsitekturCreateRequest{}
	helper.ReadFromRequestBody(request, &referensiarsitekturInsertRequest)

	referensiarsitekturResponse := controller.ReferensiArsitekturService.Insert(request.Context(), referensiarsitekturInsertRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	referensiarsitekturUpdateRequest := web.ReferensiArsitekturUpdateRequest{}
	helper.ReadFromRequestBody(request, &referensiarsitekturUpdateRequest)

	referensiarsitekturId := params.ByName("referensiarsitekturId")
	id,err := strconv.Atoi(referensiarsitekturId)
	helper.PanicIfError(err)
	
	referensiarsitekturUpdateRequest.Id = id

	referensiarsitekturResponse := controller.ReferensiArsitekturService.Update(request.Context(),referensiarsitekturUpdateRequest)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	referensiarsitekturId := params.ByName("referensiarsitekturId")
	id,err := strconv.Atoi(referensiarsitekturId)
	helper.PanicIfError(err)

	controller.ReferensiArsitekturService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	referensiarsitekturResponse := controller.ReferensiArsitekturService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}