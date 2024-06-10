package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
)

type PohonKinerjaRepositoryImpl struct {
}

func NewPohonKinerjaRepository() PohonKinerjaRepository {
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
		return pohonKinerja, errors.New("Pohon Kinerja is not found")
	}
}
