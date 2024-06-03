package app

import (
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/exception"

	"github.com/julienschmidt/httprouter"
)
type RouteController struct{
	
}

func NewRouter(referensiarsitekturController controller.ReferensiArsitekturController, 
	prosesbisnisController controller.ProsesBisnisController) *httprouter.Router{

	router := httprouter.New()

	//referensi arsitektur router
	router.GET("/v1/referensiarsitektur", referensiarsitekturController.FindAll)
	router.POST("/v1/createreferensiarsitektur", referensiarsitekturController.Insert)
	router.PUT("/v1/updatereferensiarsitektur/:referensiarsitekturId", referensiarsitekturController.Update)
	router.DELETE("/v1/deletereferensiarsitektur/:referensiarsitekturId", referensiarsitekturController.Delete)
	router.GET("/v1/referensiarsitektur/:kodeReferensi", referensiarsitekturController.FindByKodeReferensi)
	router.GET("/v1/referensiarsitekturbyid/:referensiarsitekturId", referensiarsitekturController.FindById)
	
	//proses bisnis router
	router.GET("/v1/prosesbisnis", prosesbisnisController.FindByKodeOPD)
	router.GET("/v1/prosesbisnis/:kodeOPD", prosesbisnisController.FindByKodeOPD)
	router.GET("/v1/prosesbisnisbytahun/:tahun", prosesbisnisController.FindByKodeOPD)
	router.GET("/v1/prosesbisnis/:kodeOPD/:tahun", prosesbisnisController.FindByKodeOPD)
	router.GET("/v1/prosesbisnisbyid/:prosesbisnisId", prosesbisnisController.FindById)
	router.POST("/v1/createprosesbisnis", prosesbisnisController.Insert)
	router.PUT("/v1/updateprosesbisnis/:prosesbisnisId", prosesbisnisController.Update)
	router.DELETE("/v1/deleteprosesbisnis/:prosesbisnisId", prosesbisnisController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
