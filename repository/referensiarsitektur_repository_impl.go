package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"
)

type ReferensiArsitekturRepositoryImpl struct {
}

func NewReferensiArsitekturRepository() ReferensiArsitekturRepository {
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

	script := "update referensi_arsitekturs SET kode_referensi = ?, nama_referensi = ?, level_referensi = ?, jenis_referensi = ?, updated_at = ? where id_referensi = ?"
	_, err := tx.ExecContext(ctx, script, referensiarsitektur.Kode_referensi, referensiarsitektur.Nama_referensi, referensiarsitektur.Level_referensi, referensiarsitektur.Jenis_referensi,referensiarsitektur.Updated_at, referensiarsitektur.IdReferensi)
	helper.PanicIfError(err)

	return referensiarsitektur
}

func (repository *ReferensiArsitekturRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur){
	script := "delete from referensi_arsitekturs where id_referensi =?"
	_, err := tx.ExecContext(ctx, script, referensiarsitektur.IdReferensi)
	helper.PanicIfError(err)

}

func(repository *ReferensiArsitekturRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ReferensiArsitektur{
	script := "select id_referensi,kode_referensi,nama_referensi,level_referensi,jenis_referensi,created_at,updated_at from referensi_arsitekturs"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()
	
	var reference []domain.ReferensiArsitektur
	for rows.Next() {
		referensi := domain.ReferensiArsitektur{}
		err := rows.Scan(&referensi.IdReferensi, &referensi.Kode_referensi, &referensi.Nama_referensi, &referensi.Level_referensi, &referensi.Jenis_referensi,&referensi.Created_at,&referensi.Updated_at)
		helper.PanicIfError(err)
		reference = append(reference, referensi)
	}
	return reference
}

func(repository *ReferensiArsitekturRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, referensiarsitekturId int) (domain.ReferensiArsitektur, error){
	script := "select id_referensi, kode_referensi, nama_referensi, level_referensi, jenis_referensi, created_at, updated_at from referensi_arsitekturs where id_referensi = ?"
	rows, err := tx.QueryContext(ctx, script, referensiarsitekturId)
	helper.PanicIfError(err)
	defer rows.Close()

	referensiarsitektur := domain.ReferensiArsitektur{}
	if rows.Next() {
		err := rows.Scan(&referensiarsitektur.IdReferensi, &referensiarsitektur.Kode_referensi, &referensiarsitektur.Nama_referensi, &referensiarsitektur.Level_referensi, &referensiarsitektur.Jenis_referensi, &referensiarsitektur.Created_at, &referensiarsitektur.Updated_at)
		helper.PanicIfError(err)
		return referensiarsitektur, nil
	} else {
		return referensiarsitektur, errors.New("referensiarsitektur is not found")
	}
}

func (repository *ReferensiArsitekturRepositoryImpl) FindByKodeRef(ctx context.Context, tx *sql.Tx, kodeReferensi string) ([]domain.ReferensiArsitektur, error) {
	var exists bool
	err := tx.QueryRowContext(ctx, "select exists(select 1 from referensi_arsitekturs where kode_referensi = ?)", kodeReferensi).Scan(&exists)
	helper.PanicIfError(err)

	if !exists {
		log.Println("Data not found for kodeReferensi:", kodeReferensi)
		return nil, errors.New("data not found")
	}

	kodeBody := strings.Split(kodeReferensi, ".")
	var placeholders []string

	for i := range kodeBody {
		placeholders = append(placeholders, strings.Join(kodeBody[:i+1], "."))
	}

	script := "select id_referensi, kode_referensi, nama_referensi, level_referensi, jenis_referensi, created_at, updated_at from referensi_arsitekturs where kode_referensi in ("
	for i := range placeholders {
		if i > 0 {
			script += ", "
		}
		script += "?"
	}
	script += ") order by level_referensi asc"

	rows, err := tx.QueryContext(ctx, script, helper.ConvertStringsToInterfaces(placeholders)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var referensiList []domain.ReferensiArsitektur
	for rows.Next() {
		referensi := domain.ReferensiArsitektur{}
		err := rows.Scan(&referensi.IdReferensi, &referensi.Kode_referensi, &referensi.Nama_referensi, &referensi.Level_referensi, &referensi.Jenis_referensi, &referensi.Created_at, &referensi.Updated_at)
		if err != nil {
			return nil, err
		}
		referensiList = append(referensiList, referensi)
	}

	if len(referensiList) == 0 {
		log.Println("No hierarchical data found for kodeReferensi:", kodeReferensi)
		return nil, errors.New("data not found")
	}

	return referensiList, nil
}