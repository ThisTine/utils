// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ThisTine/utils/controllers"
	"github.com/ThisTine/utils/services"
)

// Injectors from wire.go:

func initQrCodeAPI() controllers.QrController {
	iQrService := services.ProvideQrService()
	qrController := controllers.ProvideQrController(iQrService)
	return qrController
}
