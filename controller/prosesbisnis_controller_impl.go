package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// type ReferensiArsitekturControllerImpl struct {
// 	ReferensiArsitekturService service.ReferensiArsitekturService
// }

// func NewReferensiarstitekturController(referenceService service.ReferensiArsitekturService)ReferensiArsitekturController{
// 	return &ReferensiArsitekturControllerImpl{
// 		ReferensiArsitekturService: referenceService,
// 	}
// }

type ProsesBisnisControllerImpl struct {
	ProseBisnisService service.ProsesBisnisService
}

func NewProsesBisnisController(prosbisService service.ProsesBisnisService)ProsesBisnisController{
	return &ProsesBisnisControllerImpl{
		ProseBisnisService: prosbisService,
	}
}

func (controller *ProsesBisnisControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	referensiarsitekturResponse := controller.ProseBisnisService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success get all referensi arsitektur",
		Data: referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
