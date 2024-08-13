package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AplikasiControllerImpl struct {
	aplikasiService service.AplikasiService
}

func NewAplikasiControllerImpl(aplikasiService service.AplikasiService) *AplikasiControllerImpl {
	return &AplikasiControllerImpl{
		aplikasiService: aplikasiService,
	}
}

func (controller *AplikasiControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunStr := params.ByName("tahun")
	tahun, _ := strconv.Atoi(tahunStr)

	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role != "admin_kota" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	aplikasiResponse, err := controller.aplikasiService.FindByKodeOpd(request.Context(), kodeOPD, tahun)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success get data aplikasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(dataId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid ID",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := ""
	if role != "admin_kota" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	aplikasiResponse, err := controller.aplikasiService.FindById(request.Context(), id, kodeOPD)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success get aplikasi by id",
		Data:   aplikasiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *AplikasiControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aplikasiCreateRequest := web.AplikasiCreateRequest{}
	helper.ReadFromRequestBody(request, &aplikasiCreateRequest)

	kodeOPD, ok := request.Context().Value("kode_opd").(string)
	if !ok {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   "Kode OPD tidak ditemukan",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	aplikasiCreateRequest.KodeOPD = kodeOPD

	aplikasiResponse := controller.aplikasiService.Insert(request.Context(), aplikasiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create data aplikasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	if role != "asn" {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya pengguna ASN yang dapat memperbarui proses bisnis",
		})
		return
	}

	AplikasiUpdateRequest := web.AplikasiUpdateRequest{}
	helper.ReadFromRequestBody(request, &AplikasiUpdateRequest)

	layananspbeId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

	AplikasiUpdateRequest.Id = id

	existingAplikasi, err := controller.aplikasiService.FindById(request.Context(), id, kodeOPD)
	if err != nil || existingAplikasi.KodeOPD != kodeOPD {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Anda tidak memiliki akses untuk memperbarui aplikasi ini",
		})
		return
	}

	aplikasiResponse := controller.aplikasiService.Update(request.Context(), AplikasiUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update data aplikasi",
		Data:   aplikasiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *AplikasiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataId := params.ByName("aplikasiId")
	id, err := strconv.Atoi(dataId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID proses bisnis tidak valid",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	kodeOPD, ok := request.Context().Value("kode_opd").(string)
	if !ok {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   "Kode OPD tidak ditemukan",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err = controller.aplikasiService.Delete(request.Context(), id, kodeOPD)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil menghapus proses bisnis",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
