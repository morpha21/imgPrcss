package main

import (
	"image"
	"log"
	"os"

	_ "image/jpeg"
	"image/png"
)

func main() {

	imageFile, err := os.Open(os.Args[1])
	defer imageFile.Close()
	checkError(err)

	img, _, err := image.Decode(imageFile)
	checkError(err)

	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += 2 {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	imagem := image.NewRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X + 1; x < img.Bounds().Max.X; x += 2 {
			imagem.Set(x, y, img.At(x, y))
		}
		for x := (img.Bounds().Min.X); x < img.Bounds().Max.X; x += 2 {
			imagem.Set(x, y, grayImg.At(x, y))
		}
	}

	f, err := os.Create("1newimage.png")
	checkError(err)
	f2, err := os.Create("2newimage.png")
	checkError(err)
	defer f.Close()

	checkError(png.Encode(f, grayImg))
	checkError(png.Encode(f2, imagem))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
