package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BidangUrusanControllerImpl struct {
	BidangUrusanService service.BidangUrusanService
}

func NewBidangUrusanControllerImpl(bidangService service.BidangUrusanService) *BidangUrusanControllerImpl {
	return &BidangUrusanControllerImpl{
		BidangUrusanService: bidangService,
	}
}

func (controller *BidangUrusanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bidangService := controller.BidangUrusanService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success get all Bidang Urusan",
		Data:   bidangService,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BidangUrusanControllerImpl) FetchBidangUrusan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bidangService, err := controller.BidangUrusanService.FetchUrusan(request.Context())
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success Fetch bidang urusan",
		Data:   bidangService,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BidangUrusanControllerImpl) FindBidangUrusanOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)

	var kodeOPD string
	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	ctx := context.Background()

	bidangUrusanResponses, err := controller.BidangUrusanService.FindBidangUrusanOPD(ctx, kodeOPD)
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
		Status: "OK",
		Data:   bidangUrusanResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
