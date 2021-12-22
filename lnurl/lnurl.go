package lnurl

import (
	"encoding/base64"
	"fmt"

	golnurl "github.com/fiatjaf/go-lnurl"
	qrcode "github.com/skip2/go-qrcode"
)

func RandomK1() string {
	return golnurl.RandomK1()
}

func GenerateQrCode(lnUrl string) (string, error) {
	qr, err := qrcode.Encode(fmt.Sprintf("lightning:%s", lnUrl), qrcode.Highest, 256)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(qr), nil
}
