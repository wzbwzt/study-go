package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file_path := "./证件照.jpg"

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	//一寸证件照 像素大小：295px*413px
	m := resize.Resize(295, 0, img, resize.NearestNeighbor)

	out, err := os.Create("react.NearestNeighbor.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
