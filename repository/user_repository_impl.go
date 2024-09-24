package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func (repository *UserRepositoryImpl) InsertApi(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun string) (web.UserApiData, error) {
	log.Println("Memulai FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/opd/daftar_pegawai"
	method := "POST"

	// Membuat form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("kode_opd", kodeOPD)
	_ = writer.WriteField("tahun", tahun)
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		log.Println("Error membuat request:", err)
		return web.UserApiData{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error melakukan request:", err)
		return web.UserApiData{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error membaca response body:", err)
		return web.UserApiData{}, err
	}

	log.Println("Data diterima:", string(bodyBytes))
	var result web.UserApiData
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.UserApiData{}, err
	}

	log.Printf("Data yang diparse: %+v\n", result)

	stmtUser, err := tx.PrepareContext(ctx, `
		INSERT INTO users (nip, nama, kode_opd, jabatan, password)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		nama=VALUES(nama), kode_opd=VALUES(kode_opd), jabatan=VALUES(jabatan), password=VALUES(password)`)
	if err != nil {
		log.Println("Error menyiapkan statement users:", err)
		return web.UserApiData{}, err
	}
	defer stmtUser.Close()

	stmtGetUserID, err := tx.PrepareContext(ctx, `SELECT id FROM users WHERE nip = ?`)
	if err != nil {
		log.Println("Error menyiapkan statement get user ID:", err)
		return web.UserApiData{}, err
	}
	defer stmtGetUserID.Close()

	stmtRole, err := tx.PrepareContext(ctx, `
		INSERT INTO users_roles (user_id, role_id)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE role_id=VALUES(role_id)`)
	if err != nil {
		log.Println("Error menyiapkan statement users_roles:", err)
		return web.UserApiData{}, err
	}
	defer stmtRole.Close()

	defaultPassword := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error menghasilkan hash password:", err)
		return web.UserApiData{}, err
	}

	for _, pegawai := range result.Data {
		log.Printf("Menyisipkan data pegawai: NIP=%v, Nama=%v, Kode OPD=%v, Jabatan=%v\n", pegawai.Nip, pegawai.Nama, kodeOPD, pegawai.Jabatan)
		res, err := stmtUser.ExecContext(ctx, pegawai.Nip, pegawai.Nama, kodeOPD, pegawai.Jabatan, string(hashedPassword))
		if err != nil {
			log.Println("Error mengeksekusi statement users:", err)
			return web.UserApiData{}, err
		}

		var userID int
		if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
			// Jika tidak ada baris yang terpengaruh, artinya user sudah ada
			// Kita perlu mengambil ID user yang sudah ada
			err = stmtGetUserID.QueryRowContext(ctx, pegawai.Nip).Scan(&userID)
			if err != nil {
				log.Println("Error mendapatkan user ID:", err)
				return web.UserApiData{}, err
			}
		} else {
			// Jika ada baris yang terpengaruh, ambil ID yang baru saja di-insert
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				log.Println("Error mendapatkan last insert ID:", err)
				return web.UserApiData{}, err
			}
			userID = int(lastInsertID)
		}

		userRole := domain.UserRole{
			UserId: userID,
			RoleId: 3,
		}

		// Menyisipkan data ke users_roles
		_, err = stmtRole.ExecContext(ctx, userRole.UserId, userRole.RoleId)
		if err != nil {
			log.Println("Error mengeksekusi statement users_roles:", err)
			return web.UserApiData{}, err
		}
	}

	log.Println("Data berhasil diambil dan disimpan.")
	return result, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOPD string, rolesID int) []domain.User {
	script := `
		SELECT u.id, u.nip, u.nama, u.kode_opd, u.jabatan, u.password, r.id as role_id, r.nama as role_nama
		FROM users u
		JOIN users_roles ur ON u.id = ur.user_id
		JOIN role r ON ur.role_id = r.id
		WHERE 1=1
	`
	var args []interface{}

	if kodeOPD != "" {
		script += " AND u.kode_opd = ?"
		args = append(args, kodeOPD)
	}

	if rolesID != 0 {
		script += " AND ur.role_id = ?"
		args = append(args, rolesID)
	}

	rows, err := tx.QueryContext(ctx, script, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		var role domain.Role
		err := rows.Scan(&user.ID, &user.NIP, &user.Nama, &user.KodeOPD, &user.Jabatan, &user.Password, &role.ID, &role.Nama)
		helper.PanicIfError(err)
		user.Roles = append(user.Roles, role)
		users = append(users, user)
	}
	return users
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userID int, kodeOPD string) (domain.User, error) {
	SQL := "SELECT id, nip, nama, kode_opd, password, jabatan FROM users WHERE id = ?"
	args := []interface{}{userID}

	if kodeOPD != "" {
		SQL += " AND kode_opd = ?"
		args = append(args, kodeOPD)
	}

	row := tx.QueryRowContext(ctx, SQL, args...)

	var user domain.User
	err := row.Scan(&user.ID, &user.NIP, &user.Nama, &user.KodeOPD, &user.Password, &user.Jabatan)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, userID int, newPassword string) error {
	SQL := "UPDATE users SET password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, newPassword, userID)
	return err
}

func (repository *UserRepositoryImpl) ResetPassword(ctx context.Context, tx *sql.Tx, userID int) error {
	defaultPassword := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	SQL := "UPDATE users SET password = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, string(hashedPassword), userID)
	return err
}
