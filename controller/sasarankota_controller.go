package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SasaranKotaController interface {
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
