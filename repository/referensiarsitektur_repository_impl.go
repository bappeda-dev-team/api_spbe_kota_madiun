package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ReferensiArsitekturRepositoryImpl struct {
}

func NewReferensiKodeRepository() *ReferensiArsitekturRepositoryImpl {
	return &ReferensiArsitekturRepositoryImpl{}
}

func (repository *ReferensiArsitekturRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur)domain.ReferensiArsitektur {
	currentTime := time.Now()
	referensiarsitektur.Created_at = currentTime
	referensiarsitektur.Updated_at = currentTime

	script := "insert into referensi_arsitekturs(kode_referensi,nama_referensi,level_referensi,jenis_referensi,created_at,updated_at) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, referensiarsitektur.Kode_referensi, referensiarsitektur.Nama_referensi, referensiarsitektur.Level_referensi, referensiarsitektur.Jenis_referensi, referensiarsitektur.Created_at, referensiarsitektur.Updated_at)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	referensiarsitektur.IdReferensi = int(id)
	return referensiarsitektur
}

func (repository *ReferensiArsitekturRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur)domain.ReferensiArsitektur{
	currentTime := time.Now()
	referensiarsitektur.Updated_at = currentTime

	script := "update referensi_arsitekturs SET kode_referensi = ?, nama_referensi = ?, level_referensi = ?, jenis_referensi = ?, created_at = ?, updated_at = ? where id_referensi = ?"
	_, err := tx.ExecContext(ctx, script, referensiarsitektur.Kode_referensi, referensiarsitektur.Nama_referensi, referensiarsitektur.Level_referensi, referensiarsitektur.Jenis_referensi, referensiarsitektur.Created_at,referensiarsitektur.Updated_at, referensiarsitektur.IdReferensi)
	helper.PanicIfError(err)

	return referensiarsitektur
}

func (repository *ReferensiArsitekturRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur){
	script := "delete from referensi_arsitekturs where id =?"
	_, err := tx.ExecContext(ctx, script, referensiarsitektur.IdReferensi)
	helper.PanicIfError(err)

}

func(repository *ReferensiArsitekturRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ReferensiArsitektur{
	panic("find all")
}

func(repository *ReferensiArsitekturRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, referensiarsitekturId int) (domain.ReferensiArsitektur, error){
	script := "select id_referensi, kode_referensi, nama_referensi, level_referensi, jenis_referensi, created_at, updated_at from referensi_arsitekturs where id_referensi = ?"
	rows, err := tx.QueryContext(ctx, script, referensiarsitekturId)
	helper.PanicIfError(err)
	defer rows.Close()

	referensiarsitektur := domain.ReferensiArsitektur{}
	if rows.Next() {
		err := rows.Scan(&referensiarsitektur.IdReferensi, &referensiarsitektur.Kode_referensi, &referensiarsitektur.Kode_referensi, &referensiarsitektur.Nama_referensi, &referensiarsitektur.Level_referensi, &referensiarsitektur.Jenis_referensi, &referensiarsitektur.Created_at, &referensiarsitektur.Updated_at)
		helper.PanicIfError(err)
		return referensiarsitektur, nil
	} else {
		return referensiarsitektur, errors.New("referensiarsitektur is not found")
	}
}

func (repository *ReferensiArsitekturRepositoryImpl) FindByKodeRef(ctx context.Context, tx *sql.Tx, kodeReferensi string) []domain.ReferensiArsitektur {
	var result []domain.ReferensiArsitektur

	// Pertama, kita perlu menentukan level dari kode referensi yang ingin dicari
	// Di sini kita akan menggunakan jumlah titik (.) sebagai indikator level
	level := countDots(kodeReferensi)

	// Selanjutnya, kita akan mencari data dengan kode referensi yang memiliki prefiks yang sama
	// dan level yang sama atau lebih rendah dari kode referensi yang ingin dicari
	query := "SELECT * FROM referensi_arsitekturs WHERE kode_referensi LIKE ?"
	rows, err := tx.QueryContext(ctx, query, kodeReferensi+"%")
	if err != nil {
		// Handle error
		return result
	}
	defer rows.Close()

	// Iterasi melalui setiap baris hasil query
	for rows.Next() {
		var referensi domain.ReferensiArsitektur
		err := rows.Scan(&referensi.IdReferensi, &referensi.Kode_referensi, &referensi.Nama_referensi, &referensi.Level_referensi, &referensi.Jenis_referensi, &referensi.Created_at, &referensi.Updated_at)
		if err != nil {
			// Handle error
			continue
		}
		// Jika level dari kode referensi yang ditemukan kurang dari atau sama dengan level yang ditentukan
		// maka tambahkan data tersebut ke dalam hasil
		if countDots(referensi.Kode_referensi) <= level {
			result = append(result, referensi)
		}
	}

	return result
}

// Fungsi ini digunakan untuk menghitung jumlah titik (.) dalam sebuah string
func countDots(s string) int {
	count := 0
	for _, c := range s {
		if c == '.' {
			count++
		}
	}
	return count
}