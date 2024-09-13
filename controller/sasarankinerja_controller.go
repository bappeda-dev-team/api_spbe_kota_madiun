package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SasaranKinerjaPegawaiController interface {
	FetchsasaranKinerja(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByKodeOpdAndTahun(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
