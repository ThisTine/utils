package controllers

import (
	"github.com/ThisTine/utils/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QrController struct {
	QrService services.IQrService
}

func ProvideQrController(qrService services.IQrService) QrController {
	return QrController{QrService: qrService}
}

// MakeQR godoc
// @Summary Generate QRCode
// @Schemes
// @Description Generate QRCode
// @Tags QRCodeAPI
// @Accept json
// @Produce image/png
// @Param text query string true "Text to generate QRCode"
// @Param size query int false "Size of QRCode (in pixel)"
// @Param disableBorder query bool false "Border of QRCode"
// @Success 200 {string} image/png "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /qr/v1/make [get]
func (q *QrController) MakeQR(c *gin.Context) {
	//q.QrService.MakeQR()
	text := c.Query("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	size := c.Query("size")
	if size == "" {
		size = "256"
	}

	disableBorder := c.Query("disableBorder")
	if disableBorder == "" {
		disableBorder = "false"
	}

	disableBorderBool, err := strconv.ParseBool(disableBorder)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if sizeInt, err := strconv.Atoi(size); err != nil || sizeInt < 0 || sizeInt > 1024 {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	} else {
		if img, err := q.QrService.MakeQR(text, sizeInt, disableBorderBool); err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			c.Header("Content-Type", "image/png")
			c.Data(http.StatusOK, "image/png", img)

		}
	}
}
