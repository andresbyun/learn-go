package util

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	linAlgebra "learn-go/util/lin-algebra"
	"os"
	"strings"
)

type Img struct {
	Width, Height int
	Name          string
}

// Function that creates the output image
func (this Img) GenerateImage() {
	colorMat := this.raytrace()

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

// Function to raytrace
func (this Img) raytrace() [][]linAlgebra.Vec3 {
	// Initialize the 2D array that'll contain the image values
	imgMat := make([][]linAlgebra.Vec3, this.Height)
	for i := 0; i < this.Height; i++ {
		imgMat[i] = make([]linAlgebra.Vec3, this.Width)
		// Fill the background with white
		for j := 0; j < this.Width; j++ {
			imgMat[i][j] = linAlgebra.Vec3{X: 1, Y: 1, Z: 1}
		}
	}

	return imgMat
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
