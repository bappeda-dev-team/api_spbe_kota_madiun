package repository

import (
	"api_spbe_kota_madiun/model/web"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type UrusanRepositoryImpl struct {
}

func NewUrusanRepositoryImpl() *UrusanRepositoryImpl {
	return &UrusanRepositoryImpl{}
}

func (repository *UrusanRepositoryImpl) FetchUrusan(ctx context.Context, tx *sql.Tx) (web.UrusanOPD, error) {
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/opd/urusan_opd"
	method := "POST"

	formData := url.Values{}

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return web.UrusanOPD{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return web.UrusanOPD{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return web.UrusanOPD{}, err
	}

	log.Println("Received data:", string(body))

	var result web.ResponsesAPI
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.UrusanOPD{}, err
	}

	log.Printf("Parsed Data: %+v\n", result.Results)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO urusan (kode_urusan, urusan)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE urusan=VALUES(urusan)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return web.UrusanOPD{}, err
	}
	defer stmt.Close()

	for _, item := range result.Results {
		for _, urusan := range item.UrusanOPD {
			log.Printf("Inserting Urusan: KodeUrusan=%v, Urusan=%v\n", urusan.KodeUrusan, urusan.Urusan)
			_, err := stmt.ExecContext(ctx, urusan.KodeUrusan, urusan.Urusan)
			if err != nil {
				log.Println("Error executing statement for Urusan:", err)
				return web.UrusanOPD{}, nil
			}
		}
	}

	log.Println("Data successfully fetched and saved.")
	return web.UrusanOPD{}, nil
}
