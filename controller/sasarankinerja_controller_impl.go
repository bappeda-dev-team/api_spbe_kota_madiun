package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SasaranKinerjaPegawaiControllerImpl struct {
	SasaranKinerjaPegawaiService service.SasaranKinerjaPegawaiService
}

func NewSasaranKinerjaPegawaiControllerImpl(SasaranKinerjaPegawaiService service.SasaranKinerjaPegawaiService) *SasaranKinerjaPegawaiControllerImpl {
	return &SasaranKinerjaPegawaiControllerImpl{
		SasaranKinerjaPegawaiService: SasaranKinerjaPegawaiService,
	}
}

func (controller *SasaranKinerjaPegawaiControllerImpl) FetchsasaranKinerja(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
	}
	if role == "admin_opd" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	result, err := controller.SasaranKinerjaPegawaiService.FetchsasaranKinerja(request.Context(), kodeOPD, tahun)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(result)
}

func (controller *SasaranKinerjaPegawaiControllerImpl) FindByKodeOpdAndTahun(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahun := request.URL.Query().Get("tahun_sasaran")

	role := request.Context().Value("roles").(string)
	kodeOpd := ""

	fmt.Println("Role:", role)
	if role == "admin_kota" {
		kodeOpd = request.URL.Query().Get("kode_opd")
		fmt.Println("Kode OPD (admin_kota):", kodeOpd)
	} else { // Untuk role lainnya
		kodeOpd = request.Context().Value("kode_opd").(string)
		fmt.Println("Kode OPD (role lainnya):", kodeOpd)
	}

	sasaranKinerjaResponse := controller.SasaranKinerjaPegawaiService.FindByKodeOpdAndTahun(request.Context(), kodeOpd, tahun)

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil mendapatkan sasaran kinerja pegawai",
		Data:   sasaranKinerjaResponse,
	})
}

func (controller *SasaranKinerjaPegawaiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sasaranKinerjaId := params.ByName("sasaranKinerjaId")
	id, err := strconv.Atoi(sasaranKinerjaId)
	helper.PanicIfError(err)

	sasaranResponse, err := controller.SasaranKinerjaPegawaiService.FindById(request.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Succes get pohon kinerja by id",
		Data:   sasaranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
