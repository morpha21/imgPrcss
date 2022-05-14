package main

import (
	"image"
	"log"
	"os"

	_ "image/jpeg"
	"image/png"
)

func main() {
	//imports image
	imageFile, err := os.Open(os.Args[1])
	defer imageFile.Close()
	checkError(err)

	//prepares image for manipulation
	img, _, err := image.Decode(imageFile)
	checkError(err)

	//creates a gray image with tiny black stripes
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += 2 {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	//I'm trying to create a colored image with black stripes, but I'm kinda failing
	imagem := image.NewRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X + 1; x < img.Bounds().Max.X; x += 2 {
			imagem.Set(x, y, img.At(x, y))
		}
		for x := (img.Bounds().Min.X); x < img.Bounds().Max.X; x += 2 {
			imagem.Set(x, y, grayImg.At(x, y))
		}
	}

	//exports the gray image to the 1newimage.png file
	f, err := os.Create("1newimage.png")
	checkError(err)
	checkError(png.Encode(f, grayImg))

	defer f.Close()

	//exports the colored image to the 2newimage.png file
	f, err = os.Create("2newimage.png")
	checkError(err)
	checkError(png.Encode(f, imagem))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
