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

func NewSasaranKotaControllerImpl(sasaranService service.SasaranKotaService) *SasaranKotaControllerImpl {
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

func (controller *SasaranKotaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunParam := request.URL.Query().Get("tahun")
	tahun := 0

	if tahunParam != "" {
		var err error
		tahun, err = strconv.Atoi(tahunParam)
		if err != nil {
			http.Error(writer, "Parameter tahun tidak valid", http.StatusBadRequest)
			return
		}
	}

	sasarankotaResponse := controller.SasaranKotaService.FindAll(request.Context(), tahun)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan semua sasaran kota",
		Data:   sasarankotaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SasaranKotaControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sasarankotaApiResponse, err := controller.SasaranKotaService.InsertApi(request.Context())
	if err != nil {
		helper.PanicIfError(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success fetching and inserting Sasaran Kota",
		Data:   sasarankotaApiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
