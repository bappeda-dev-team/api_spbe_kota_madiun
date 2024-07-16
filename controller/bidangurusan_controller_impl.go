package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BidangUrusanControllerImpl struct {
	BidangUrusanService service.BidangUrusanService
}

func NewBidangUrusanController(bidangService service.BidangUrusanService) BidangUrusanController {
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
