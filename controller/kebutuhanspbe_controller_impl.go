package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KebutuhanSPBEControllerImpl struct {
	KebutuhanSPBEService service.KebutuhanSPBEService
}

func NewKebutuhanSPBEControllerImpl(kebutuhanSPBEService service.KebutuhanSPBEService) *KebutuhanSPBEControllerImpl {
	return &KebutuhanSPBEControllerImpl{
		KebutuhanSPBEService: kebutuhanSPBEService,
	}
}

func (controller *KebutuhanSPBEControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBECreateRequest := web.KebutuhanSPBECreateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBECreateRequest)

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

	kebutuhanSPBECreateRequest.KodeOpd = kodeOPD

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.Create(request.Context(), kebutuhanSPBECreateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	kebutuhanSPBEUpdateRequest := web.KebutuhanSPBEUpdateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBEUpdateRequest)

	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	kebutuhanSPBEUpdateRequest.ID = id

	existingKebutuhanSPBE, err := controller.KebutuhanSPBEService.FindById(request.Context(), id, kodeOPD)
	if err != nil || existingKebutuhanSPBE.KodeOpd != kodeOPD {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Anda tidak memiliki akses untuk memperbarui kebutuhan SPBE ini",
		})
		return
	}

	kebutuhanSPBEUpdateRequest.KodeOpd = kodeOPD

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.Update(request.Context(), kebutuhanSPBEUpdateRequest)
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
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
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

	err = controller.KebutuhanSPBEService.Delete(request.Context(), id, kodeOPD)
	if err != nil {
		if err.Error() == "kebutuhan spbe tidak ditemukan untuk OPD ini" {
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk menghapus kebutuhan spbe ini",
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

func (controller *KebutuhanSPBEControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	role := request.Context().Value("roles").(string)
	kodeOPD := ""
	if role != "admin_kota" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.FindById(request.Context(), id, kodeOPD)
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
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) FindByKodeOpdAndTahun(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunStr := params.ByName("tahun")
	tahun, _ := strconv.Atoi(tahunStr)

	role := request.Context().Value("roles").(string)
	kodeOpd := ""

	if role != "admin_kota" {
		kodeOpd = request.Context().Value("kode_opd").(string)
	}

	kebutuhanSPBEResponses, err := controller.KebutuhanSPBEService.FindByKodeOpdAndTahun(request.Context(), kodeOpd, tahun)
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
		Data:   kebutuhanSPBEResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
