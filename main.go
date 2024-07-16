package main

import (
	"api_spbe_kota_madiun/app"
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/repository"
	"api_spbe_kota_madiun/service"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func main() {

	db := app.GetConnection()
	validate := validator.New()

	//referensi arsitektur
	referensiarsitekturRepository := repository.NewReferensiArsitekturRepository()
	referesiarsitekturService := service.NewReferensiArsitekturService(referensiarsitekturRepository, db, validate)
	referensiarsitekturController := controller.NewReferensiarstitekturController(referesiarsitekturService)

	//proses bisnis
	prosesbisnisRepository := repository.NewProsesBisnisRepository()
	pohonkinerjaRepository := repository.NewPohonKinerjaRepository()
	sasarankotaRepository := repository.NewSasaranKotaRepository()
	bidangurusanRepository := repository.NewBidangUrusanRepository()
	prosesbisnisService := service.NewProsesBisnisService(referensiarsitekturRepository, sasarankotaRepository, bidangurusanRepository, prosesbisnisRepository, pohonkinerjaRepository, db, validate)
	prosesbisnisController := controller.NewProsesBisnisController(prosesbisnisService)

	//sasaran kota
	sasarankotaService := service.NewSasaranKotaService(sasarankotaRepository, db)
	sasarankotaController := controller.NewSasaranKotaController(sasarankotaService)

	//pohon kinerja
	pohonkinerjaService := service.NewPohonKinerjaService(pohonkinerjaRepository, db)
	pohonkinerjaController := controller.NewPohonKinerjaController(pohonkinerjaService)

	//bidang Urusan
	bidangurusanService := service.NewBidangUrusanService(bidangurusanRepository, db)
	bidangurusanController := controller.NewBidangUrusanController(bidangurusanService)

	//layanan spbe
	layananspbeRepository := repository.NewLayananSPBERepository()
	layananspbeService := service.NewLayananSpbeService(layananspbeRepository, pohonkinerjaRepository, referensiarsitekturRepository, db, validate)
	layananspbeController := controller.NewLayananSPBEController(layananspbeService)

	//fetch opd
	opdRepository := repository.NewOpdRepository()
	opdService := service.NewOpdService(opdRepository, db)
	opdController := controller.NewOpdController(opdService)

	// fetch urusan
	urusanRepository := repository.NewUrusanRepository()
	urusanService := service.NewUrusanService(urusanRepository, db)
	urusanController := controller.NewUrusanController(urusanService)

	router := app.NewRouter(referensiarsitekturController, prosesbisnisController, sasarankotaController, pohonkinerjaController, bidangurusanController, opdController, urusanController, layananspbeController)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

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
		//with auth
		// Handler: middleware.NewAuthMiddleware(handler),
	}
	fmt.Println("running", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
