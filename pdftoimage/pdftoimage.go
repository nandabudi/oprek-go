package main

import (
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {

	pdfName := "ai.pdf"

	if err := ConvertPdfToJpg(pdfName); err != nil {
		log.Fatal(err)
	}
}

// ConvertPdfToJpg is
func ConvertPdfToJpg(pdfName string) error {

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Make sure our image is high quality
	if err := mw.SetResolution(300, 300); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
		return err
	}

	// Set any compression (100 = max quality)
	if err := mw.SetCompressionQuality(95); err != nil {
		return err
	}

	// Convert into JPG
	if err := mw.SetFormat("jpg"); err != nil {
		return err
	}

	// Save File
	return mw.WriteImages("files/test.jpg", false)
}
