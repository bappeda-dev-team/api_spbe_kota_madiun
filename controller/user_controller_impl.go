package controller

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/service"
	"encoding/json"
	"net/http"
	"strconv"

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

func (controller *UserControllerImpl) ChangePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var changePasswordRequest web.ChangePasswordRequest
	err := json.NewDecoder(request.Body).Decode(&changePasswordRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Format permintaan tidak valid",
		})
		return
	}

	role := request.Context().Value("roles").(string)
	var userID int

	if role == "admin_kota" {
		userID, err = strconv.Atoi(params.ByName("userId"))
		if err != nil {
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "ID pengguna tidak valid",
			})
			return
		}
	} else {
		userIDValue := request.Context().Value("user_id")
		switch v := userIDValue.(type) {
		case int:
			userID = v
		case float64:
			userID = int(v)
		case string:
			userID, err = strconv.Atoi(v)
			if err != nil {
				helper.WriteToResponseBody(writer, web.WebResponse{
					Code:   http.StatusInternalServerError,
					Status: "INTERNAL SERVER ERROR",
					Data:   "ID pengguna tidak valid dalam token",
				})
				return
			}
		default:
			helper.WriteToResponseBody(writer, web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Data:   "Tipe ID pengguna tidak valid",
			})
			return
		}
	}

	loginResponse, err := controller.userService.ChangePassword(request.Context(), userID, changePasswordRequest)
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
		Data:   loginResponse,
	})
}
