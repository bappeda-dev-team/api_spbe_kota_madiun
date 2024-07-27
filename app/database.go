package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		// log.Fatal("Error loading .env file")
	}

	// !! diganti dengan DB_URL agar bisa konek di docker-compose
	// contoh: "root@tcp(localhost:3306)/db_spbe?parseTime=true"
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT") // Ensure this is set to something like "3306"
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	// if dbHost == "" {
	// 	dbHost = "localhost"
	// }

	connStr := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	// connStr := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
