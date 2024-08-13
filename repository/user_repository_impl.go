package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (domain.User, error) {
	query := "SELECT id, nip, nama, kode_opd, password FROM users WHERE nip = ?"
	row := tx.QueryRowContext(ctx, query, nip)

	var user domain.User
	err := row.Scan(&user.ID, &user.NIP, &user.Nama, &user.KodeOPD, &user.Password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) GetUserRoles(ctx context.Context, tx *sql.Tx, userID int) ([]domain.Role, error) {
	query := `
        SELECT r.id, r.nama
        FROM role r
        JOIN users_roles ur ON r.id = ur.role_id
        WHERE ur.user_id = ?
    `
	rows, err := tx.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		var role domain.Role
		err := rows.Scan(&role.ID, &role.Nama)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
