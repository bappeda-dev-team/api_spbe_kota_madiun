package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (domain.User, error)
	GetUserRoles(ctx context.Context, tx *sql.Tx, userID int) ([]domain.Role, error)
}
