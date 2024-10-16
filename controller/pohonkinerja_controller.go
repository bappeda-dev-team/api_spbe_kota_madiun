package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PohonKinerjaController interface {
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchApiPohon(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByOperational(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetHierarchy(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
