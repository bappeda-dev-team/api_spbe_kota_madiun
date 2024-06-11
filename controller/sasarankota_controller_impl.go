package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SasaranKotaControllerImpl struct {
	SasaranKotaService service.SasaranKotaService
}

func NewSasaranKotaController(sasaranService service.SasaranKotaService) SasaranKotaController {
	return &SasaranKotaControllerImpl{
		SasaranKotaService: sasaranService,
	}
}

func (controller *SasaranKotaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sasarankotaId := params.ByName("sasarankotaId")
	id, err := strconv.Atoi(sasarankotaId)
	helper.PanicIfError(err)

	sasaranResponse := controller.SasaranKotaService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Succes get sasaran kota by id",
		Data:   sasaranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
