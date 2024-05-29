package main

import (
	"api_spbe_kota_madiun/app"
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/repository"
	"api_spbe_kota_madiun/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func main() {

	db := app.GetConnection()
	validate :=validator.New()

	referensiarsitekturRepository := repository.NewReferensiArsitekturRepository()
	referesiarsitekturService := service.NewReferensiArsitekturService(referensiarsitekturRepository, db, validate)
	referensiarsitekturController := controller.NewReferensiarstitekturController(referesiarsitekturService)

	prosesbisnisRepository := repository.NewProsesBisnisRepository()
	prosesbisnisService  := service.NewProsesBisnisService(referensiarsitekturRepository, prosesbisnisRepository, db)
	prosesbisnisController := controller.NewProsesBisnisController(prosesbisnisService)

	router := app.NewRouter(referensiarsitekturController, prosesbisnisController)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})
	
	handler := c.Handler(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
		
	}

	fmt.Println("running", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
