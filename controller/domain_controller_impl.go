package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type DomainSPBEControllerImpl struct {
	DomainSPBEService service.DomainSPBEService
}

func NewDomainSPBEControllerImpl(domainSPBEService service.DomainSPBEService) *DomainSPBEControllerImpl {
	return &DomainSPBEControllerImpl{
		DomainSPBEService: domainSPBEService,
	}
}

func (controller *DomainSPBEControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	domainSPBECreateRequest := web.DomainSPBECreateRequest{}
	helper.ReadFromRequestBody(request, &domainSPBECreateRequest)

	domainSPBEResponse := controller.DomainSPBEService.Insert(request.Context(), domainSPBECreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   domainSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DomainSPBEControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	domainSPBEUpdateRequest := web.DomainSPBEUpdateRequest{}
	helper.ReadFromRequestBody(request, &domainSPBEUpdateRequest)

	domainSPBEId := params.ByName("domainSPBEId")
	id, err := strconv.Atoi(domainSPBEId)
	helper.PanicIfError(err)

	domainSPBEUpdateRequest.Id = id

	domainSPBEResponse := controller.DomainSPBEService.Update(request.Context(), domainSPBEUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   domainSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DomainSPBEControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	domainSPBEId := params.ByName("domainSPBEId")
	id, err := strconv.Atoi(domainSPBEId)
	helper.PanicIfError(err)

	controller.DomainSPBEService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DomainSPBEControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	domainSPBEResponses := controller.DomainSPBEService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   domainSPBEResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DomainSPBEControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	domainSPBEId := params.ByName("domainSPBEId")
	id, err := strconv.Atoi(domainSPBEId)
	helper.PanicIfError(err)

	domainSPBEResponse := controller.DomainSPBEService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   domainSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
