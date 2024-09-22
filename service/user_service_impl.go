package service

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"api_spbe_kota_madiun/repository"
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, req web.LoginRequest) (web.LoginResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.LoginResponse{}, err
	}
	defer tx.Rollback()

	user, err := service.UserRepository.FindByNIP(ctx, tx, req.NIP)
	if err != nil {
		return web.LoginResponse{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return web.LoginResponse{}, errors.New("invalid credentials")
	}

	domainRoles, err := service.UserRepository.GetUserRoles(ctx, tx, user.ID)
	if err != nil {
		return web.LoginResponse{}, err
	}

	// Konversi []domain.Role menjadi []web.Role
	responseRoles := make([]web.Role, len(domainRoles))
	for i, role := range domainRoles {
		responseRoles[i] = web.Role{
			ID:   role.ID,
			Nama: role.Nama,
		}
	}

	// Pastikan user memiliki roles sebelum generate token
	if len(domainRoles) > 0 {
		user.Roles = domainRoles
	} else {
		user.Roles = []domain.Role{{Nama: "default_role"}}
	}

	token, err := helper.GenerateJWTToken(user)
	if err != nil {
		return web.LoginResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		return web.LoginResponse{}, err
	}

	return web.LoginResponse{
		Token: token,
		User: web.UserResponse{
			ID:      user.ID,
			NIP:     user.NIP,
			Nama:    user.Nama,
			KodeOPD: user.KodeOPD,
			Roles:   responseRoles,
		},
	}, nil
}

func (service *UserServiceImpl) InsertApi(ctx context.Context, kodeOPD string, tahun string) (web.UserApiData, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserApiData{}, err
	}
	defer tx.Rollback()

	result, err := service.UserRepository.InsertApi(ctx, tx, kodeOPD, tahun)
	if err != nil {
		return web.UserApiData{}, err
	}

	err = tx.Commit()
	if err != nil {
		return web.UserApiData{}, err
	}

	return result, nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)

}

func (service *UserServiceImpl) FindByNIP(ctx context.Context, nip string) (web.UserResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByNIP(ctx, tx, nip)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user), nil
}

func (service *UserServiceImpl) ChangePassword(ctx context.Context, userID int, request web.ChangePasswordRequest) (web.LoginResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.LoginResponse{}, err
	}
	defer tx.Rollback()

	user, err := service.UserRepository.FindByID(ctx, tx, userID)
	if err != nil {
		return web.LoginResponse{}, err
	}

	// Verifikasi old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		return web.LoginResponse{}, errors.New("password lama tidak sesuai")
	}

	// Enkripsi password baru
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return web.LoginResponse{}, err
	}

	// Update password di database
	err = service.UserRepository.UpdatePassword(ctx, tx, userID, string(hashedNewPassword))
	if err != nil {
		return web.LoginResponse{}, err
	}

	// Ambil roles pengguna
	roles, err := service.UserRepository.GetUserRoles(ctx, tx, user.ID)
	if err != nil {
		return web.LoginResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		return web.LoginResponse{}, err
	}

	// Konversi domain.Role ke web.Role
	webRoles := make([]web.Role, len(roles))
	for i, role := range roles {
		webRoles[i] = web.Role{
			ID:   role.ID,
			Nama: role.Nama,
		}
	}

	// Generate token JWT baru
	token, err := helper.GenerateJWTToken(user)
	if err != nil {
		return web.LoginResponse{}, err
	}

	return web.LoginResponse{
		Token: token,
		User: web.UserResponse{
			ID:      user.ID,
			NIP:     user.NIP,
			Nama:    user.Nama,
			KodeOPD: user.KodeOPD,
			Jabatan: user.Jabatan.String, // Gunakan .String untuk mengakses nilai string
			Roles:   webRoles,
		},
	}, nil
}
