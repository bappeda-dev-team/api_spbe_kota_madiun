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

type OpdRepositoryImpl struct {
}

func NewOpdRepository() OpdRepository {
	return &OpdRepositoryImpl{}
}

func (rrepository *OpdRepositoryImpl) FetchKodeOpd(ctx context.Context, tx *sql.Tx) (web.Opd, error) {
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/opd/urusan_opd"
	method := "POST"

	formData := url.Values{}
	// formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return web.Opd{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return web.Opd{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return web.Opd{}, err
	}

	log.Println("Received data:", string(body))

	var result web.ResponsesAPI
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.Opd{}, err
	}

	log.Printf("Parsed Data: %+v\n", result.Results)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO opd (kode_opd, nama_opd)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE nama_opd=VALUES(nama_opd)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return web.Opd{}, err
	}
	defer stmt.Close()

	for _, item := range result.Results {
		log.Printf("Inserting OPD: KodeOpd=%v, NamaOpd=%v\n", item.KodeOpd, item.NamaOpd)
		_, err := stmt.ExecContext(ctx, item.KodeOpd, item.NamaOpd)
		if err != nil {
			log.Println("Error executing statement:", err)
			return web.Opd{}, err
		}
	}

	log.Println("Data successfully fetched and saved.")
	return web.Opd{}, nil
}
