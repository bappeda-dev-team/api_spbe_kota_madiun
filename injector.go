//go:build wireinject
// +build wireinject

package main

import (
	"api_spbe_kota_madiun/app"
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/middleware"
	"api_spbe_kota_madiun/repository"
	"api_spbe_kota_madiun/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var referensiArsitekturSet = wire.NewSet(
	repository.NewReferensiArsitekturRepositoryImpl,
	wire.Bind(new(repository.ReferensiArsitekturRepository), new(*repository.ReferensiArsitekturRepositoryImpl)),
	service.NewReferensiArsitekturServiceImpl,
	wire.Bind(new(service.ReferensiArsitekturService), new(*service.ReferensiArsitekturServiceImpl)),
	controller.NewReferensiArsitekturControllerImpl,
	wire.Bind(new(controller.ReferensiArsitekturController), new(*controller.ReferensiArsitekturControllerImpl)),
)

var prosesbisnisSet = wire.NewSet(
	repository.NewProsesBisnisRepositoryImpl,
	wire.Bind(new(repository.ProsesBisnisRepository), new(*repository.ProsesBisnisRepositoryImpl)),
	service.NewProsesBisnisServiceImpl,
	wire.Bind(new(service.ProsesBisnisService), new(*service.ProsesBisnisServiceImpl)),
	controller.NewProsesBisnisControllerImpl,
	wire.Bind(new(controller.ProsesBisnisController), new(*controller.ProsesBisnisControllerImpl)),
)

var layananSet = wire.NewSet(
	repository.NewLayananSPBERepositoryImpl,
	wire.Bind(new(repository.LayananSPBERepository), new(*repository.LayananSPBERepositoryImpl)),
	service.NewLayananSpbeServiceImpl,
	wire.Bind(new(service.LayananSpbeService), new(*service.LayananSpbeServiceImpl)),
	controller.NewLayananSPBEControllerImpl,
	wire.Bind(new(controller.LayananSpbeController), new(*controller.LayananSpbeControllerImpl)),
)

var datadaninformasiSet = wire.NewSet(
	repository.NewDataDanInformasiRepositoryImpl,
	wire.Bind(new(repository.DataDanInformasiRepository), new(*repository.DataDanInformasiRepositoryImpl)),
	service.NewDataDanInformasiServiceImpl,
	wire.Bind(new(service.DataDanInformasiService), new(*service.DataDanInformasiServiceImpl)),
	controller.NewDataDanInformasiControllerImpl,
	wire.Bind(new(controller.DataDanInformasiController), new(*controller.DataDanInformasiControllerImpl)),
)

var aplikasiSet = wire.NewSet(
	repository.NewAplikasiRepositoryImpl,
	wire.Bind(new(repository.AplikasiRepository), new(*repository.AplikasiRepositoryImpl)),
	service.NewAplikasiServiceImpl,
	wire.Bind(new(service.AplikasiService), new(*service.AplikasiServiceImpl)),
	controller.NewAplikasiControllerImpl,
	wire.Bind(new(controller.AplikasiController), new(*controller.AplikasiControllerImpl)),
)

var bidangurusanSet = wire.NewSet(
	repository.NewBidangUrusanRepositoryImpl,
	wire.Bind(new(repository.BidangUrusanRepository), new(*repository.BidangUrusanRepositoryImpl)),
	service.NewBidangUrusanServiceImpl,
	wire.Bind(new(service.BidangUrusanService), new(*service.BidangUrusanServiceImpl)),
	controller.NewBidangUrusanControllerImpl,
	wire.Bind(new(controller.BidangUrusanController), new(*controller.BidangUrusanControllerImpl)),
)

var opdSet = wire.NewSet(
	repository.NewOpdRepositoryImpl,
	wire.Bind(new(repository.OpdRepository), new(*repository.OpdRepositoryImpl)),
	service.NewOpdServiceImpl,
	wire.Bind(new(service.OpdService), new(*service.OpdServiceImpl)),
	controller.NewOpdControllerImpl,
	wire.Bind(new(controller.OpdController), new(*controller.OpdControllerImpl)),
)

