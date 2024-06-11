package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PohonKinerjaControllerImpl struct {
	PohonKinerjaService service.PohonKinerjaService
}

func NewPohonKinerjaController(pohonService service.PohonKinerjaService) PohonKinerjaController {
	return &PohonKinerjaControllerImpl{
		PohonKinerjaService: pohonService,
	}
}

func (controller *PohonKinerjaControllerImpl)FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	pohonId := params.ByName("pohonId")
	id, err := strconv.Atoi(pohonId)
	helper.PanicIfError(err)

	pohonResponse := controller.PohonKinerjaService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Succes get pohon kinerja by id",
		Data:   pohonResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}