package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProsesBisnisController interface {
	FindByKodeOPD(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProsesBisnisGrouped(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProsesBisnisNoGap(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ExportExcel(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
