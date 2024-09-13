package middleware

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"context"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/v1/login" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	tokenString := request.Header.Get("Authorization")
	if tokenString == "" {
		middleware.sendUnauthorizedResponse(writer, "Token tidak ditemukan")
		return
	}

	claims, err := helper.ValidateJWTToken(tokenString)
	if err != nil {
		middleware.sendUnauthorizedResponse(writer, "Token tidak valid")
		return
	}

	roles, ok := claims["roles"].([]interface{})
	if !ok || len(roles) == 0 {
		middleware.sendUnauthorizedResponse(writer, "Role tidak valid")
		return
	}

	kodeOPD, ok := claims["kode_opd"].(string)
	if !ok {
		middleware.sendUnauthorizedResponse(writer, "Kode OPD tidak valid")
		return
	}

	// Fungsi untuk menentukan prioritas peran
	getPriorityRole := func(roles []interface{}) string {
		priorityOrder := map[string]int{
			"admin_kota": 3,
			"admin_opd":  2,
			"asn":        1,
		}

		highestPriority := -1
		priorityRole := ""

		for _, r := range roles {
			role := r.(string)
			if priority, exists := priorityOrder[role]; exists && priority > highestPriority {
				highestPriority = priority
				priorityRole = role
			}
		}

		return priorityRole
	}

	role := getPriorityRole(roles)

	// Menyimpan role dan kode OPD ke dalam context
	ctx := context.WithValue(request.Context(), "roles", role)
	ctx = context.WithValue(ctx, "kode_opd", kodeOPD)

	publicEndpoints := []string{
		"/v1/referensiarsitektur",
		"/v1/referensiarsitektur/:kodeReferensi",
		"/v1/referensiarsitekturbyid/:referensiarsitekturId",
		"/v1/sasarankota",
		"/v1/sasarankota/:sasarankotaId",
		"/v1/pohonkinerja/:pohonId",
		"/v1/bidangurusan",
		"/sasarankotafetch",
		"/opdafetch",
		"/urusanfetch",
		"/bidangurusanfetch",
	}

	// Periksa apakah endpoint saat ini adalah endpoint publik
	for _, endpoint := range publicEndpoints {
		if request.URL.Path == endpoint {
			middleware.Handler.ServeHTTP(writer, request)
			return
		}
	}

	// Jika metode adalah POST atau PUT, tambahkan kode OPD ke body request untuk admin_opd dan asn
	if (role == "admin_opd" || role == "asn") && (request.Method == "POST" || request.Method == "PUT") {
		err := request.ParseForm()
		if err != nil {
			middleware.sendUnauthorizedResponse(writer, "Gagal memproses form data")
			return
		}
		request.Form.Set("kode_opd", kodeOPD)
	}

	request = request.WithContext(ctx)

	// Mengatur akses berdasarkan role
	switch role {
	case "admin_kota":
		// Admin kota melakukan crud semua data
		if request.Method == "GET" || request.Method == "POST" || request.Method == "PUT" || request.Method == "DELETE" {
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			middleware.sendUnauthorizedResponse(writer, "Admin kota tidak ada akses")
		}
	case "admin_opd":
		// Admin OPD dapat melakukan crud data berdasarkan kode OPD
		if request.Method == "GET" || request.Method == "POST" || request.Method == "PUT" || request.Method == "DELETE" {
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			middleware.sendUnauthorizedResponse(writer, "Metode tidak diizinkan untuk admin kota")
		}
	case "asn":
		// ASN dapat melakukan CRUD pada data OPD mereka sendiri
		if request.Method == "GET" || request.Method == "POST" || request.Method == "PUT" || request.Method == "DELETE" {
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			middleware.sendUnauthorizedResponse(writer, "Metode tidak diizinkan untuk ASN")
		}
	default:
		middleware.sendUnauthorizedResponse(writer, "Role tidak dikenali")
	}
}

func (middleware *AuthMiddleware) sendUnauthorizedResponse(writer http.ResponseWriter, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   message,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
