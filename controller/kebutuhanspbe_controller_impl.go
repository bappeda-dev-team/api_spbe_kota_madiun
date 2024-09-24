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

type KebutuhanSPBEControllerImpl struct {
	KebutuhanSPBEService service.KebutuhanSPBEService
}

func NewKebutuhanSPBEControllerImpl(kebutuhanSPBEService service.KebutuhanSPBEService) *KebutuhanSPBEControllerImpl {
	return &KebutuhanSPBEControllerImpl{
		KebutuhanSPBEService: kebutuhanSPBEService,
	}
}

func (controller *KebutuhanSPBEControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)

	if role != "admin_opd" && role != "admin_kota" {
		webResponse := web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya admin_opd dan admin_kota yang diizinkan untuk membuat kebutuhan SPBE",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Ambil id_prosesbisnis dari URL query parameters
	idProsesBisnisStr := request.URL.Query().Get("id_prosesbisnis")
	idProsesBisnis, err := strconv.Atoi(idProsesBisnisStr)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID proses bisnis tidak valid atau tidak ditemukan",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	kebutuhanSPBECreateRequest := web.KebutuhanSPBECreateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBECreateRequest)

	// Set id_prosesbisnis dari query parameter
	kebutuhanSPBECreateRequest.IdProsesbisnis = idProsesBisnis

	if role == "admin_kota" {
		kebutuhanSPBECreateRequest.KodeOpd = request.URL.Query().Get("kode_opd")
	}

	if role == "admin_opd" {
		kodeOPD := request.Context().Value("kode_opd").(string)
		kebutuhanSPBECreateRequest.KodeOpd = kodeOPD
	}

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.Create(request.Context(), kebutuhanSPBECreateRequest)
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
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// func (controller *KebutuhanSPBEControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	role := request.Context().Value("roles").(string)
// 	kodeOPD := request.Context().Value("kode_opd").(string)

// 	kebutuhanSPBEUpdateRequest := web.KebutuhanSPBEUpdateRequest{}
// 	helper.ReadFromRequestBody(request, &kebutuhanSPBEUpdateRequest)

// 	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
// 	id, err := strconv.Atoi(kebutuhanSPBEId)
// 	helper.PanicIfError(err)

// 	kebutuhanSPBEUpdateRequest.ID = id

// 	// cek == kode opd untuk roles asn dan admin_opd
// 	if role == "admin_opd" || role == "asn" {
// 		existingKebutuhanSPBE, err := controller.KebutuhanSPBEService.FindById(request.Context(), id, kodeOPD)
// 		if err != nil || existingKebutuhanSPBE.KodeOpd != kodeOPD {
// 			helper.WriteToResponseBody(writer, web.WebResponse{
// 				Code:   http.StatusForbidden,
// 				Status: "FORBIDDEN",
// 				Data:   "Anda tidak memiliki akses untuk memperbarui Kebutuhan SPBE ini",
// 			})
// 			return
// 		}
// 		kebutuhanSPBEUpdateRequest.KodeOpd = kodeOPD
// 	} else if role == "admin_kota" {
// 		kebutuhanSPBEUpdateRequest.KodeOpd = request.URL.Query().Get("kode_opd")
// 	} else {
// 		helper.WriteToResponseBody(writer, web.WebResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "UNAUTHORIZED",
// 			Data:   "Role tidak diizinkan untuk memperbarui proses bisnis",
// 		})
// 		return
// 	}

// 	kebutuhanSPBEUpdateRequest.KodeOpd = kodeOPD

// 	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.Update(request.Context(), kebutuhanSPBEUpdateRequest)
// 	if err != nil {
// 		webResponse := web.WebResponse{
// 			Code:   500,
// 			Status: "INTERNAL SERVER ERROR",
// 			Data:   err.Error(),
// 		}
// 		helper.WriteToResponseBody(writer, webResponse)
// 		return
// 	}

// 	webResponse := web.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   kebutuhanSPBEResponse,
// 	}

// 	helper.WriteToResponseBody(writer, webResponse)
// }

func (controller *KebutuhanSPBEControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	kebutuhanSPBEUpdateRequest := web.KebutuhanSPBEUpdateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBEUpdateRequest)

	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	kebutuhanSPBEUpdateRequest.ID = id

	if role == "admin_kota" {
		kebutuhanSPBEUpdateRequest.KodeOpd = request.URL.Query().Get("kode_opd")

		if kebutuhanSPBEUpdateRequest.IndikatorPj == "internal" {
			kebutuhanSPBEUpdateRequest.PenanggungJawab = kebutuhanSPBEUpdateRequest.KodeOpd
		}

	} else if role == "admin_opd" || role == "asn" {
		if kebutuhanSPBEUpdateRequest.IndikatorPj == "internal" {
			kebutuhanSPBEUpdateRequest.PenanggungJawab = kodeOPD
		}
		existingKebutuhanSPBE, err := controller.KebutuhanSPBEService.FindById(request.Context(), id, kodeOPD)
		if err != nil || existingKebutuhanSPBE.KodeOpd != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk memperbarui Kebutuhan SPBE ini",
			})
			return
		}
		kebutuhanSPBEUpdateRequest.KodeOpd = kodeOPD
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Role tidak diizinkan untuk memperbarui proses bisnis",
		})
		return
	}

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

