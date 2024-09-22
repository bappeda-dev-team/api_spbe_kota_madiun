package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
)

type UserRepository interface {
	FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (domain.User, error)
	GetUserRoles(ctx context.Context, tx *sql.Tx, userID int) ([]domain.Role, error)
	InsertApi(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun string) (web.UserApiData, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByID(ctx context.Context, tx *sql.Tx, userID int) (domain.User, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, userID int, newPassword string) error
}
