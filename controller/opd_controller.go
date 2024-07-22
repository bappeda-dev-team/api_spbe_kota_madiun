package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OpdController interface {
	FetchApiOpd(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
