package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OpdControllerImpl struct {
	OpdService service.OpdService
}

func NewOpdControllerImpl(opdService service.OpdService) *OpdControllerImpl {
	return &OpdControllerImpl{
		OpdService: opdService,
	}
}

func (controller *OpdControllerImpl) FetchApiOpd(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	opdApiesponse, err := controller.OpdService.FetchKodeOpd(request.Context())
	if err != nil {
		helper.PanicIfError(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success fetching and inserting Opd",
		Data:   opdApiesponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
