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

func (out Img) GenerateImage() {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{out.Width, out.Height}
	output := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	// Set color for each pixel.
	for x := 0; x < out.Width; x++ {
		for y := 0; y < out.Height; y++ {
			// TODO: Convert values in matrix to color
			output.Set(x, y, color.White)
		}
	}

	// Change directory to outside src
	os.Chdir("..")

	if !strings.HasSuffix(out.Name, ".png") {
		out.Name += ".png"
	}
	filename := out.Name

	// Generate the image
	f, _ := os.Create(out.Name)
	png.Encode(f, output)
	fmt.Println("Created image: " + filename)
}
