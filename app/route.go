package app

import (
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/exception"

	"github.com/julienschmidt/httprouter"
)

type RouteController struct {
}

func NewRouter(referensiarsitekturController controller.ReferensiArsitekturController,
	prosesbisnisController controller.ProsesBisnisController, sasarankotaController controller.SasaranKotaController, pohonkinerja controller.PohonKinerjaController, bidangurusan controller.BidangUrusanController, opdController controller.OpdController, urusanController controller.UrusanController, layananspbeController controller.LayananSpbeController, datainformasiController controller.DataDanInformasiController, aplikasiController controller.AplikasiController, domainspbeController controller.DomainSPBEController, kebutuhanSPBEController controller.KebutuhanSPBEController, userController controller.UserController) *httprouter.Router {
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
	router.GET("/v1/prosesbisnisnogap", prosesbisnisController.GetProsesBisnisNoGap)
	router.GET("/v1/prosesbisnisnogap/:kodeOPD", prosesbisnisController.GetProsesBisnisNoGap)
	router.GET("/v1/prosesbisnisnogap/:kodeOPD/:tahun", prosesbisnisController.GetProsesBisnisNoGap)
	router.GET("/v1/prosesbisnisbyid/:prosesbisnisId", prosesbisnisController.FindById)
	router.POST("/v1/createprosesbisnis", prosesbisnisController.Insert)
	router.PUT("/v1/updateprosesbisnis/:prosesbisnisId", prosesbisnisController.Update)
	router.DELETE("/v1/deleteprosesbisnis/:prosesbisnisId", prosesbisnisController.Delete)

	//sasaran kota
	router.GET("/v1/sasarankota/:sasarankotaId", sasarankotaController.FindById)
	router.GET("/v1/sasarankota", sasarankotaController.FindAll)

	//pohon kinerja
	router.GET("/v1/pohonkinerja/:pohonId", pohonkinerja.FindById)
	router.GET("/v1/pohonkinerja", pohonkinerja.FindAll)

	//bidangurusan
	router.GET("/v1/bidangurusan", bidangurusan.FindAll)

	//layanan spbe
	router.GET("/v1/layananspbe", layananspbeController.FindByKodeOPD)
	router.GET("/v1/layananspbe/:kodeOPD", layananspbeController.FindByKodeOPD)
	router.GET("/v1/layananspbebytahun/:tahun", layananspbeController.FindByKodeOPD)
	router.GET("/v1/layananspbe/:kodeOPD/:tahun", layananspbeController.FindByKodeOPD)
	router.GET("/v1/layananspbebyid/:layananspbeId", layananspbeController.FindById)
	router.POST("/v1/createlayananspbe", layananspbeController.Insert)
	router.PUT("/v1/updatelayananspbe/:layananspbeId", layananspbeController.Update)
	router.DELETE("/v1/deletelayananspbe/:layananspbeId", layananspbeController.Delete)

	//data dan informasi
	router.GET("/v1/datainformasi", datainformasiController.FindByKodeOPD)
	router.GET("/v1/datainformasi/:kodeOPD", datainformasiController.FindByKodeOPD)
	router.GET("/v1/datainformasibytahun/:tahun", datainformasiController.FindByKodeOPD)
	router.GET("/v1/datainformasi/:kodeOPD/:tahun", datainformasiController.FindByKodeOPD)
	router.GET("/v1/datainformasibyid/:dataId", datainformasiController.FindById)
	router.POST("/v1/createdatainformasi", datainformasiController.Insert)
	router.PUT("/v1/updatedatainformasi/:dataId", datainformasiController.Update)
	router.DELETE("/v1/deletedatainformasi/:dataId", datainformasiController.Delete)

	//aplikasi
	router.GET("/v1/aplikasi", aplikasiController.FindByKodeOPD)
	router.GET("/v1/aplikasi/:kodeOPD", aplikasiController.FindByKodeOPD)
	router.GET("/v1/aplikasibytahun/:tahun", aplikasiController.FindByKodeOPD)
	router.GET("/v1/aplikasi/:kodeOPD/:tahun", aplikasiController.FindByKodeOPD)
	router.GET("/v1/aplikasibyid/:aplikasiId", aplikasiController.FindById)
	router.POST("/v1/createaplikasi", aplikasiController.Insert)
	router.PUT("/v1/updateaplikasi/:aplikasiId", aplikasiController.Update)
	router.DELETE("/v1/deleteaplikasi/:aplikasiId", aplikasiController.Delete)

	//domainspbe
	router.GET("/v1/domainspbe", domainspbeController.FindAll)
	router.GET("/v1/domainspbe/:domainSPBEId", domainspbeController.FindById)
	router.POST("/v1/createdomainspbe", domainspbeController.Insert)
	router.PUT("/v1/updatedomainspbe/:domainSPBEId", domainspbeController.Update)
	router.DELETE("/v1/deletedomainspbe/:domainSPBEId", domainspbeController.Delete)

	//kebutuhanspbe
	router.GET("/v1/kebutuhanspbe", kebutuhanSPBEController.FindByKodeOpdAndTahun)
	router.GET("/v1/kebutuhanspbe/:kodeOpd", kebutuhanSPBEController.FindByKodeOpdAndTahun)
	router.GET("/v1/kebutuhanspbebytahun/:tahun", kebutuhanSPBEController.FindByKodeOpdAndTahun)
	router.GET("/v1/kebutuhanspbe/:kodeOpd/:tahun", kebutuhanSPBEController.FindByKodeOpdAndTahun)
	router.GET("/v1/kebutuhanspbebyid/:kebutuhanSPBEId", kebutuhanSPBEController.FindById)
	router.POST("/v1/createkebutuhanspbe", kebutuhanSPBEController.Create)
	router.PUT("/v1/updatekebutuhanspbe/:kebutuhanSPBEId", kebutuhanSPBEController.Update)
	router.DELETE("/v1/deletekebutuhanspbe/:kebutuhanSPBEId", kebutuhanSPBEController.Delete)

	//fetch api
	router.GET("/sasarankotafetch", sasarankotaController.Insert)
	router.GET("/pohonkinerjafetch", pohonkinerja.FetchApiPohon)
	router.GET("/opdafetch", opdController.FetchApiOpd)
	router.GET("/urusanfetch", urusanController.FetchApiUrusan)
	router.GET("/bidangurusanfetch", bidangurusan.FetchBidangUrusan)
	router.GET("/userapifetch", userController.InsertApi)

	//GAP SPBE
	router.GET("/v1/GapSPBE", prosesbisnisController.GetProsesBisnisGrouped)
	router.GET("/v1/gapSPBE/:kodeOPD", prosesbisnisController.GetProsesBisnisGrouped)

	//user
	router.POST("/v1/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
