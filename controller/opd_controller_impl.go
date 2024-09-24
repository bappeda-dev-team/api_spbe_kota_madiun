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
	opdApiesponse, err := controller.OpdService.FetchAllData(request.Context())
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

func (controller *OpdControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	opdResponses := controller.OpdService.FindAll(request.Context(), kodeOPD)

	// Filter untuk menghapus OPD dengan kode "super_admin"
	var filteredOpdResponses []web.Opd
	for _, opd := range opdResponses {
		if opd.KodeOpd != "super_admin" {
			filteredOpdResponses = append(filteredOpdResponses, opd)
		}
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan kode opd",
		Data:   filteredOpdResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OpdControllerImpl) FindAllOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	kodeOPD := ""

	opdResponses := controller.OpdService.FindAll(request.Context(), kodeOPD)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan kode opd",
		Data:   opdResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OpdControllerImpl) FindAllEksternal(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil role dari context
	role := request.Context().Value("roles").(string)

	var kodeOPD string

	// Cek role dan ambil kode_opd sesuai dengan rolenya
	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	// Panggil service untuk mendapatkan semua OPD
	allOPD := controller.OpdService.FindAll(request.Context(), "")

	// Filter OPD yang kodenya tidak sama dengan kodeOPD yang diberikan
	var filteredOPD []web.Opd
	for _, opd := range allOPD {
		if opd.KodeOpd != kodeOPD && opd.KodeOpd != "super_admin" {
			filteredOPD = append(filteredOPD, opd)
		}
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan daftar OPD eksternal",
		Data:   filteredOPD,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OpdControllerImpl) FindKodeOpdUrusan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	opdUrusanResponses, err := controller.OpdService.FindKodeOpdUrusan(request.Context(), kodeOPD)
	if err != nil {
		helper.PanicIfError(err)
	}

	// Filter untuk menghapus OPD dengan kode "super_admin"
	var filteredOpdUrusanResponses []web.OPD
	for _, opd := range opdUrusanResponses {
		if opd.KodeOpd != "super_admin" {
			filteredOpdUrusanResponses = append(filteredOpdUrusanResponses, opd)
		}
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan kode opd",
		Data:   filteredOpdUrusanResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
