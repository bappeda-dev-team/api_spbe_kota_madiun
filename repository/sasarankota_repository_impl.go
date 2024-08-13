package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
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

type SasaranKotaRepositoryImpl struct {
}

func NewSasaranKotaRepositoryImpl() *SasaranKotaRepositoryImpl {
	return &SasaranKotaRepositoryImpl{}
}

func (repository *SasaranKotaRepositoryImpl) InsertApi(ctx context.Context, tx *sql.Tx) (web.SasaranKotaApi, error) {
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/skp/sasaran_kota"
	method := "POST"

	formData := url.Values{}
	formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return web.SasaranKotaApi{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return web.SasaranKotaApi{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return web.SasaranKotaApi{}, err
	}

	log.Println("Received data:", string(body))

	var result web.APIResponseDataSasaranKota
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.SasaranKotaApi{}, err
	}

	log.Printf("Parsed Data: %+v\n", result.Data)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO sasaran_kota (id, sasaran, strategi_kota, tujuan_kota, tahun)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		sasaran=VALUES(sasaran), strategi_kota=VALUES(strategi_kota), tujuan_kota=VALUES(tujuan_kota)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return web.SasaranKotaApi{}, err
	}
	defer stmt.Close()

	for _, item := range result.Data.SasaranKota {
		log.Printf("Insert Sasaran Kota: ID=%v, Sasaran=%v, Strategi Kota=%v, Tujuan Kota=%v\n", item.ID, item.Sasaran, item.StrategiKota, item.TujuanKota)
		_, err := stmt.ExecContext(ctx, item.ID, item.Sasaran, item.StrategiKota, item.TujuanKota, result.Data.Tahun)
		if err != nil {
			log.Println("Error executing statement:", err)
			return web.SasaranKotaApi{}, err
		}
	}

	log.Println("Data successfully fetched and saved.")
	return web.SasaranKotaApi{}, nil
}

func (repository *SasaranKotaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, sasarankotaId int) (domain.SasaranKota, error) {
	script := "select id, sasaran, strategi_kota, tujuan_kota, tahun, created_at, updated_at from sasaran_kota where id = ?"
	rows, err := tx.QueryContext(ctx, script, sasarankotaId)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		helper.PanicIfError(err)
	}
	defer rows.Close()

	sasaran := domain.SasaranKota{}
	if rows.Next() {
		err := rows.Scan(&sasaran.ID, &sasaran.Sasaran, &sasaran.StrategiKota, &sasaran.TujuanKota, &sasaran.Tahun, &sasaran.CreatedAt, &sasaran.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			helper.PanicIfError(err)
		}
		log.Printf("Successfully retrieved Sasaran Kota: %+v", sasarankotaId)
		return sasaran, nil
	} else {
		return sasaran, errors.New("sasaran kota is not found")
	}
}

func (repository *SasaranKotaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.SasaranKota {
	script := "select id, sasaran, strategi_kota, tujuan_kota, tahun, created_at, updated_at from sasaran_kota"

	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var sasaranKota []domain.SasaranKota
	for rows.Next() {
		sasaran := domain.SasaranKota{}
		err := rows.Scan(&sasaran.ID, &sasaran.Sasaran, &sasaran.StrategiKota, &sasaran.TujuanKota, &sasaran.Tahun, &sasaran.CreatedAt, &sasaran.UpdatedAt)
		helper.PanicIfError(err)
		sasaranKota = append(sasaranKota, sasaran)
	}
	return sasaranKota
}
