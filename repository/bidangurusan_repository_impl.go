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

type BidangUrusanRepositoryImpl struct {
}

func NewBidangUrusanRepositoryImpl() *BidangUrusanRepositoryImpl {
	return &BidangUrusanRepositoryImpl{}
}

func (repository *BidangUrusanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bidangurusanId int) (domain.BidangUrusan, error) {
	script := "select  id, kode_bidang_urusan, bidang_urusan, created_at, updated_at  from bidang_urusan where id = ?"
	rows, err := tx.QueryContext(ctx, script, bidangurusanId)
	if err != nil {
		log.Printf("Error executing query: %v", err) // Log the error
		helper.PanicIfError(err)
	}
	defer rows.Close()

	bidangUrusan := domain.BidangUrusan{}
	if rows.Next() {
		err := rows.Scan(&bidangUrusan.ID, &bidangUrusan.KodeBidangUrusan, &bidangUrusan.BidangUrusan, &bidangUrusan.CreatedAt, &bidangUrusan.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			helper.PanicIfError(err)
		}
		log.Printf("Successfully retrieved Bidang urusan: %+v", bidangUrusan)
		return bidangUrusan, nil
	} else {
		return bidangUrusan, errors.New("bidang urusan is not found")
	}
}

func (repository *BidangUrusanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.BidangUrusan {
	script := "select id, kode_bidang_urusan, bidang_urusan, created_at, updated_at from bidang_urusan"

	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var bidang []domain.BidangUrusan
	for rows.Next() {
		bidangUrusan := domain.BidangUrusan{}
		err := rows.Scan(&bidangUrusan.ID, &bidangUrusan.KodeBidangUrusan, &bidangUrusan.BidangUrusan, &bidangUrusan.CreatedAt, &bidangUrusan.UpdatedAt)
		helper.PanicIfError(err)
		bidang = append(bidang, bidangUrusan)
	}
	return bidang
}

func (repository *BidangUrusanRepositoryImpl) FetchBidangUrusan(ctx context.Context, tx *sql.Tx) (web.BidangUrusanOPD, error) {
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/opd/urusan_opd"
	method := "POST"

	formData := url.Values{}
	// formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return web.BidangUrusanOPD{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return web.BidangUrusanOPD{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return web.BidangUrusanOPD{}, err
	}

	log.Println("Received data:", string(body))

	var result web.ResponsesAPI
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.BidangUrusanOPD{}, err
	}

	log.Printf("Parsed Data: %+v\n", result.Results)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO bidang_urusan (kode_bidang_urusan, bidang_urusan)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE kode_bidang_urusan=VALUES(kode_bidang_urusan), bidang_urusan=VALUES(bidang_urusan)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return web.BidangUrusanOPD{}, err
	}
	defer stmt.Close()

	stmtBidangUrusanUrusan, err := tx.PrepareContext(ctx, `
		INSERT INTO urusan_bidang_opd (kode_opd, kode_urusan, bidang_urusan)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE kode_opd=VALUES(kode_opd), kode_urusan=VALUES(kode_urusan),bidang_urusan=VALUES(bidang_urusan)
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return web.BidangUrusanOPD{}, err
	}
	defer stmtBidangUrusanUrusan.Close()

	for _, item := range result.Results {
		for _, urusan := range item.UrusanOPD {
			for _, bidang := range urusan.BidangUrusanOPD {
				log.Printf("Inserting BidangUrusan: KodeBidangUrusan=%v, BidangUrusan=%v\n", bidang.KodeBidangUrusan, bidang.BidangUrusan)
				_, err := stmt.ExecContext(ctx, bidang.KodeBidangUrusan, bidang.BidangUrusan)
				if err != nil {
					log.Println("Error executing statement for Bidang Urusan:", err)
					return web.BidangUrusanOPD{}, nil
				}

				log.Printf("Inserting into urusan_bidang_opd: KodeOpd=%v, KodeUrusan=%v, BidangUrusan=%v\n", item.KodeOpd, urusan.KodeUrusan, bidang.BidangUrusan)
				_, err2 := stmtBidangUrusanUrusan.ExecContext(ctx, item.KodeOpd, urusan.KodeUrusan, bidang.BidangUrusan)
				if err2 != nil {
					log.Println("Error executing statement for urusan_bidang_opd:", err2)
					return web.BidangUrusanOPD{}, nil
				}
			}
		}
	}

	log.Println("Data successfully fetched and saved.")
	return web.BidangUrusanOPD{}, nil
}

func (repository *BidangUrusanRepositoryImpl) FindByBidangUrusan(ctx context.Context, tx *sql.Tx, bidangUrusan string) (domain.BidangUrusan, error) {
	script := "select  id, kode_bidang_urusan, bidang_urusan, created_at, updated_at  from bidang_urusan where bidang_urusan = ?"
	rows, err := tx.QueryContext(ctx, script, bidangUrusan)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		helper.PanicIfError(err)
	}
	defer rows.Close()

	bidang := domain.BidangUrusan{}
	if rows.Next() {
		err := rows.Scan(&bidang.ID, &bidang.KodeBidangUrusan, &bidang.BidangUrusan, &bidang.CreatedAt, &bidang.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			helper.PanicIfError(err)
		}
		log.Printf("Successfully retrieved Bidang urusan: %+v", bidang)
		return bidang, nil
	} else {
		return bidang, errors.New("bidang urusan is not found")
	}
}
