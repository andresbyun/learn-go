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
	// Variables
	aspectRatio := this.Width / this.Height
	viewportHeight := float64(2.0)
	viewportWidth := float64(aspectRatio) * viewportHeight
	focalLength := 1.0

	origin := linAlgebra.Vec3{X: 0, Y: 0, Z: 0}
	horizontal := linAlgebra.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := linAlgebra.Vec3{X: 0, Y: viewportHeight, Z: 0}

	// TODO: Rewrite this so it's actually behaving the way it's supposed to
	lowerLeft := ((origin.Sub(horizontal.Div(2))).Sub(vertical.Div(2))).Sub(linAlgebra.Vec3{X: 0, Y: 0, Z: focalLength})

	// Initialize the 2D array that'll contain the image values
	imgMat := make([][]linAlgebra.Vec3, this.Height)
	for i := 0; i < this.Height; i++ {
		imgMat[i] = make([]linAlgebra.Vec3, this.Width)

		for j := 0; j < this.Width; j++ {
			u := float64(j) / float64(this.Width-1)
			uVec := horizontal.Mult(u)

			v := float64(i) / float64(this.Height-1)
			vVec := vertical.Mult(v)

			temp := (uVec.Add(vVec)).Add(lowerLeft)
			r := Ray{Orig: origin, Dir: temp}

			imgMat[i][j] = r.getColor()
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
