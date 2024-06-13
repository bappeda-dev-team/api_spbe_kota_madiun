package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
)

type BidangUrusanRepositoryImpl struct {
}

func NewBidangUrusanRepository() BidangUrusanRepository {
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
		return bidangUrusan, errors.New("Bidang urusan is not found")
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
