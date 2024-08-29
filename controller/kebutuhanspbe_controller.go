package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type KebutuhanSPBEController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateKeterangan(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePenanggungJawab(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByKodeOpdAndTahun(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
