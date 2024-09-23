// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"api_spbe_kota_madiun/app"
	"api_spbe_kota_madiun/controller"
	"api_spbe_kota_madiun/middleware"
	"api_spbe_kota_madiun/repository"
	"api_spbe_kota_madiun/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

// Injectors from injector.go:

func InitializeServer() *http.Server {
	referensiArsitekturRepositoryImpl := repository.NewReferensiArsitekturRepositoryImpl()
	db := app.GetConnection()
	v := _wireValue
	validate := validator.New(v...)
	referensiArsitekturServiceImpl := service.NewReferensiArsitekturServiceImpl(referensiArsitekturRepositoryImpl, db, validate)
	referensiArsitekturControllerImpl := controller.NewReferensiArsitekturControllerImpl(referensiArsitekturServiceImpl)
	sasaranKotaRepositoryImpl := repository.NewSasaranKotaRepositoryImpl()
	bidangUrusanRepositoryImpl := repository.NewBidangUrusanRepositoryImpl()
	prosesBisnisRepositoryImpl := repository.NewProsesBisnisRepositoryImpl()
	pohonKinerjaRepositoryImpl := repository.NewPohonKinerjaRepositoryImpl()
	prosesBisnisServiceImpl := service.NewProsesBisnisServiceImpl(referensiArsitekturRepositoryImpl, sasaranKotaRepositoryImpl, bidangUrusanRepositoryImpl, prosesBisnisRepositoryImpl, pohonKinerjaRepositoryImpl, db, validate)
	prosesBisnisControllerImpl := controller.NewProsesBisnisControllerImpl(prosesBisnisServiceImpl)
	sasaranKotaServiceImpl := service.NewSasaranKotaServiceImpl(sasaranKotaRepositoryImpl, db)
	sasaranKotaControllerImpl := controller.NewSasaranKotaControllerImpl(sasaranKotaServiceImpl)
	pohonKinerjaServiceImpl := service.NewPohonKinerjaServiceImpl(pohonKinerjaRepositoryImpl, db)
	pohonKinerjaControllerImpl := controller.NewPohonKinerjaControllerImpl(pohonKinerjaServiceImpl)
	bidangUrusanServiceImpl := service.NewBidangUrusanServiceImpl(bidangUrusanRepositoryImpl, db)
	bidangUrusanControllerImpl := controller.NewBidangUrusanControllerImpl(bidangUrusanServiceImpl)
	opdRepositoryImpl := repository.NewOpdRepositoryImpl()
	urusanRepositoryImpl := repository.NewUrusanRepositoryImpl()
	opdServiceImpl := service.NewOpdServiceImpl(opdRepositoryImpl, urusanRepositoryImpl, bidangUrusanRepositoryImpl, db)
	opdControllerImpl := controller.NewOpdControllerImpl(opdServiceImpl)
	urusanServiceImpl := service.NewUrusanServiceImpl(urusanRepositoryImpl, db)
	urusanControllerImpl := controller.NewUrusanControllerImpl(urusanServiceImpl)
	layananSPBERepositoryImpl := repository.NewLayananSPBERepositoryImpl()
	layananSpbeServiceImpl := service.NewLayananSpbeServiceImpl(layananSPBERepositoryImpl, pohonKinerjaRepositoryImpl, referensiArsitekturRepositoryImpl, db, validate)
	layananSpbeControllerImpl := controller.NewLayananSPBEControllerImpl(layananSpbeServiceImpl)
	dataDanInformasiRepositoryImpl := repository.NewDataDanInformasiRepositoryImpl()
	dataDanInformasiServiceImpl := service.NewDataDanInformasiServiceImpl(dataDanInformasiRepositoryImpl, pohonKinerjaRepositoryImpl, referensiArsitekturRepositoryImpl, db, validate)
	dataDanInformasiControllerImpl := controller.NewDataDanInformasiControllerImpl(dataDanInformasiServiceImpl)
	aplikasiRepositoryImpl := repository.NewAplikasiRepositoryImpl()
	aplikasiServiceImpl := service.NewAplikasiServiceImpl(aplikasiRepositoryImpl, pohonKinerjaRepositoryImpl, referensiArsitekturRepositoryImpl, db, validate)
	aplikasiControllerImpl := controller.NewAplikasiControllerImpl(aplikasiServiceImpl)
	domainSPBERepositoryImpl := repository.NewDomainSPBERepositoryImpl()
	domainSPBEServiceImpl := service.NewDomainSPBEServiceImpl(domainSPBERepositoryImpl, db, validate)
	domainSPBEControllerImpl := controller.NewDomainSPBEControllerImpl(domainSPBEServiceImpl)
	kebutuhanSPBERepositoryImpl := repository.NewKebutuhanSPBERepositoryImpl()
	rencanaPelaksanaanRepositoryImpl := repository.NewRencanaPelaksanaanRepositoryImpl()
	sasaranKinerjaPegawaiRepositoryImpl := repository.NewSasaranKinerjaPegawaiRepositoryImpl()
	kebutuhanSPBEServiceImpl := service.NewKebutuhanSPBEServiceImpl(kebutuhanSPBERepositoryImpl, prosesBisnisRepositoryImpl, db, rencanaPelaksanaanRepositoryImpl, sasaranKinerjaPegawaiRepositoryImpl, opdRepositoryImpl)
	kebutuhanSPBEControllerImpl := controller.NewKebutuhanSPBEControllerImpl(kebutuhanSPBEServiceImpl)
	userRepositoryImpl := repository.NewUserRepositoryImpl()
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl, db)
	userControllerImpl := controller.NewUserControllerImpl(userServiceImpl)
	sasaranKinerjaPegawaiServiceImpl := service.NewSasaranKinerjaPegawaiServiceImpl(sasaranKinerjaPegawaiRepositoryImpl, db)
	sasaranKinerjaPegawaiControllerImpl := controller.NewSasaranKinerjaPegawaiControllerImpl(sasaranKinerjaPegawaiServiceImpl)
	rencanaPelaksanaanServiceImpl := service.NewRencanaPelaksanaanServiceImpl(rencanaPelaksanaanRepositoryImpl, sasaranKinerjaPegawaiRepositoryImpl, kebutuhanSPBERepositoryImpl, opdRepositoryImpl, db)
	rencanaPelaksanaanControllerImpl := controller.NewRencanaPelaksanaanControllerImpl(rencanaPelaksanaanServiceImpl)
	router := app.NewRouter(referensiArsitekturControllerImpl, prosesBisnisControllerImpl, sasaranKotaControllerImpl, pohonKinerjaControllerImpl, bidangUrusanControllerImpl, opdControllerImpl, urusanControllerImpl, layananSpbeControllerImpl, dataDanInformasiControllerImpl, aplikasiControllerImpl, domainSPBEControllerImpl, kebutuhanSPBEControllerImpl, userControllerImpl, sasaranKinerjaPegawaiControllerImpl, rencanaPelaksanaanControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

var (
	_wireValue = []validator.Option{}
)

// injector.go:

var referensiArsitekturSet = wire.NewSet(repository.NewReferensiArsitekturRepositoryImpl, wire.Bind(new(repository.ReferensiArsitekturRepository), new(*repository.ReferensiArsitekturRepositoryImpl)), service.NewReferensiArsitekturServiceImpl, wire.Bind(new(service.ReferensiArsitekturService), new(*service.ReferensiArsitekturServiceImpl)), controller.NewReferensiArsitekturControllerImpl, wire.Bind(new(controller.ReferensiArsitekturController), new(*controller.ReferensiArsitekturControllerImpl)))

var prosesbisnisSet = wire.NewSet(repository.NewProsesBisnisRepositoryImpl, wire.Bind(new(repository.ProsesBisnisRepository), new(*repository.ProsesBisnisRepositoryImpl)), service.NewProsesBisnisServiceImpl, wire.Bind(new(service.ProsesBisnisService), new(*service.ProsesBisnisServiceImpl)), controller.NewProsesBisnisControllerImpl, wire.Bind(new(controller.ProsesBisnisController), new(*controller.ProsesBisnisControllerImpl)))

var layananSet = wire.NewSet(repository.NewLayananSPBERepositoryImpl, wire.Bind(new(repository.LayananSPBERepository), new(*repository.LayananSPBERepositoryImpl)), service.NewLayananSpbeServiceImpl, wire.Bind(new(service.LayananSpbeService), new(*service.LayananSpbeServiceImpl)), controller.NewLayananSPBEControllerImpl, wire.Bind(new(controller.LayananSpbeController), new(*controller.LayananSpbeControllerImpl)))

var datadaninformasiSet = wire.NewSet(repository.NewDataDanInformasiRepositoryImpl, wire.Bind(new(repository.DataDanInformasiRepository), new(*repository.DataDanInformasiRepositoryImpl)), service.NewDataDanInformasiServiceImpl, wire.Bind(new(service.DataDanInformasiService), new(*service.DataDanInformasiServiceImpl)), controller.NewDataDanInformasiControllerImpl, wire.Bind(new(controller.DataDanInformasiController), new(*controller.DataDanInformasiControllerImpl)))

var aplikasiSet = wire.NewSet(repository.NewAplikasiRepositoryImpl, wire.Bind(new(repository.AplikasiRepository), new(*repository.AplikasiRepositoryImpl)), service.NewAplikasiServiceImpl, wire.Bind(new(service.AplikasiService), new(*service.AplikasiServiceImpl)), controller.NewAplikasiControllerImpl, wire.Bind(new(controller.AplikasiController), new(*controller.AplikasiControllerImpl)))

var bidangurusanSet = wire.NewSet(repository.NewBidangUrusanRepositoryImpl, wire.Bind(new(repository.BidangUrusanRepository), new(*repository.BidangUrusanRepositoryImpl)), service.NewBidangUrusanServiceImpl, wire.Bind(new(service.BidangUrusanService), new(*service.BidangUrusanServiceImpl)), controller.NewBidangUrusanControllerImpl, wire.Bind(new(controller.BidangUrusanController), new(*controller.BidangUrusanControllerImpl)))

var opdSet = wire.NewSet(repository.NewOpdRepositoryImpl, wire.Bind(new(repository.OpdRepository), new(*repository.OpdRepositoryImpl)), service.NewOpdServiceImpl, wire.Bind(new(service.OpdService), new(*service.OpdServiceImpl)), controller.NewOpdControllerImpl, wire.Bind(new(controller.OpdController), new(*controller.OpdControllerImpl)))

var pohonkinerjaSet = wire.NewSet(repository.NewPohonKinerjaRepositoryImpl, wire.Bind(new(repository.PohonKinerjaRepository), new(*repository.PohonKinerjaRepositoryImpl)), service.NewPohonKinerjaServiceImpl, wire.Bind(new(service.PohonKinerjaService), new(*service.PohonKinerjaServiceImpl)), controller.NewPohonKinerjaControllerImpl, wire.Bind(new(controller.PohonKinerjaController), new(*controller.PohonKinerjaControllerImpl)))

var sasarankotaSet = wire.NewSet(repository.NewSasaranKotaRepositoryImpl, wire.Bind(new(repository.SasaranKotaRepository), new(*repository.SasaranKotaRepositoryImpl)), service.NewSasaranKotaServiceImpl, wire.Bind(new(service.SasaranKotaService), new(*service.SasaranKotaServiceImpl)), controller.NewSasaranKotaControllerImpl, wire.Bind(new(controller.SasaranKotaController), new(*controller.SasaranKotaControllerImpl)))

var urusanSet = wire.NewSet(repository.NewUrusanRepositoryImpl, wire.Bind(new(repository.UrusanRepository), new(*repository.UrusanRepositoryImpl)), service.NewUrusanServiceImpl, wire.Bind(new(service.UrusanService), new(*service.UrusanServiceImpl)), controller.NewUrusanControllerImpl, wire.Bind(new(controller.UrusanController), new(*controller.UrusanControllerImpl)))

var domainspbeSet = wire.NewSet(repository.NewDomainSPBERepositoryImpl, wire.Bind(new(repository.DomainSPBERepository), new(*repository.DomainSPBERepositoryImpl)), service.NewDomainSPBEServiceImpl, wire.Bind(new(service.DomainSPBEService), new(*service.DomainSPBEServiceImpl)), controller.NewDomainSPBEControllerImpl, wire.Bind(new(controller.DomainSPBEController), new(*controller.DomainSPBEControllerImpl)))

var kebutuhanspbeSet = wire.NewSet(repository.NewKebutuhanSPBERepositoryImpl, wire.Bind(new(repository.KebutuhanSPBERepository), new(*repository.KebutuhanSPBERepositoryImpl)), service.NewKebutuhanSPBEServiceImpl, wire.Bind(new(service.KebutuhanSPBEService), new(*service.KebutuhanSPBEServiceImpl)), controller.NewKebutuhanSPBEControllerImpl, wire.Bind(new(controller.KebutuhanSPBEController), new(*controller.KebutuhanSPBEControllerImpl)))

var sasarankinerjaSet = wire.NewSet(repository.NewSasaranKinerjaPegawaiRepositoryImpl, wire.Bind(new(repository.SasaranKinerjaPegawaiRepository), new(*repository.SasaranKinerjaPegawaiRepositoryImpl)), service.NewSasaranKinerjaPegawaiServiceImpl, wire.Bind(new(service.SasaranKinerjaPegawaiService), new(*service.SasaranKinerjaPegawaiServiceImpl)), controller.NewSasaranKinerjaPegawaiControllerImpl, wire.Bind(new(controller.SasaranKinerjaPegawaiController), new(*controller.SasaranKinerjaPegawaiControllerImpl)))

var userSet = wire.NewSet(repository.NewUserRepositoryImpl, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)), service.NewUserServiceImpl, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)), controller.NewUserControllerImpl, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var rencanaPelaksanaanSet = wire.NewSet(repository.NewRencanaPelaksanaanRepositoryImpl, wire.Bind(new(repository.RencanaPelaksanaanRepository), new(*repository.RencanaPelaksanaanRepositoryImpl)), service.NewRencanaPelaksanaanServiceImpl, wire.Bind(new(service.RencanaPelaksanaanService), new(*service.RencanaPelaksanaanServiceImpl)), controller.NewRencanaPelaksanaanControllerImpl, wire.Bind(new(controller.RencanaPelaksanaanController), new(*controller.RencanaPelaksanaanControllerImpl)))
