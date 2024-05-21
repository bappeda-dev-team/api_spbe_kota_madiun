package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProsesBisnisController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}