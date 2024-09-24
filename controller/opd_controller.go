package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OpdController interface {
	FetchApiOpd(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllEksternal(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindKodeOpdUrusan(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
