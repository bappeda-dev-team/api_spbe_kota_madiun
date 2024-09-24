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
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Format permintaan tidak valid",
			Data:   nil,
		})
		return
	}

	loginResponse, err := controller.userService.Login(request.Context(), loginRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(loginResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	role := request.Context().Value("roles").(string)
	kodeOPD := ""

	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else if role == "admin_opd" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "Akses ditolak",
			Data:   nil,
		})
		return
	}

	users := controller.userService.FindAll(request.Context(), kodeOPD)

	// Kirim response sukses
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil mendapatkan daftar pengguna",
		Data:   users,
	})
}

func (controller *UserControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(params.ByName("userId"))
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "ID pengguna tidak valid",
			Data:   nil,
		})
		return
	}

	// Ambil role dan kodeOPD dari context
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	var user web.UserResponse

	if role == "admin_kota" {
		// Admin kota dapat melihat semua pengguna
		user, err = controller.userService.FindByID(request.Context(), userID, "")
	} else if role == "admin_opd" {
		// Admin OPD hanya dapat melihat pengguna dengan kodeOPD yang sama
		user, err = controller.userService.FindByID(request.Context(), userID, kodeOPD)
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "Akses ditolak",
			Data:   nil,
		})
		return
	}

	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Pengguna tidak ditemukan",
			Data:   nil,
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil mendapatkan data pengguna",
		Data:   user,
	})
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
	} else if role == "admin_opd" {
		kodeOPD = request.Context().Value("kode_opd").(string)
	} else {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "Akses ditolak",
			Data:   nil,
		})
		return
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
	// Decode request body
	var changePasswordRequest web.ChangePasswordRequest
	err := json.NewDecoder(request.Body).Decode(&changePasswordRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Format permintaan tidak valid",
			Data:   nil,
		})
		return
	}

	// Ambil user ID dari context (diasumsikan telah diset oleh middleware autentikasi)
	userIDFloat64 := request.Context().Value("user_id").(float64)
	userID := int(userIDFloat64)

	// Panggil service untuk mengubah password
	loginResponse, err := controller.userService.ChangePassword(request.Context(), userID, changePasswordRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	// Kirim response sukses
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Password berhasil diubah",
		Data:   loginResponse,
	})
}

func (controller *UserControllerImpl) ResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil user ID dari parameter
	userID, err := strconv.Atoi(params.ByName("userId"))
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "ID pengguna tidak valid",
			Data:   nil,
		})
		return
	}

	// Ambil role dan kodeOPD dari context
	role := request.Context().Value("roles").(string)
	kodeOPD := request.Context().Value("kode_opd").(string)

	// Panggil service untuk mendapatkan user
	user, err := controller.userService.FindByID(request.Context(), userID, "")
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Pengguna tidak ditemukan",
			Data:   nil,
		})
		return
	}

	// Pengecekan role
	if role != "admin_kota" && (role != "admin_opd" || user.KodeOPD != kodeOPD) {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusForbidden,
			Status: "Tidak memiliki izin untuk mereset password",
			Data:   nil,
		})
		return
	}

	// Panggil service untuk reset password
	resetResponse, err := controller.userService.ResetPassword(request.Context(), userID)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	// Kirim response sukses
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Password berhasil direset",
		Data:   resetResponse,
	})
}
