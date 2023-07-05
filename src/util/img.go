package util

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"learn-go/util/linAlgebra"
	"os"
	"strings"
)

type Img struct {
	Width, Height int
	Name          string
}

// Function that creates the output image
func (this Img) GenerateImage(colorMat [][]linAlgebra.Vec3) {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{this.Width, this.Height}
	outputImg := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	// Set color for each pixel.
	for i := 0; i < this.Height; i++ {
		for j := 0; j < this.Width; j++ {
			curr := Vec3ToRGBA(colorMat[i][j])
			outputImg.Set(j, i, curr)
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

// Convert Vec3 to RGBA
func Vec3ToRGBA(v linAlgebra.Vec3) color.RGBA {
	uColor := map[uint8]uint8{
		'r': uint8(v.X * 0xFF),
		'g': uint8(v.Y * 0xFF),
		'b': uint8(v.Z * 0xFF),
	}

	return color.RGBA{uColor['r'], uColor['g'], uColor['b'], 0xFF}
}
