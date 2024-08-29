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

type PohonKinerjaRepositoryImpl struct {
}

func NewPohonKinerjaRepositoryImpl() *PohonKinerjaRepositoryImpl {
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
		return pohonKinerja, errors.New("pohon Kinerja is not found")
	}
}

func (repository *PohonKinerjaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) []domain.PohonKinerja {
	script := "SELECT id, nama_pohon, jenis_pohon, level_pohon, created_at, updated_at, tahun, kode_opd FROM pohon_kinerja WHERE 1=1"
	args := []interface{}{}

	if tahun > 0 {
		script += " AND tahun = ?"
		args = append(args, tahun)
	}

	if kodeOpd != "" {
		script += " AND kode_opd = ?"
		args = append(args, kodeOpd)
	}

	rows, err := tx.QueryContext(ctx, script, args...)
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

func (repository *PohonKinerjaRepositoryImpl) InsertApi(ctx context.Context, tx *sql.Tx) (web.PohonKinerjaApi, error) {
	log.Println("Starting FetchKodeOpd")
	apiURL := "https://kak.madiunkota.go.id/api/pohon_kinerja/pohon_kinerja_opd.json"
	method := "POST"

	formData := url.Values{}
	formData.Set("kode_opd", "5.01.5.05.0.00.02.0000")
	formData.Set("tahun", "2024")

	client := &http.Client{}
	req, err := http.NewRequest(method, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
		return web.PohonKinerjaApi{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return web.PohonKinerjaApi{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return web.PohonKinerjaApi{}, err
	}

	log.Println("Received data:", string(body))
	var result web.APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return web.PohonKinerjaApi{}, err
	}

	log.Printf("Parsed Data: %+v\n", result.Results.Data)

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO pohon_kinerja (id, jenis_pohon, parent ,level_pohon, kode_opd, nama_pohon, tahun)
		VALUES (?,?, ?, ?, ?, ?, 2024)
		ON DUPLICATE KEY UPDATE
		id=(VALUES(id)),jenis_pohon=VALUES(jenis_pohon), parent=VALUES(parent), level_pohon=VALUES(level_pohon),kode_opd=VALUES(kode_opd), nama_pohon=VALUES(nama_pohon)`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return web.PohonKinerjaApi{}, err
	}
	defer stmt.Close()

	for _, item := range result.Results.Data.PohonKinerjas {
		log.Printf("Insert Pohon Kinerja:Id-%v, Jenis Pohon=%v, Parent=%v, Level Pohon=%v, Kode OPD=%v, Nama Pohon=%v \n", item.ID, item.JenisPohon, item.Parent, item.LevelPohon, result.Results.Data.KodeOpd, item.Strategi)
		_, err := stmt.ExecContext(ctx, item.ID, item.JenisPohon, item.Parent, item.LevelPohon, result.Results.Data.KodeOpd, item.Strategi)
		if err != nil {
			log.Println("Error executing statement:", err)
			return web.PohonKinerjaApi{}, err
		}
	}

	log.Println("Data successfully fetched and saved.")
	return web.PohonKinerjaApi{}, nil
}

func (repository *PohonKinerjaRepositoryImpl) FindByOperational(ctx context.Context, tx *sql.Tx, pohonkinerjaId int) (domain.PohonKinerja, []domain.PohonKinerja, []domain.PohonKinerja, error) {
	// Ambil data operational
	operationalScript := "SELECT id, nama_pohon, jenis_pohon, level_pohon, parent, created_at, updated_at, tahun, kode_opd FROM pohon_kinerja WHERE id = ?"
	operationalRow := tx.QueryRowContext(ctx, operationalScript, pohonkinerjaId)

	var operational domain.PohonKinerja
	err := operationalRow.Scan(&operational.ID, &operational.NamaPohon, &operational.JenisPohon, &operational.LevelPohon, &operational.Parent, &operational.CreatedAt, &operational.UpdatedAt, &operational.Tahun, &operational.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.PohonKinerja{}, nil, nil, errors.New("pohon kinerja not found")
		}
		return domain.PohonKinerja{}, nil, nil, err
	}

	// Jika tidak ada parent, kembalikan hanya data operational
	if operational.Parent == "" {
		return domain.PohonKinerja{}, nil, []domain.PohonKinerja{operational}, nil
	}

	// Ambil data tactical
	tacticalScript := "SELECT id, nama_pohon, jenis_pohon, level_pohon, parent, created_at, updated_at, tahun, kode_opd FROM pohon_kinerja WHERE id = ?"
	tacticalRow := tx.QueryRowContext(ctx, tacticalScript, operational.Parent)

	var tactical domain.PohonKinerja
	err = tacticalRow.Scan(&tactical.ID, &tactical.NamaPohon, &tactical.JenisPohon, &tactical.LevelPohon, &tactical.Parent, &tactical.CreatedAt, &tactical.UpdatedAt, &tactical.Tahun, &tactical.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.PohonKinerja{}, nil, []domain.PohonKinerja{operational}, nil
		}
		return domain.PohonKinerja{}, nil, nil, err
	}

	// Jika tidak ada parent untuk tactical, kembalikan data tactical dan operational
	if tactical.Parent == "" {
		return domain.PohonKinerja{}, []domain.PohonKinerja{tactical}, []domain.PohonKinerja{operational}, nil
	}

	// Ambil data strategic
	strategicScript := "SELECT id, nama_pohon, jenis_pohon, level_pohon, parent, created_at, updated_at, tahun, kode_opd FROM pohon_kinerja WHERE id = ?"
	strategicRow := tx.QueryRowContext(ctx, strategicScript, tactical.Parent)

	var strategic domain.PohonKinerja
	err = strategicRow.Scan(&strategic.ID, &strategic.NamaPohon, &strategic.JenisPohon, &strategic.LevelPohon, &strategic.Parent, &strategic.CreatedAt, &strategic.UpdatedAt, &strategic.Tahun, &strategic.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.PohonKinerja{}, []domain.PohonKinerja{tactical}, []domain.PohonKinerja{operational}, nil
		}
		return domain.PohonKinerja{}, nil, nil, err
	}

	return strategic, []domain.PohonKinerja{tactical}, []domain.PohonKinerja{operational}, nil
}
