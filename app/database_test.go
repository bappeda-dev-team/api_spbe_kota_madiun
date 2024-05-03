package app

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TestDatabase(t *testing.T) {

	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("%s@tcp(localhost:%s)/%s?parseTime=true", dbUser, dbPort, dbName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	
	defer db.Close()
}