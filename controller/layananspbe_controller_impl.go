package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LayananSpbeControllerImpl struct {
	layananSpbeService service.LayananSpbeService
}

func NewLayananSPBEControllerImpl(layananspbeService service.LayananSpbeService) *LayananSpbeControllerImpl {
	return &LayananSpbeControllerImpl{
		layananSpbeService: layananspbeService,
	}
}

func (controller *LayananSpbeControllerImpl) FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunStr := params.ByName("tahun")
	tahun, _ := strconv.Atoi(tahunStr)

	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role != "admin_kota" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	layananSpbeResponse, err := controller.layananSpbeService.FindByKodeOpd(request.Context(), kodeOPD, tahun)
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
		Status: "Success get layanan spbe",
		Data:   layananSpbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
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

	layananspbeResponse, err := controller.layananSpbeService.FindById(request.Context(), id, kodeOPD)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Mengirimkan respons dengan data proses bisnis
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success get layanan spbe by id",
		Data:   layananspbeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	LayananSpbeCreateRequest := web.LayananSpbeCreateRequest{}
	helper.ReadFromRequestBody(request, &LayananSpbeCreateRequest)

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

	LayananSpbeCreateRequest.KodeOPD = kodeOPD

	layananspbeResponse := controller.layananSpbeService.Insert(request.Context(), LayananSpbeCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success create layanan spbe",
		Data:   layananspbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *LayananSpbeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	if role != "asn" {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya pengguna ASN yang dapat memperbarui layanan spbe",
		})
		return
	}

	LayananSpbeUpdateRequest := web.LayananSpbeUpdateRequest{}
	helper.ReadFromRequestBody(request, &LayananSpbeUpdateRequest)
	layananspbeId := params.ByName("layananspbeId")

	id, err := strconv.Atoi(layananspbeId)

	helper.PanicIfError(err)

	LayananSpbeUpdateRequest.Id = id

	existingLayananSpbe, err := controller.layananSpbeService.FindById(request.Context(), id, kodeOPD)
	if err != nil || existingLayananSpbe.KodeOPD != kodeOPD {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Anda tidak memiliki akses untuk memperbarui layanan spbe ini",
		})
		return
	}

	LayananSpbeUpdateRequest.KodeOPD = kodeOPD

	layananspbeResponse := controller.layananSpbeService.Update(request.Context(), LayananSpbeUpdateRequest)

	LayananSpbeUpdateRequest.KodeOPD = kodeOPD

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success update layanan spbe",
		Data:   layananspbeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LayananSpbeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananspbeId := params.ByName("layananspbeId")
	id, err := strconv.Atoi(layananspbeId)
	helper.PanicIfError(err)

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

	err = controller.layananSpbeService.Delete(request.Context(), id, kodeOPD)
	if err != nil {
		if err.Error() == "layanan spbe tidak ditemukan untuk OPD ini" {
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk menghapus layanan spbe ini",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		panic(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete layanan spbe",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
