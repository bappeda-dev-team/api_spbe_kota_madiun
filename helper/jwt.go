package helper

import (
	"api_spbe_kota_madiun/model/domain"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWTToken(user domain.User) (string, error) {
	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = role.Nama
	}

	// Pastikan roleNames tidak kosong
	if len(roleNames) == 0 {
		roleNames = append(roleNames, "default_role")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"nip":      user.NIP,
		"nama":     user.Nama,
		"kode_opd": user.KodeOPD,
		"roles":    roleNames,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func RequestKodeOpd(request *http.Request) (role string, kodeOPD string, tahun int, err error) {
	// Ambil role dari middleware
	role, ok := request.Context().Value("roles").(string)
	if !ok {
		return "", "", 0, errors.New("gagal mendapatkan roles dari JWT")
	}

	// Ambil kode OPD dari token
	kodeOPDFromToken, ok := request.Context().Value("kode_opd").(string)
	if !ok {
		return "", "", 0, errors.New("kode OPD tidak ditemukan dalam context")
	}

	// Tentukan kode OPD berdasarkan role
	if role == "admin_kota" {
		kodeOPD = request.URL.Query().Get("kode_opd")
	} else {
		kodeOPD = kodeOPDFromToken
	}

	// Ambil tahun dari query parameter
	tahunStr := request.URL.Query().Get("tahun")
	if tahunStr != "" {
		tahun, err = strconv.Atoi(tahunStr)
		if err != nil {
			return "", "", 0, errors.New("tahun harus berupa angka")
		}
	}

	return role, kodeOPD, tahun, nil
}

func CheckRoleAndOPD(request *http.Request, allowedRole string) (string, string, error) {
	role, ok := request.Context().Value("roles").(string)
	if !ok {
		return "", "", errors.New("gagal mendapatkan peran pengguna")
	}

	kodeOPD, ok := request.Context().Value("kode_opd").(string)
	if !ok {
		return "", "", errors.New("gagal mendapatkan kode OPD")
	}

	if role != allowedRole {
		return "", "", fmt.Errorf("hanya pengguna dengan peran %s yang diizinkan", allowedRole)
	}

	return role, kodeOPD, nil
}
