package repository

// import (
// 	"api_spbe_kota_madiun/app"
// 	"api_spbe_kota_madiun/model/domain"
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func TestReferensiArsitekturInsert(t *testing.T) {
// 	referensiRepository := NewReferensiArsitekturRepository()
// 	db := app.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		t.Fatalf("Failed to begin transaction: %v", err)
// 	}
// 	defer tx.Rollback()

// 	currentTime := time.Now()
// 	referensiArsitektur := domain.ReferensiArsitektur{
// 		Kode_referensi: "RAB.01.01.02 ",
// 		Nama_referensi: "PERTAHANAN DAN LUAR NEGERI",
// 		Level_referensi: "1",
// 		Jenis_referensi: "ProsesBisnis",
// 		Created_at:     currentTime,

// 	}

// 	result := referensiRepository.Insert(ctx, tx, referensiArsitektur)
// 	if err != nil {
// 		t.Fatalf("Failed to insert kode: %v", err)
// 	}

// 	if err := tx.Commit(); err != nil {
// 		t.Fatalf("Failed to commit transaction: %v", err)
// 	}

// 	fmt.Println(result)
// }

// func TestReferensiArsitekturUpdate(t *testing.T) {
// 	referensiRepository := NewReferensiArsitekturRepository()
// 	db := app.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		t.Fatalf("Failed to begin transaction: %v", err)
// 	}
// 	defer tx.Rollback()

// 	currentTime := time.Now()
// 	referensiArsitektur := domain.ReferensiArsitektur{
// 		IdReferensi:      2,
// 		Kode_referensi:  "RAB.01.01.01",
// 		Nama_referensi: "PERTAHANAN",
// 		Level_referensi: "2",
// 		Jenis_referensi: "ProsesBisnis",
// 		Created_at: currentTime,
// 		Updated_at:     currentTime,
// 	}

// 	result := referensiRepository.Update(ctx, tx, referensiArsitektur)
// 	if err != nil {
// 		t.Fatalf("Failed to insert kode: %v", err)
// 	}

// 	if err := tx.Commit(); err != nil {
// 		t.Fatalf("Failed to commit transaction: %v", err)
// 	}

// 	fmt.Println(result)
// }

// func TestReferensiArsitekturFinfByKode(t *testing.T) {
// 	// Inisialisasi repository dan koneksi database
// 	referensiRepository := NewReferensiArsitekturRepository()
// 	db := app.GetConnection()
// 	defer db.Close()

// 	// Mulai transaksi
// 	ctx := context.Background()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		t.Fatalf("Failed to begin transaction: %v", err)
// 	}
// 	defer tx.Rollback()

// 	// Kode referensi yang akan dicari
// 	kodeReferensi := "RAB.0"
// 	result := referensiRepository.FindByKodeRef(ctx, tx, kodeReferensi)

// 	// // Periksa apakah hasilnya tidak kosong
// 	// if len(result) == 0 {
// 	// 	t.Fatalf("No results found for kode: %s", kodeReferensi)
// 	// }

// 	// Komit transaksi jika tidak ada kesalahan
// 	if err := tx.Commit(); err != nil {
// 		t.Fatalf("Failed to commit transaction: %v", err)
// 	}

// 	// Cetak hasil pencarian
// 	fmt.Println(result)

// }
