package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RencanaPelaksanaanControllerImpl struct {
	rencanaPelaksanaanService service.RencanaPelaksanaanService
}

func NewRencanaPelaksanaanControllerImpl(rencanaPelaksanaanService service.RencanaPelaksanaanService) *RencanaPelaksanaanControllerImpl {
	return &RencanaPelaksanaanControllerImpl{rencanaPelaksanaanService: rencanaPelaksanaanService}
}

func (controller *RencanaPelaksanaanControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)

	// Ambil id_kebutuhan dari URL query parameters
	idKebutuhanStr := request.URL.Query().Get("id_kebutuhan")
	idKebutuhan, err := strconv.Atoi(idKebutuhanStr)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID kebutuhan tidak valid atau tidak ditemukan",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	rencanaPelaksanaanCreateRequest := web.RencanaPelaksanaanCreateRequest{}
	helper.ReadFromRequestBody(request, &rencanaPelaksanaanCreateRequest)

	// Set id_prosesbisnis dari query parameter
	rencanaPelaksanaanCreateRequest.IdKebutuhan = idKebutuhan

	if role == "admin_kota" {
		rencanaPelaksanaanCreateRequest.KodeOpd = request.URL.Query().Get("kode_opd")
		if rencanaPelaksanaanCreateRequest.IndikatorPD == "internal" {
			rencanaPelaksanaanCreateRequest.PerangkatDaerah = rencanaPelaksanaanCreateRequest.KodeOpd
		}
	}

	if role == "admin_opd" || role == "asn" {
		kodeOpd := request.Context().Value("kode_opd").(string)
		rencanaPelaksanaanCreateRequest.KodeOpd = kodeOpd
		if rencanaPelaksanaanCreateRequest.IndikatorPD == "internal" {
			rencanaPelaksanaanCreateRequest.PerangkatDaerah = kodeOpd
		}
	}

	rencanaPelaksanaanCreateResponse, err := controller.rencanaPelaksanaanService.Create(request.Context(), rencanaPelaksanaanCreateRequest)
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
		Code:   200,
		Status: "OK",
		Data:   rencanaPelaksanaanCreateResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaPelaksanaanControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	rencanaPelaksanaanUpdateRequest := web.RencanaPelaksanaanUpdateRequest{}
	helper.ReadFromRequestBody(request, &rencanaPelaksanaanUpdateRequest)

	rencanaPelaksanaanId := params.ByName("rencanaId")
	id, err := strconv.Atoi(rencanaPelaksanaanId)
	helper.PanicIfError(err)

	rencanaPelaksanaanUpdateRequest.Id = id

	// cek == kode opd untuk roles asn dan admin_opd
	if role == "admin_opd" || role == "asn" {
		if rencanaPelaksanaanUpdateRequest.IndikatorPD == "internal" {
			rencanaPelaksanaanUpdateRequest.PerangkatDaerah = kodeOPD
		}
		existingRencanaPelaksanaan, err := controller.rencanaPelaksanaanService.FindById(request.Context(), id, kodeOPD)
		if err != nil || existingRencanaPelaksanaan.KodeOpd != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk memperbarui Rencana Pelaksanaan ini",
			})
			return
		}
		rencanaPelaksanaanUpdateRequest.KodeOpd = kodeOPD
	} else if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
		rencanaPelaksanaanUpdateRequest.KodeOpd = kodeOPD

		if rencanaPelaksanaanUpdateRequest.IndikatorPD == "internal" {
			rencanaPelaksanaanUpdateRequest.PerangkatDaerah = rencanaPelaksanaanUpdateRequest.KodeOpd
		}
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Role tidak diizinkan untuk memperbarui proses bisnis",
		})
		return
	}

	rencanaPelaksanaanUpdateRequest.KodeOpd = kodeOPD

	rencanaPelaksanaanUpdateResponse, err := controller.rencanaPelaksanaanService.Update(request.Context(), rencanaPelaksanaanUpdateRequest)
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
		Data:   rencanaPelaksanaanUpdateResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaPelaksanaanControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaPelaksanaanId := params.ByName("rencanaId")
	id, err := strconv.Atoi(rencanaPelaksanaanId)
	helper.PanicIfError(err)

	kodeOPD, _ := request.Context().Value("kode_opd").(string)
	role, _ := request.Context().Value("roles").(string)

	err = controller.rencanaPelaksanaanService.Delete(request.Context(), id, kodeOPD, role)
	if err != nil {
		if err.Error() == "rencana pelaksanaan tidak ditemukan untuk OPD ini" {
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk menghapus rencana pelaksanaan ini",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		panic(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete rencana pelaksanaan",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaPelaksanaanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanIdStr := request.URL.Query().Get("kebutuhanId")
	kebutuhanId := 0
	var err error
	if kebutuhanIdStr != "" {
		kebutuhanId, err = strconv.Atoi(kebutuhanIdStr)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Id kebutuhan tidak valid",
				Data:   nil,
			})
			return
		}
	}

	role := request.Context().Value("roles").(string)
	kodeOpd := ""

	if role == "admin_kota" {
		kodeOpd = request.URL.Query().Get("kode_opd")
	} else {
		kodeOpd = request.Context().Value("kode_opd").(string)
	}

	// Tambahkan log untuk kode OPD
	log.Printf("Kode OPD: %s", kodeOpd)

	rencanaPelaksanaanResponses, err := controller.rencanaPelaksanaanService.FindAll(request.Context(), kodeOpd, kebutuhanId)
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
		Data:   rencanaPelaksanaanResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaPelaksanaanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaId := params.ByName("rencanaId")
	id, err := strconv.Atoi(rencanaId)
	helper.PanicIfError(err)

	role := request.Context().Value("roles").(string)
	kodeOpd := ""
	if role != "admin_kota" {
		kodeOpd = request.Context().Value("kode_opd").(string)
	}

	rencanaPelaksanaanResponse, err := controller.rencanaPelaksanaanService.FindById(request.Context(), id, kodeOpd)
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
		Data:   rencanaPelaksanaanResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