func (controller *KebutuhanSPBEControllerImpl) UpdateKeterangan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)

	if role != "admin_opd" && role != "admin_kota" {
		webResponse := web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya admin_opd dan admin_kota yang diizinkan untuk membuat kebutuhan SPBE",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	kebutuhanSPBEKeteranganUpdateRequest := web.KebutuhanSPBEKeteranganUpdateRequest{}
	helper.ReadFromRequestBody(request, &kebutuhanSPBEKeteranganUpdateRequest)

	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	kebutuhanSPBEKeteranganUpdateRequest.ID = id

	kebutuhanSPBEKeteranganResponse, err := controller.KebutuhanSPBEService.UpdateKeterangan(request.Context(), kebutuhanSPBEKeteranganUpdateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEKeteranganResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *KebutuhanSPBEControllerImpl) UpdatePenanggungJawab(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)

	if role != "admin_opd" && role != "admin_kota" {
		webResponse := web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   "Hanya admin_opd dan admin_kota yang diizinkan untuk membuat kebutuhan SPBE",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	KebutuhanSPBEPjUpdateRequest := web.KebutuhanSPBEPjUpdateRequest{}
	helper.ReadFromRequestBody(request, &KebutuhanSPBEPjUpdateRequest)

	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	KebutuhanSPBEPjUpdateRequest.ID = id

	kebutuhanSPBEPjResponse, err := controller.KebutuhanSPBEService.UpdatePenanggungJawab(request.Context(), KebutuhanSPBEPjUpdateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kebutuhanSPBEPjResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KebutuhanSPBEControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanSPBEId := params.ByName("kebutuhanSPBEId")
	id, err := strconv.Atoi(kebutuhanSPBEId)
	helper.PanicIfError(err)

	kodeOPD, _ := request.Context().Value("kode_opd").(string)
	role, _ := request.Context().Value("roles").(string)

	err = controller.KebutuhanSPBEService.Delete(request.Context(), id, kodeOPD, role)
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
	tahunStr := request.URL.Query().Get("tahun")
	tahun := 0
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Format tahun tidak valid",
				Data:   nil,
			})
			return
		}
	}
	prosesbisnisStr := request.URL.Query().Get("prosesbisnis")
	prosesbisnis := 0
	var err2 error
	if prosesbisnisStr != "" {
		prosesbisnis, err2 = strconv.Atoi(prosesbisnisStr)
		if err2 != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Id prosesbisnis tidak valid",
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

	kebutuhanSPBEResponses, err := controller.KebutuhanSPBEService.FindByKodeOpdAndTahun(request.Context(), kodeOpd, tahun, prosesbisnis)
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

func (controller *KebutuhanSPBEControllerImpl) FindDataPemenuhanKebutuhan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunStr := request.URL.Query().Get("tahun")
	tahun := 0
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Format tahun tidak valid",
				Data:   nil,
			})
			return
		}
	}
	prosesbisnisStr := request.URL.Query().Get("prosesbisnis")
	prosesbisnis := 0
	var err2 error
	if prosesbisnisStr != "" {
		prosesbisnis, err2 = strconv.Atoi(prosesbisnisStr)
		if err2 != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Id prosesbisnis tidak valid",
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

	kebutuhanSPBEResponses, err := controller.KebutuhanSPBEService.FindDataPemenuhanKebutuhan(request.Context(), kodeOpd, tahun, prosesbisnis)
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

func (controller *KebutuhanSPBEControllerImpl) FindPenanggungJawab(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	pj := ""

	if role == "admin_kota" {
		pj = request.URL.Query().Get("kode_opd")
	} else {
		pj = request.Context().Value("kode_opd").(string)
	}

	tahunStr := request.URL.Query().Get("tahun")
	tahun := 0
	var err error
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Format tahun tidak valid",
				Data:   nil,
			})
			return
		}
	}

	kebutuhanSPBEResponses, err := controller.KebutuhanSPBEService.FindPenanggungJawab(request.Context(), pj, tahun)
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

func (controller *KebutuhanSPBEControllerImpl) FindByIdPenanggungJawab(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kebutuhanId := params.ByName("kebutuhanId")
	id, err := strconv.Atoi(kebutuhanId)
	helper.PanicIfError(err)

	role := request.Context().Value("roles").(string)
	pj := ""

	if role == "admin_kota" {
		pj = request.URL.Query().Get("kode_opd")
	} else {
		pj = request.Context().Value("kode_opd").(string)
	}

	// Tambahkan logging untuk debug
	log.Printf("Controller - Role: %s, PJ: %s", role, pj)

	kebutuhanSPBEResponse, err := controller.KebutuhanSPBEService.FindByIdPenanggungJawab(request.Context(), id, role, pj)
	if err != nil {
		if err.Error() == "penanggung jawab tidak sesuai" {
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses ke data ini",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		if err.Error() == "data not found" {
			webResponse := web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   "Data tidak ditemukan",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
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
		Status: "OK",
		Data:   kebutuhanSPBEResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
