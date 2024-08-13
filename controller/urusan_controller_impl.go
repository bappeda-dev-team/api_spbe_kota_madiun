package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UrusanControllerImpl struct {
	UrusanService service.UrusanService
}

func NewUrusanControllerImpl(urusanService service.UrusanService) *UrusanControllerImpl {
	return &UrusanControllerImpl{
		UrusanService: urusanService,
	}
}
func (controller *UrusanControllerImpl) FetchApiUrusan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	urusanApiResponse, err := controller.UrusanService.FetchUrusan(request.Context())
	if err != nil {
		helper.PanicIfError(err)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success fetching and inserting Urusan",
		Data:   urusanApiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
