package main

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	host := os.Getenv("host")
	port := os.Getenv("port")
	addr := fmt.Sprintf("%s:%s", host, port)

	if addr == ":" {
		addr = "localhost:8080"
	}

	return &http.Server{
		Addr:    addr,
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializeServer()

	log.Printf("Server berjalan di %s", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
