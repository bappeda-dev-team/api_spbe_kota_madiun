package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"api_spbe_kota_madiun/model/web"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type SasaranKinerjaPegawaiRepositoryImpl struct {
}

func NewSasaranKinerjaPegawaiRepositoryImpl() *SasaranKinerjaPegawaiRepositoryImpl {
	return &SasaranKinerjaPegawaiRepositoryImpl{}
}

func (repository *SasaranKinerjaPegawaiRepositoryImpl) FetchsasaranKinerja(ctx context.Context, tx *sql.Tx, kodeOPD string, tahun string) (web.ApiResponsSasaranKinerjaPegawai, error) {
	apiURL := "https://kak.madiunkota.go.id/api/pohon_kinerja/list_sasaran_operational"
	method := "GET"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("kode_opd", kodeOPD)
	_ = writer.WriteField("tahun", tahun)
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		log.Println("Error membuat request:", err)
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error melakukan request:", err)
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error membaca response body:", err)
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}

	log.Println("Data diterima:", string(bodyBytes))
	var result web.ApiResponsSasaranKinerjaPegawai
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}

	log.Printf("Data yang diparse: %+v\n", result)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO sasaran_kinerja_pegawai (id, kode_opd, kode_sasaran, sasaran_kinerja, tahun_sasaran, anggaran_sasaran, pelaksana_sasaran, kode_subkegiatan_sasaran, subkegiatan_sasaran)
		VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		id=VALUES(id), kode_opd=VALUES(kode_opd), kode_sasaran=VALUES(kode_sasaran), sasaran_kinerja=VALUES(sasaran_kinerja), tahun_sasaran=VALUES(tahun_sasaran), anggaran_sasaran=VALUES(anggaran_sasaran), pelaksana_sasaran=VALUES(pelaksana_sasaran), subkegiatan_sasaran=VALUES(subkegiatan_sasaran)
	`)
	if err != nil {
		log.Println("Error menyiapkan statement:", err)
		return web.ApiResponsSasaranKinerjaPegawai{}, err
	}
	defer stmt.Close()

	for _, sasarankinerja := range result.Results {
		_, err = stmt.ExecContext(ctx, sasarankinerja.Id, sasarankinerja.KodeOPD, sasarankinerja.KodeSasaranPegawai, sasarankinerja.SasaranPegawai, sasarankinerja.TahunSasaran, sasarankinerja.AnggaranSasaran, sasarankinerja.PelaksanaSasaran, sasarankinerja.KodeSubKegiatan, sasarankinerja.SubKegiatan)
		if err != nil {
			log.Println("Error mengeksekusi statement:", err)
			return web.ApiResponsSasaranKinerjaPegawai{}, err
		}
	}
	log.Println("Data berhasil diambil dan disimpan.")
	return result, nil
}

func (repository *SasaranKinerjaPegawaiRepositoryImpl) FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun string) []domain.SasaranKinerja {
	script := "SELECT id, kode_opd, kode_sasaran, tahun_sasaran, sasaran_kinerja, anggaran_sasaran, pelaksana_sasaran, kode_subkegiatan_sasaran, subkegiatan_sasaran FROM sasaran_kinerja_pegawai WHERE 1=1"
	var args []interface{}

	if kodeOpd != "" { // Menggunakan string kosong sebagai indikator
		script += " AND kode_opd = ?"
		args = append(args, kodeOpd)
	}
	if tahun != "" { // Menggunakan string kosong sebagai indikator
		script += " AND tahun_sasaran = ?"
		args = append(args, tahun)
	}

	script += " ORDER BY id DESC"
	rows, err := tx.QueryContext(ctx, script, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var sasaran []domain.SasaranKinerja
	for rows.Next() {
		var sasaranKinerja domain.SasaranKinerja
		err := rows.Scan(&sasaranKinerja.Id, &sasaranKinerja.KodeOpd, &sasaranKinerja.KodeSasaran, &sasaranKinerja.Tahun, &sasaranKinerja.SasaranKinerjaPegawai, &sasaranKinerja.AnggaranSasaran, &sasaranKinerja.PelaksanaSasaran, &sasaranKinerja.KodeSubKegiatan, &sasaranKinerja.SubKegiatan)
		helper.PanicIfError(err)
		sasaran = append(sasaran, sasaranKinerja)
	}
	return sasaran
}

func (repository *SasaranKinerjaPegawaiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, sasarankinerjaId int) (domain.SasaranKinerja, error) {
	script := "SELECT id, kode_opd, kode_sasaran, tahun_sasaran, sasaran_kinerja, anggaran_sasaran, pelaksana_sasaran, kode_subkegiatan_sasaran, subkegiatan_sasaran FROM sasaran_kinerja_pegawai WHERE id = ?"
	row := tx.QueryRowContext(ctx, script, sasarankinerjaId)

	sasaranKinerja := domain.SasaranKinerja{}

	err := row.Scan(&sasaranKinerja.Id, &sasaranKinerja.KodeOpd, &sasaranKinerja.KodeSasaran, &sasaranKinerja.Tahun, &sasaranKinerja.SasaranKinerjaPegawai, &sasaranKinerja.AnggaranSasaran, &sasaranKinerja.PelaksanaSasaran, &sasaranKinerja.KodeSubKegiatan, &sasaranKinerja.SubKegiatan)
	if err == sql.ErrNoRows {
		return sasaranKinerja, errors.New("Data sasaran kinerja pegawai dengan ID " + strconv.Itoa(sasarankinerjaId) + " tidak ditemukan")
	}
	if err != nil {
		return sasaranKinerja, err
	}

	return sasaranKinerja, nil
}
