package repository

import (
	"api_spbe_kota_madiun/app"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestReferensiArsitekturInsert(t *testing.T) {
	referensiRepository := NewReferensiKodeRepository()
	db := app.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin() 
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback() 

	currentTime := time.Now()
	referensiArsitektur := domain.ReferensiArsitektur{
		Kode_referensi: "RAB.01 ",
		Nama_referensi: "PERTAHANAN DAN LUAR NEGERI",
		Level_referensi: "1",
		Jenis_referensi: "ProsesBisnis",
		Created_at:     currentTime,
		
	}

	result := referensiRepository.Insert(ctx, tx, referensiArsitektur) 
	if err != nil {
		t.Fatalf("Failed to insert kode: %v", err)
	}

	if err := tx.Commit(); err != nil { 
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println(result)
}

func TestReferensiArsitekturUpdate(t *testing.T) {
	referensiRepository := NewReferensiKodeRepository()
	db := app.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin() 
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback() 

	currentTime := time.Now()
	referensiArsitektur := domain.ReferensiArsitektur{
		IdReferensi:      2,
		Kode_referensi:  "RAB.01.01.01",
		Nama_referensi: "PERTAHANAN",
		Level_referensi: "2",
		Jenis_referensi: "ProsesBisnis",
		Created_at: currentTime,
		Updated_at:     currentTime,
	}

	result := referensiRepository.Update(ctx, tx, referensiArsitektur) 
	if err != nil {
		t.Fatalf("Failed to insert kode: %v", err)
	}

	if err := tx.Commit(); err != nil { 
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println(result)
}

func TestReferensiArsitekturFinfByKode(t *testing.T) {
	referensiRepository := NewReferensiKodeRepository()
	db := app.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin() 
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback() 

	kodeReferensi := "RAB.01.01.01"
	result := referensiRepository.FindByKodeRef(ctx, tx, kodeReferensi)


	if err != nil {
		t.Fatalf("Failed to insert kode: %v", err)
	}

	if err := tx.Commit(); err != nil { 
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println(result)

}