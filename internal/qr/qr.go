package qrc

import (
	"github.com/skip2/go-qrcode"
)

func NewQr(text string) string {
	qr, _ := qrcode.New(text, qrcode.Medium)
	return qr.ToSmallString(false)
}
