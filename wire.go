//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ThisTine/utils/controllers"
	"github.com/ThisTine/utils/services"
	"github.com/google/wire"
)

func initQrCodeAPI() controllers.QrController {
	wire.Build(services.ProvideQrService, controllers.ProvideQrController)

	return controllers.QrController{}
}
