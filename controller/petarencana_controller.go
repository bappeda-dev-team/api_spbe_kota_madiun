package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PetarencanaController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
