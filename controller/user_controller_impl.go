package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}
func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var loginRequest web.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse, err := controller.userService.Login(request.Context(), loginRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(loginResponse)
}

func (controller *UserControllerImpl) InsertApi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// kodeOPD := request.URL.Query().Get("kode_opd")
	// tahun := request.URL.Query().Get("tahun")

	// if kodeOPD == "" {
	// 	http.Error(writer, "Parameter kode_opd harus diisi", http.StatusBadRequest)
	// 	return
	// }

	tahun := request.URL.Query().Get("tahun")
	if tahun != "" {
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
	} else {
		kodeOPD = request.Context().Value("kode_opd").(string)
	}

	result, err := controller.userService.InsertApi(request.Context(), kodeOPD, tahun)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(result)
}
