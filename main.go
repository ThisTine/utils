package main

import (
	"github.com/ThisTine/utils/docs"
	"github.com/ThisTine/utils/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Thistine's Utils API
// @version 1.0
// @description This contains a set of utilities API
// @contact.name Thistine
// @contact.url https://thistine.com
// @contact.email tine@thistine.com
// BasePath: /api/v1
func main() {
	r := gin.Default()

	api := r.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"

	// Routes
	qr := initQrCodeAPI()
	routes.QrRoute(api, qr)

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
