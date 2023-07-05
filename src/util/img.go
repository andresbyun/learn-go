package util

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type Img struct {
	Width, Height int
	Name          string
}

func (this Img) GenerateImage() {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{this.Width, this.Height}
	outputImg := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	// Set color for each pixel.
	for x := 0; x < this.Width; x++ {
		for y := 0; y < this.Height; y++ {
			// TODO: Convert values in matrix to color
			outputImg.Set(x, y, color.White)
		}
	}

	// Change directory to output directory
	os.Chdir("../output")

	if !strings.HasSuffix(this.Name, ".png") {
		this.Name += ".png"
	}
	filename := this.Name

	// Generate the image
	f, _ := os.Create(this.Name)
	png.Encode(f, outputImg)
	fmt.Println("Created image: " + filename)
}
