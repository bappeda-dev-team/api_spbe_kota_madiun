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

func NewPohonKinerjaControllerImpl(pohonService service.PohonKinerjaService) *PohonKinerjaControllerImpl {
	return &PohonKinerjaControllerImpl{
		PohonKinerjaService: pohonService,
	}
}

func (controller *PohonKinerjaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

func (controller *PohonKinerjaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunParam := request.URL.Query().Get("tahun")
	var tahun int
	var err error
	if tahunParam != "" {
		tahun, err = strconv.Atoi(tahunParam)
		if err != nil {
			http.Error(writer, "Parameter tahun harus berupa angka", http.StatusBadRequest)
			return
		}
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	// Hanya admin_kota yang dapat menggunakan parameter kode_opd
	if role == "admin_kota" {
		if paramKodeOPD := request.URL.Query().Get("kode_opd"); paramKodeOPD != "" {
			kodeOPD = paramKodeOPD
		}
	}

	pohonResponse := controller.PohonKinerjaService.FindAll(request.Context(), kodeOPD, tahun)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil mendapatkan semua pohon kinerja",
		Data:   pohonResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PohonKinerjaControllerImpl) FetchApiPohon(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahun := request.URL.Query().Get("tahun")
	if tahun == "" {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Format tahun tidak valid",
			Data:   nil,
		})
		return
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	_, err := controller.PohonKinerjaService.InsertApi(request.Context(), kodeOPD, tahun)
	if err != nil {
		helper.PanicIfError(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success fetching and inserting Pohon Kinerja",
		Data:   "berhasil fetch",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PohonKinerjaControllerImpl) FindByOperational(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pohonId := request.URL.Query().Get("pohonId")
	if pohonId == "" {
		pohonId = params.ByName("pohonId")
	}
	id, err := strconv.Atoi(pohonId)
	helper.PanicIfError(err)

	pohonKinerjaResponse := controller.PohonKinerjaService.FindByOperational(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Berhasil",
		Data:   pohonKinerjaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PohonKinerjaControllerImpl) GetHierarchy(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pohonId := params.ByName("pohonId")
	id, err := strconv.Atoi(pohonId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	hierarchyResponse, err := controller.PohonKinerjaService.GetHierarchy(request.Context(), id)
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
		Data:   hierarchyResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
