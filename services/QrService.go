package services

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"image/png"
)

type QrService struct {
}

func ProvideQrService() QrService {
	return QrService{}
}

func (q *QrService) MakeQR(text string, size int, disableBorder bool) (image []byte, err error) {
	if qrc, err := qrcode.New(text, qrcode.High); err != nil {
		return nil, err
	} else {
		qrc.DisableBorder = disableBorder
		img := qrc.Image(size)
		buf := new(bytes.Buffer)
		err := png.Encode(buf, img)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}
}
