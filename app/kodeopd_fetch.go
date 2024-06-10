package app

import (
	"api_spbe_kota_madiun/model/web"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

func FetchKodeOpd(ctx context.Context, db *sql.DB, wg *sync.WaitGroup) {

	log.Println("Starting FetchKodeOpd")
	url := "https://kak.madiunkota.go.id/api/opd/urusan_opd" // Gantilah dengan URL yang benar
	method := "POST"

	payload := bytes.NewReader([]byte(""))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
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

	log.Println("Response Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: Response status is not OK:", resp.Status)
		return
	}
	log.Println("Response Headers:", resp.Header)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	body := buf.Bytes()
	log.Println("Received data:", string(body))

	var response web.ResponsesAPI
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	results := response.Results
	if len(results) == 0 {
		log.Println("No 'results' found in JSON")
		return
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return
	}
	defer tx.Rollback()

	// kode opd
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO opd (kode_opd, nama_opd)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE nama_opd=VALUES(nama_opd)
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	// urusan opd
	stmtUrusan, err := tx.PrepareContext(ctx, `
		INSERT INTO urusan (kode_urusan, urusan)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE urusan=VALUES(urusan)
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmtUrusan.Close()

	// bidang urusan opd
	stmtBidangUrusan, err := tx.PrepareContext(ctx, `
		INSERT INTO bidang_urusan (kode_bidang_urusan, bidang_urusan)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE kode_bidang_urusan=VALUES(kode_bidang_urusan), bidang_urusan=VALUES(bidang_urusan)
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmtBidangUrusan.Close()

	// gabungan
	stmtBidangUrusanUrusan, err := tx.PrepareContext(ctx, `
		INSERT INTO urusan_bidang_opd (kode_opd, kode_urusan, bidang_urusan)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE kode_opd=VALUES(kode_opd), kode_urusan=VALUES(kode_urusan),bidang_urusan=VALUES(bidang_urusan)
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmtBidangUrusanUrusan.Close()

	for _, item := range results {
		log.Printf("Inserting OPD: KodeOpd=%v, NamaOpd=%v\n", item.KodeOpd, item.NamaOpd)
		_, err := stmt.ExecContext(ctx, item.KodeOpd, item.NamaOpd)
		if err != nil {
			log.Println("Error executing statement for OPD:", err)
			return
		}
		for _, urusan := range item.UrusanOPD {
			log.Printf("Inserting Urusan: KodeUrusan=%v, Urusan=%v\n", urusan.KodeUrusan, urusan.Urusan)
			_, err := stmtUrusan.ExecContext(ctx, urusan.KodeUrusan, urusan.Urusan)
			if err != nil {
				log.Println("Error executing statement for Urusan:", err)
				return
			}
			for _, bidang := range urusan.BidangUrusanOPD {
				log.Printf("Inserting BidangUrusan: KodeBidangUrusan=%v, BidangUrusan=%v\n", bidang.KodeBidangUrusan, bidang.BidangUrusan)
				_, err := stmtBidangUrusan.ExecContext(ctx, bidang.KodeBidangUrusan, bidang.BidangUrusan)
				if err != nil {
					log.Println("Error executing statement for Bidang Urusan:", err)
					return
				}

				log.Printf("Inserting into urusan_bidang_opd: KodeOpd=%v, KodeUrusan=%v, BidangUrusan=%v\n", item.KodeOpd, urusan.KodeUrusan, bidang.BidangUrusan)
				_, err2 := stmtBidangUrusanUrusan.ExecContext(ctx, item.KodeOpd, urusan.KodeUrusan, bidang.BidangUrusan)
				if err2 != nil {
					log.Println("Error executing statement for urusan_bidang_opd:", err2)
					return
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return
	}

	log.Println("Data successfully fetched and saved.")

}
