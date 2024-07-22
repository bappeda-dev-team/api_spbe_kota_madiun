package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UrusanController interface {
	FetchApiUrusan(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
