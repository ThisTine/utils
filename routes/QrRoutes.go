package routes

import (
	"github.com/ThisTine/utils/controllers"
	"github.com/gin-gonic/gin"
)

func QrRoute(r *gin.RouterGroup, qr controllers.QrController) {
	qrAPI := r.Group("/qr")
	{
		v1 := qrAPI.Group("/v1")
		{
			v1.GET("/make", qr.MakeQR)
		}
	}
}
