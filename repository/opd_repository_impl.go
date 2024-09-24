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

type OpdRepositoryImpl struct {
}

func NewOpdRepositoryImpl() *OpdRepositoryImpl {
	return &OpdRepositoryImpl{}
}

func (repository *OpdRepositoryImpl) FetchAllData(ctx context.Context, tx *sql.Tx) ([]web.Opd, error) {
	log.Println("Memulai FetchAllData")
	apiURL := "https://kak.madiunkota.go.id/api/opd/urusan_opd"
	method := "POST"

	formData := url.Values{}

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error membuat permintaan:", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error melakukan permintaan:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error membaca body respons:", err)
		return nil, err
	}

	log.Println("Data diterima:", string(body))

	var result web.ResponsesAPI
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshal JSON:", err)
		return nil, err
	}

	log.Printf("Data yang diurai: %+v\n", result.Results)

	if len(result.Results) == 0 {
		return nil, errors.New("tidak ada data OPD yang diterima")
	}

	// Persiapkan statements
	stmtOpd, err := tx.PrepareContext(ctx, `INSERT INTO opd (kode_opd, nama_opd) VALUES (?, ?) ON DUPLICATE KEY UPDATE nama_opd=VALUES(nama_opd)`)
	if err != nil {
		return nil, err
	}
	defer stmtOpd.Close()

	stmtUrusan, err := tx.PrepareContext(ctx, `INSERT INTO urusan (kode_urusan, urusan) VALUES (?, ?) ON DUPLICATE KEY UPDATE urusan=VALUES(urusan)`)
	if err != nil {
		return nil, err
	}
	defer stmtUrusan.Close()

	stmtBidangUrusan, err := tx.PrepareContext(ctx, `INSERT INTO bidang_urusan (kode_bidang_urusan, bidang_urusan) VALUES (?, ?) ON DUPLICATE KEY UPDATE bidang_urusan=VALUES(bidang_urusan)`)
	if err != nil {
		return nil, err
	}
	defer stmtBidangUrusan.Close()

	stmtUrusanBidangOpd, err := tx.PrepareContext(ctx, `INSERT INTO urusan_bidang_opd (kode_opd, kode_urusan, bidang_urusan) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE kode_opd=VALUES(kode_opd), kode_urusan=VALUES(kode_urusan), bidang_urusan=VALUES(bidang_urusan)`)
	if err != nil {
		return nil, err
	}
	defer stmtUrusanBidangOpd.Close()

	for _, opd := range result.Results {
		// Masukkan data OPD
		_, err = stmtOpd.ExecContext(ctx, opd.KodeOpd, opd.NamaOpd)
		if err != nil {
			return nil, err
		}

		for _, urusan := range opd.UrusanOPD {
			// Masukkan data urusan
			_, err := stmtUrusan.ExecContext(ctx, urusan.KodeUrusan, urusan.Urusan)
			if err != nil {
				return nil, err
			}

			for _, bidangUrusan := range urusan.BidangUrusanOPD {
				// Masukkan data bidang urusan
				_, err := stmtBidangUrusan.ExecContext(ctx, bidangUrusan.KodeBidangUrusan, bidangUrusan.BidangUrusan)
				if err != nil {
					return nil, err
				}

				// Masukkan data ke urusan_bidang_opd
				_, err = stmtUrusanBidangOpd.ExecContext(ctx, opd.KodeOpd, urusan.KodeUrusan, bidangUrusan.BidangUrusan)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	log.Println("Semua data berhasil diambil dan disimpan.")
	return result.Results, nil
}

func (repository *OpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOPD string) []domain.Opd {
	script := "select id, kode_opd, nama_opd from opd where 1=1"
	args := []interface{}{}

	if kodeOPD != "" {
		script += " AND kode_opd = ?"
		args = append(args, kodeOPD)
	}

	rows, err := tx.QueryContext(ctx, script, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var opd []domain.Opd
	for rows.Next() {
		getOpd := domain.Opd{}
		err := rows.Scan(&getOpd.Id, &getOpd.KodeOpd, &getOpd.NamaOpd)
		helper.PanicIfError(err)
		opd = append(opd, getOpd)
	}
	return opd
}

func (repository *OpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kodeOpd string) (domain.Opd, error) {
	script := "SELECT id, kode_opd, nama_opd FROM opd WHERE kode_opd = ?"
	rows, err := tx.QueryContext(ctx, script, kodeOpd)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	opd := domain.Opd{}
	if rows.Next() {
		err := rows.Scan(&opd.Id, &opd.KodeOpd, &opd.NamaOpd)
		if err != nil {
			panic(err)
		}
		return opd, nil
	} else {
		return opd, errors.New("opd tidak ditemukan")
	}

}

func (repository *OpdRepositoryImpl) FindKodeOpdUrusan(ctx context.Context, tx *sql.Tx, kodeOPD string) ([]domain.OpdUrusanBidang, error) {
	script := "select id, kode_opd, kode_urusan, bidang_urusan from urusan_bidang_opd where 1=1"
	args := []interface{}{}

	if kodeOPD != "" {
		script += " AND kode_opd = ?"
		args = append(args, kodeOPD)
	}

	rows, err := tx.QueryContext(ctx, script, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var opd []domain.OpdUrusanBidang
	for rows.Next() {
		getOpd := domain.OpdUrusanBidang{}
		err := rows.Scan(&getOpd.ID, &getOpd.KodeOpd, &getOpd.KodeUrusan, &getOpd.BidangUrusan)
		helper.PanicIfError(err)
		opd = append(opd, getOpd)
	}
	return opd, nil
}
