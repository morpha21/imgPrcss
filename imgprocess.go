package main

import (
	"image"
	"log"
	"net/http"
	"os"

	_ "image/jpeg"
	"image/png"
)

func main() {
	resp, err := http.Get("https://qph.fs.quoracdn.net/main-thumb-78855837-200-lrecovmpceibxumvmwbsjolvvjhntddv.jpeg")
	checkError(err)
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	checkError(err)
	log.Printf("Image Type: %T", img)

	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += 2 {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	f, err := os.Create("newimage.png")
	checkError(err)
	defer f.Close()

	checkError(png.Encode(f, grayImg))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
