package main

import (
	"fmt"
	"learn-go/util"
	"learn-go/util/linAlgebra"
	"learn-go/util/worldObj"
)

func main() {
	// Set the image's height and width
	fmt.Println("Initializing!")
	Height := 720
	Width := 1280

	// Variables
	aspectRatio := float64(Width) / float64(Height)
	viewportHeight := float64(2.0)
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := linAlgebra.Vec3{X: 0, Y: 0, Z: 0}
	horizontal := linAlgebra.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := linAlgebra.Vec3{X: 0, Y: viewportHeight, Z: 0}
	lowerLeft := ((origin.Sub(horizontal.Div(2))).Sub(vertical.Div(2))).Sub(linAlgebra.Vec3{X: 0, Y: 0, Z: focalLength})

	// Initialize the 2D array that'll contain the image values
	imgMat := make([][]linAlgebra.Vec3, Height)
	for i := 0; i < Height; i++ {
		imgMat[i] = make([]linAlgebra.Vec3, Width)

		for j := 0; j < Width; j++ {
			u := float64(j) / float64(Width-1)
			uVec := horizontal.Mult(u)

			v := float64(i) / float64(Height-1)
			vVec := vertical.Mult(v)

			direction := (uVec.Add(vVec)).Add(lowerLeft)
			r := util.Ray{Orig: origin, Dir: direction}

			imgMat[i][j] = getColor(r)
		}
	}

	// Generate the output image
	output := util.Img{Width: Width, Height: Height, Name: "test.png"}
	output.GenerateImage(imgMat)
}

// Get the color of the ray
func getColor(r util.Ray) linAlgebra.Vec3 {
	// TEST hit with a black sphere at {X:0, Y:-1, Z:-1} with radius 0.5
	black := linAlgebra.Vec3{X: 0, Y: 0, Z: 0}
	radius := 0.5
	pos := linAlgebra.Vec3{X: 0, Y: 0, Z: -1}

	testSphere := worldObj.Sphere{Position: pos, Color: black, Radius: radius}
	if testSphere.HitIntersect(r) {
		return testSphere.GetColor()
	}
	// END TEST

	unitVec := r.Dir.Normalize()
	t := 0.5 * (unitVec.Y + 1.0)

	lerpWhite := linAlgebra.Vec3{X: 1, Y: 1, Z: 1}.Mult(t)
	lerpBlue := linAlgebra.Vec3{X: 0.5, Y: 0.7, Z: 1}.Mult(1.0 - t)

	return lerpWhite.Add(lerpBlue)
}
