package image

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/internal/template"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/skip2/go-qrcode"
)

func (r *ImageResolver) GetReferralCodeImage(rw http.ResponseWriter, request *http.Request) {
	referralCode := chi.URLParam(request, "referral_code")
	content := fmt.Sprintf("%s/circuit/%s", os.Getenv("WEB_DOMAIN"), referralCode)
	circuitBytes, err := template.ReadFile("image/circuit.png")

	if err != nil {
		util.LogOnError("API037", "Error reading image file", err)
		log.Printf("API037: ReferralCode=%v", referralCode)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	circuitBytesReader := bytes.NewReader(circuitBytes)
	circuitImageData, err := png.Decode(circuitBytesReader)

	if err != nil {
		util.LogOnError("API038", "Error decoding png", err)
		log.Printf("API038: ReferralCode=%v", referralCode)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	imageData := image.NewRGBA(image.Rect(0, 0, circuitImageData.Bounds().Dx(), circuitImageData.Bounds().Dy()))
	draw.Draw(imageData, imageData.Bounds(), circuitImageData, image.Point{X: 0, Y: 0}, draw.Src)
	qr, err := qrcode.New(content, qrcode.Medium)
	
	if err != nil {
		util.LogOnError("API039", "Error creating QR code", err)
		log.Printf("API039: Content=%v", content)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	qr.DisableBorder = true
	foregroundColor := color.RGBA{255, 255, 255, 255}
	bitmap := qr.Bitmap()
	size := 512
	border := int((imageData.Bounds().Dy() - size) / 2)
	xOffset := imageData.Bounds().Dx() - border - size
	yOffset := border
	modulesPerPixel := float64(len(bitmap)) / float64(size)

	for y := 0; y < size; y++ {
		y2 := int(float64(y) * modulesPerPixel)
		for x := 0; x < size; x++ {
			x2 := int(float64(x) * modulesPerPixel)

			v := bitmap[y2][x2]

			if v {
				imageData.SetRGBA(x + xOffset, y + yOffset, foregroundColor)
			}
		}
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, imageData)

	if err != nil {
		util.LogOnError("API040", "Error encoding png", err)
		log.Printf("API040: Content=%v", content)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Write(buffer.Bytes())
}
