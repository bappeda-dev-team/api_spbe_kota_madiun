package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type UserService interface {
	Login(ctx context.Context, req web.LoginRequest) (web.LoginResponse, error)
	InsertApi(ctx context.Context, kodeOPD string, tahun string) (web.UserApiData, error)
	FindAll(ctx context.Context, kodeOPD string, rolesID int) []web.UserResponse
	FindByNIP(ctx context.Context, nip string) (web.UserResponse, error)
	FindByID(ctx context.Context, userID int, kodeOPD string) (web.UserResponse, error)
	ChangePassword(ctx context.Context, userID int, request web.ChangePasswordRequest) (web.LoginResponse, error)
	ResetPassword(ctx context.Context, userID int) (web.LoginResponse, error)
}
