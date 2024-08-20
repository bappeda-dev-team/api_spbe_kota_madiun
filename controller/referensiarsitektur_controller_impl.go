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

func NewReferensiArsitekturControllerImpl(referenceService service.ReferensiArsitekturService) *ReferensiArsitekturControllerImpl {
	return &ReferensiArsitekturControllerImpl{
		ReferensiArsitekturService: referenceService,
	}
}

func (controller *ReferensiArsitekturControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	referensiarsitekturInsertRequest := web.ReferensiArsitekturCreateRequest{}
	helper.ReadFromRequestBody(request, &referensiarsitekturInsertRequest)

	referensiarsitekturResponse := controller.ReferensiArsitekturService.Insert(request.Context(), referensiarsitekturInsertRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create referensi arsitektur",
		Data:   referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	referensiarsitekturUpdateRequest := web.ReferensiArsitekturUpdateRequest{}
	helper.ReadFromRequestBody(request, &referensiarsitekturUpdateRequest)

	referensiarsitekturId := params.ByName("referensiarsitekturId")
	id, err := strconv.Atoi(referensiarsitekturId)
	helper.PanicIfError(err)

	referensiarsitekturUpdateRequest.Id = id

	referensiarsitekturResponse := controller.ReferensiArsitekturService.Update(request.Context(), referensiarsitekturUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update referensi arsitektur",
		Data:   referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	referensiarsitekturId := params.ByName("referensiarsitekturId")
	id, err := strconv.Atoi(referensiarsitekturId)
	helper.PanicIfError(err)

	controller.ReferensiArsitekturService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete referensi arsitektur",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil parameter tahun dari query string
	tahun := request.URL.Query().Get("tahun")

	var tahunInt int
	var err error

	// Jika tahun tidak kosong, konversi ke integer
	if tahun != "" {
		tahunInt, err = strconv.Atoi(tahun)
		if err != nil {
			webResponse := web.WebResponse{
				Code:   400,
				Status: "Bad Request",
				Data:   "Tahun harus berupa angka",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
	}

	referensiarsitekturResponse := controller.ReferensiArsitekturService.FindAll(request.Context(), tahunInt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan semua referensi arsitektur",
		Data:   referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReferensiArsitekturControllerImpl) FindByKodeReferensi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	kodeReferensi := params.ByName("kodeReferensi")

	referensiarsitekturResponse, err := controller.ReferensiArsitekturService.GetDataHierarchy(request.Context(), kodeReferensi)
	if err != nil {
		if err.Error() == "data not found" {
			webResponse := web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   "Data not found",
			}
			helper.WriteToResponseBody(writer, webResponse)
		} else {
			webResponse := web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Data:   err.Error(),
			}
			helper.WriteToResponseBody(writer, webResponse)
		}
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success get code referensi arsitektur",
		Data:   referensiarsitekturResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *ReferensiArsitekturControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	referensiarsitekturId := params.ByName("referensiarsitekturId")
	id, err := strconv.Atoi(referensiarsitekturId)
	helper.PanicIfError(err)

	referensiarsitekturResponse := controller.ReferensiArsitekturService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Succes get code referensi arsitektur by id",
		Data:   referensiarsitekturResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
