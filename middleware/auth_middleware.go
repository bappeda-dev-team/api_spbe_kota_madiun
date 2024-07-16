package middleware

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") == "espbe" {
		//oke
		middleware.Handler.ServeHTTP(writter, request)
	} else {
		//error
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writter, webResponse)

	}
}