var pohonkinerjaSet = wire.NewSet(
	repository.NewPohonKinerjaRepositoryImpl,
	wire.Bind(new(repository.PohonKinerjaRepository), new(*repository.PohonKinerjaRepositoryImpl)),
	service.NewPohonKinerjaServiceImpl,
	wire.Bind(new(service.PohonKinerjaService), new(*service.PohonKinerjaServiceImpl)),
	controller.NewPohonKinerjaControllerImpl,
	wire.Bind(new(controller.PohonKinerjaController), new(*controller.PohonKinerjaControllerImpl)),
)

var sasarankotaSet = wire.NewSet(
	repository.NewSasaranKotaRepositoryImpl,
	wire.Bind(new(repository.SasaranKotaRepository), new(*repository.SasaranKotaRepositoryImpl)),
	service.NewSasaranKotaServiceImpl,
	wire.Bind(new(service.SasaranKotaService), new(*service.SasaranKotaServiceImpl)),
	controller.NewSasaranKotaControllerImpl,
	wire.Bind(new(controller.SasaranKotaController), new(*controller.SasaranKotaControllerImpl)),
)

var urusanSet = wire.NewSet(
	repository.NewUrusanRepositoryImpl,
	wire.Bind(new(repository.UrusanRepository), new(*repository.UrusanRepositoryImpl)),
	service.NewUrusanServiceImpl,
	wire.Bind(new(service.UrusanService), new(*service.UrusanServiceImpl)),
	controller.NewUrusanControllerImpl,
	wire.Bind(new(controller.UrusanController), new(*controller.UrusanControllerImpl)),
)

var domainspbeSet = wire.NewSet(
	repository.NewDomainSPBERepositoryImpl,
	wire.Bind(new(repository.DomainSPBERepository), new(*repository.DomainSPBERepositoryImpl)),
	service.NewDomainSPBEServiceImpl,
	wire.Bind(new(service.DomainSPBEService), new(*service.DomainSPBEServiceImpl)),
	controller.NewDomainSPBEControllerImpl,
	wire.Bind(new(controller.DomainSPBEController), new(*controller.DomainSPBEControllerImpl)),
)

var kebutuhanspbeSet = wire.NewSet(
	repository.NewKebutuhanSPBERepositoryImpl,
	wire.Bind(new(repository.KebutuhanSPBERepository), new(*repository.KebutuhanSPBERepositoryImpl)),
	service.NewKebutuhanSPBEServiceImpl,
	wire.Bind(new(service.KebutuhanSPBEService), new(*service.KebutuhanSPBEServiceImpl)),
	controller.NewKebutuhanSPBEControllerImpl,
	wire.Bind(new(controller.KebutuhanSPBEController), new(*controller.KebutuhanSPBEControllerImpl)),
)

var sasarankinerjaSet = wire.NewSet(
	repository.NewSasaranKinerjaPegawaiRepositoryImpl,
	wire.Bind(new(repository.SasaranKinerjaPegawaiRepository), new(*repository.SasaranKinerjaPegawaiRepositoryImpl)),
	service.NewSasaranKinerjaPegawaiServiceImpl,
	wire.Bind(new(service.SasaranKinerjaPegawaiService), new(*service.SasaranKinerjaPegawaiServiceImpl)),
	controller.NewSasaranKinerjaPegawaiControllerImpl,
	wire.Bind(new(controller.SasaranKinerjaPegawaiController), new(*controller.SasaranKinerjaPegawaiControllerImpl)),
)

var userSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserServiceImpl,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserControllerImpl,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var rencanaPelaksanaanSet = wire.NewSet(
	repository.NewRencanaPelaksanaanRepositoryImpl,
	wire.Bind(new(repository.RencanaPelaksanaanRepository), new(*repository.RencanaPelaksanaanRepositoryImpl)),
	service.NewRencanaPelaksanaanServiceImpl,
	wire.Bind(new(service.RencanaPelaksanaanService), new(*service.RencanaPelaksanaanServiceImpl)),
	controller.NewRencanaPelaksanaanControllerImpl,
	wire.Bind(new(controller.RencanaPelaksanaanController), new(*controller.RencanaPelaksanaanControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		referensiArsitekturSet,
		prosesbisnisSet,
		layananSet,
		datadaninformasiSet,
		aplikasiSet,
		bidangurusanSet,
		opdSet,
		urusanSet,
		domainspbeSet,
		kebutuhanspbeSet,
		pohonkinerjaSet,
		sasarankotaSet,
		sasarankinerjaSet,
		userSet,
		rencanaPelaksanaanSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
