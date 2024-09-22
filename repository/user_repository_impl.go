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

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO users (nip, nama, kode_opd, jabatan, password)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		nip=VALUES(nip), nama=VALUES(nama), kode_opd=VALUES(kode_opd), jabatan=VALUES(jabatan), password=VALUES(password)`)
	if err != nil {
		log.Println("Error menyiapkan statement:", err)
		return web.UserApiData{}, err
	}
	defer stmt.Close()

	defaultPassword := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error menghasilkan hash password:", err)
		return web.UserApiData{}, err
	}

	for _, pegawai := range result.Data {
		log.Printf("Menyisipkan data pegawai: NIP=%v, Nama=%v, Kode OPD=%v, Jabatan=%v\n", pegawai.Nip, pegawai.Nama, kodeOPD, pegawai.Jabatan)
		_, err = stmt.ExecContext(ctx, pegawai.Nip, pegawai.Nama, kodeOPD, pegawai.Jabatan, string(hashedPassword))
		if err != nil {
			log.Println("Error mengeksekusi statement:", err)
			return web.UserApiData{}, err
		}
	}

	log.Println("Data berhasil diambil dan disimpan.")
	return result, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	script := "SELECT id, nip, nama, kode_opd, jabatan, password FROM users"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.NIP, &user.Nama, &user.KodeOPD, &user.Jabatan, &user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userID int) (domain.User, error) {
	SQL := "SELECT id, nip, nama, kode_opd, password, jabatan FROM users WHERE id = ?"
	row := tx.QueryRowContext(ctx, SQL, userID)

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
