package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	InsertApi(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ChangePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
