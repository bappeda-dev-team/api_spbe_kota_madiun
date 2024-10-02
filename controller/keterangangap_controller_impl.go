package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"context"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KeteranganGapControllerImpl struct {
	KeteranganGapService service.KeteranganGapService
}

func NewKeteranganGapControllerImpl(keteranganGapService service.KeteranganGapService) *KeteranganGapControllerImpl {
	return &KeteranganGapControllerImpl{
		KeteranganGapService: keteranganGapService,
	}
}

func (controller *KeteranganGapControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil roles dari context
	roles := request.Context().Value("roles").(string)

	keteranganGapCreateRequest := web.KeteranganGapCreateRequest{}
	helper.ReadFromRequestBody(request, &keteranganGapCreateRequest)

	// Ambil kode_opd berdasarkan roles
	if roles == "asn" || roles == "admin_opd" {
		kodeOPD := request.Context().Value("kode_opd").(string)
		keteranganGapCreateRequest.KodeOpd = kodeOPD
	} else if roles == "admin_kota" {
		kodeOPD := request.URL.Query().Get("kode_opd")
		keteranganGapCreateRequest.KodeOpd = kodeOPD
	}

	// Ambil prosesbisnis dari query URL
	prosesBisnis := request.URL.Query().Get("prosesbisnis")
	idProsesBisnis, err := strconv.Atoi(prosesBisnis)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID proses bisnis tidak valid",
		})
		return
	}
	keteranganGapCreateRequest.IdProsesBisnis = idProsesBisnis

	keteranganGapResponse, err := controller.KeteranganGapService.Insert(context.Background(), keteranganGapCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   keteranganGapResponse,
	})
}

func (controller *KeteranganGapControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	keteranganGapResponse, err := controller.KeteranganGapService.FindById(context.Background(), id)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   keteranganGapResponse,
	})
}

func (controller *KeteranganGapControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keteranganGapUpdateRequest := web.KeteranganGapUpdateRequest{}
	helper.ReadFromRequestBody(request, &keteranganGapUpdateRequest)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}
	keteranganGapUpdateRequest.Id = id

	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	if role == "admin_kota" {
		kodeOpdBaru := request.URL.Query().Get("kode_opd")
		if kodeOpdBaru == "" {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Kode OPD harus diisi untuk admin kota",
			})
			return
		}
		existingKeteranganGap, err := controller.KeteranganGapService.FindById(request.Context(), id)
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Data:   "Gagal memeriksa keterangan gap yang ada",
			})
			return
		}
		if existingKeteranganGap.KodeOpd != kodeOpdBaru {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Kode OPD tidak sesuai dengan keterangan gap yang ada",
			})
			return
		}
		keteranganGapUpdateRequest.KodeOpd = kodeOpdBaru
	} else if role == "admin_opd" || role == "asn" {
		existingKeteranganGap, err := controller.KeteranganGapService.FindById(request.Context(), id)
		if err != nil || existingKeteranganGap.KodeOpd != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk memperbarui Keterangan Gap ini",
			})
			return
		}
		keteranganGapUpdateRequest.KodeOpd = kodeOPD
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Role tidak diizinkan untuk memperbarui Keterangan Gap",
		})
		return
	}

	keteranganGapResponse, err := controller.KeteranganGapService.Update(request.Context(), keteranganGapUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   keteranganGapResponse,
	})
}

func (controller *KeteranganGapControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	var keteranganGapResponse web.KeteranganGapRespons
	var deleteErr error

	if role == "admin_kota" {
		kodeOPDQuery := request.URL.Query().Get("kode_opd")
		if kodeOPDQuery == kodeOPD {
			keteranganGapResponse, deleteErr = controller.KeteranganGapService.Delete(request.Context(), id)
		} else {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk menghapus Keterangan Gap ini",
			})
			return
		}
	} else if role == "admin_opd" || role == "asn" {
		existingKeteranganGap, err := controller.KeteranganGapService.FindById(request.Context(), id)
		if err != nil || existingKeteranganGap.KodeOpd != kodeOPD {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Anda tidak memiliki akses untuk menghapus Keterangan Gap ini",
			})
			return
		}
		keteranganGapResponse, deleteErr = controller.KeteranganGapService.Delete(request.Context(), id)
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Role tidak diizinkan untuk menghapus Keterangan Gap",
		})
		return
	}

	if deleteErr != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   deleteErr.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   keteranganGapResponse,
	})
}
