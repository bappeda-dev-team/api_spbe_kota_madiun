package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BidangUrusanController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
