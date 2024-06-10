package main

import (
	"api_spbe_kota_madiun/app"
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/repository"
	"api_spbe_kota_madiun/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func main() {

	db := app.GetConnection()
	validate := validator.New()

	referensiarsitekturRepository := repository.NewReferensiArsitekturRepository()
	referesiarsitekturService := service.NewReferensiArsitekturService(referensiarsitekturRepository, db, validate)
	referensiarsitekturController := controller.NewReferensiarstitekturController(referesiarsitekturService)

	prosesbisnisRepository := repository.NewProsesBisnisRepository()
	pohonkinerjaRepository := repository.NewPohonKinerjaRepository()
	prosesbisnisService := service.NewProsesBisnisService(referensiarsitekturRepository, prosesbisnisRepository, pohonkinerjaRepository, db, validate)
	prosesbisnisController := controller.NewProsesBisnisController(prosesbisnisService)

	router := app.NewRouter(referensiarsitekturController, prosesbisnisController)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	var wg sync.WaitGroup
	wg.Add(3)
	ctx := context.Background()
	go app.FetchPohonKinerja(ctx, db, &wg)
	go app.FetchKodeOpd(ctx, db, &wg)
	go app.FetchSaranKota(ctx, db, &wg)

	handler := c.Handler(router)
	// ambil host dan port di env
	// agar bisa disetting diluar build
	// localhost tidak bisa dipakai untuk production
	host := os.Getenv("host")
	port := os.Getenv("port")
	addr := fmt.Sprintf("%s:%s", host, port)

	// set default value host and port
	if addr == ":" {
		addr = "localhost:8080"
	}

	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	fmt.Println("running", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
