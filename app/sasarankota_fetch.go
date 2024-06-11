package app

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
	"sync"
)

func FetchSaranKota(ctx context.Context, db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/skp/sasaran_kota"
	method := "POST"

	formData := url.Values{}
	formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	log.Println("Received data:", string(body))

	var result web.APIResponseDataSasaranKota
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	log.Printf("Parsed Data: %+v\n", result.Data)

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO sasaran_kota ( sasaran, strategi_kota, tujuan_kota, tahun)
		VALUES (?, ?, ?, 2024)
		ON DUPLICATE KEY UPDATE
		sasaran=VALUES(sasaran),strategi_kota=VALUES(strategi_kota), tujuan_kota=VALUES(tujuan_kota)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	for _, item := range result.Data.SasaranKota {
		log.Printf("Insert Sasaran Kota: Sasaran=%v, Strategi Kota=%v, Tujuan Kota=%v \n", item.Sasaran, item.StrategiKota, item.TujuanKota)
		_, err := stmt.ExecContext(ctx, item.Sasaran, item.StrategiKota, item.TujuanKota)
		if err != nil {
			log.Println("Error executing statement:", err)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return
	}

	log.Println("Data successfully fetched and saved.")
}
