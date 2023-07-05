package main

import (
	"fmt"
	"learn-go/util"
	linAlgebra "learn-go/util/lin-algebra"
)

func main() {
	// Set the image's height and width
	fmt.Println("Initializing!")
	Height := 400
	Width := 400

	// Initialize the 2D array that'll contain the image values
	imgMat := make([][]linAlgebra.Vec3, Height)
	for i := 0; i < Height; i++ {
		imgMat[i] = make([]linAlgebra.Vec3, Width)
		// Fill the background with white
		for j := 0; j < Width; j++ {
			imgMat[i][j] = linAlgebra.Vec3{X: 1, Y: 1, Z: 1}
		}
	}

	// Generate the output image
	output := util.Img{Width: Width, Height: Height, Name: "test.png"}
	output.GenerateImage(imgMat)
}
