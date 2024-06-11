package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
)

type SasaranKotaRepositoryImpl struct {
}

func NewSasaranKotaRepository() SasaranKotaRepository {
	return &SasaranKotaRepositoryImpl{}
}

func (repository *SasaranKotaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, sasarankotaId int) (domain.SasaranKota, error) {
	script := "select id, sasaran, strategi_kota, tujuan_kota, tahun, created_at, updated_at from sasaran_kota where id = ?"
	rows, err := tx.QueryContext(ctx, script, sasarankotaId)
	if err != nil {
		log.Printf("Error executing query: %v", err) 
		helper.PanicIfError(err)
	}
	defer rows.Close()

	sasaran := domain.SasaranKota{}
	if rows.Next() {
		err := rows.Scan(&sasaran.ID, &sasaran.Sasaran, &sasaran.StrategiKota, &sasaran.TujuanKota, &sasaran.Tahun, &sasaran.CreatedAt, &sasaran.UpdatedAt )
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			helper.PanicIfError(err)
		}
		log.Printf("Successfully retrieved Sasaran Kota: %+v", sasarankotaId)
		return sasaran, nil
	} else {
		return sasaran, errors.New("Sasaran kota is not found")
	}
}