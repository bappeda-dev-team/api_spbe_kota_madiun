package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type PohonKinerjaRepositoryImpl struct {
}

func NewPohonKinerjaRepository() PohonKinerjaRepository {
	return &PohonKinerjaRepositoryImpl{}
}

func (repository *PohonKinerjaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pohonkinerjaId int) (domain.PohonKinerja, error) {
	script := "select id, nama_pohon, jenis_pohon, level_pohon, created_at, updated_at, tahun, kode_opd from pohon_kinerja where id = ?"
	rows, err := tx.QueryContext(ctx, script, pohonkinerjaId)
	if err != nil {
		log.Printf("Error executing query: %v", err) // Log the error
		helper.PanicIfError(err)
	}
	defer rows.Close()

	pohonKinerja := domain.PohonKinerja{}
	if rows.Next() {
		err := rows.Scan(&pohonKinerja.ID, &pohonKinerja.NamaPohon, &pohonKinerja.JenisPohon, &pohonKinerja.LevelPohon, &pohonKinerja.CreatedAt, &pohonKinerja.UpdatedAt, &pohonKinerja.Tahun, &pohonKinerja.KodeOpd)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			helper.PanicIfError(err)
		}
		log.Printf("Successfully retrieved Pohon Kinerja: %+v", pohonKinerja)
		return pohonKinerja, nil
	} else {
		return pohonKinerja, errors.New("Pohon Kinerja is not found")
	}
}

func (repository *PohonKinerjaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.PohonKinerja {
	script := "select id, nama_pohon, jenis_pohon, level_pohon, created_at, updated_at, tahun, kode_opd from pohon_kinerja"

	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var pohon []domain.PohonKinerja
	for rows.Next() {
		pohonKinerja := domain.PohonKinerja{}
		err := rows.Scan(&pohonKinerja.ID, &pohonKinerja.NamaPohon, &pohonKinerja.JenisPohon, &pohonKinerja.LevelPohon, &pohonKinerja.CreatedAt, &pohonKinerja.UpdatedAt, &pohonKinerja.Tahun, &pohonKinerja.KodeOpd)
		helper.PanicIfError(err)
		pohon = append(pohon, pohonKinerja)
	}
	return pohon
}

func (repository *PohonKinerjaRepositoryImpl)InsertApi(ctx context.Context, tx *sql.Tx) (domain.PohonKinerja, error){
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/pohon_kinerja/pohon_kinerja_opd.json"
	method := "POST"

	formData := url.Values{}
	formData.Set("kode_op", "5.01.5.05.0.00.02.0000")
	formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return domain.PohonKinerja{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return domain.PohonKinerja{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return domain.PohonKinerja{}, err
	}

	log.Println("Received data:", string(body))

	var result domain.APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return domain.PohonKinerja{}, err
	}

	log.Printf("Parsed Data: %+v\n", result.Results.Data)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO pohon_kinerja (jenis_pohon, level_pohon, kode_opd, nama_pohon, tahun)
		VALUES (?, ?, ?, ?, 2024)
		ON DUPLICATE KEY UPDATE
		jenis_pohon=VALUES(jenis_pohon), level_pohon=VALUES(level_pohon), kode_opd=VALUES(kode_opd), nama_pohon=VALUES(nama_pohon)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return domain.PohonKinerja{}, err
	}
	defer stmt.Close()

	for _, item := range result.Results.Data.PohonKinerjas {
		log.Printf("Insert Pohon Kinerja: Jenis Pohon=%v, Level Pohon=%v, Kode OPD=%v, Nama Pohon=%v \n", item.JenisPohon, item.LevelPohon, result.Results.Data.KodeOpd, item.NamaPohon)
		_, err := stmt.ExecContext(ctx, item.JenisPohon, item.LevelPohon, result.Results.Data.KodeOpd, item.NamaPohon)
		if err != nil {
			log.Println("Error executing statement:", err)
			return domain.PohonKinerja{}, err
		}
	}

	log.Println("Data successfully fetched and saved.")
	return domain.PohonKinerja{}, nil
}