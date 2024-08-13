package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type UserService interface {
	Login(ctx context.Context, req web.LoginRequest) (web.LoginResponse, error)
}
