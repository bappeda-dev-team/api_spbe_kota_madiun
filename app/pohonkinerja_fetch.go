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

func FetchPohonKinerja(ctx context.Context, db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	apiURL := "https://kak.madiunkota.go.id/api/pohon_kinerja/pohon_kinerja_opd.json"
	method := "POST"

	formData := url.Values{}
	formData.Set("kode_opd", "5.01.5.05.0.00.02.0000")
	formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

	// Tambahkan log untuk memeriksa data yang diterima
	log.Println("Received data:", string(body))

	// Menguraikan respons JSON
	var result web.APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	// Tambahkan log untuk memeriksa data yang diurai
	log.Printf("Parsed Data: %+v\n", result.Results.Data)

	// Menggunakan data langsung dalam eksekusi pernyataan SQL
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO pohon_kinerja (jenis_pohon, level_pohon, kode_opd, nama_pohon, tahun)
		VALUES (?, ?, ?, ?, 2024)
		ON DUPLICATE KEY UPDATE
		jenis_pohon=VALUES(jenis_pohon), level_pohon=VALUES(level_pohon), kode_opd=VALUES(kode_opd), nama_pohon=VALUES(nama_pohon)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	for _, item := range result.Results.Data.PohonKinerjas {
		log.Printf("Insert Pohon Kinerja: Jenis Pohon=%v, Level Pohon=%v, Kode OPD=%v, Nama Pohon=%v \n", item.JenisPohon, item.LevelPohon, result.Results.Data.KodeOpd, item.Strategi)
		_, err := stmt.ExecContext(ctx, item.JenisPohon, item.LevelPohon, result.Results.Data.KodeOpd, item.Strategi)
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
