package main

import (
	"fmt"
	"learn-go/util"
	"learn-go/util/linAlgebra"
	"learn-go/util/worldObj"
	"math/rand"
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

	// worldCollection
	test := []worldObj.WorldObj{}

	// Add a black sphere in the center
	black := linAlgebra.Vec3{X: 0, Y: 0, Z: 0}
	radius := 0.25
	pos := linAlgebra.Vec3{X: 0, Y: 0, Z: -1}
	test = append(test, worldObj.Sphere{Position: pos, Color: black, Radius: radius})

	// Add randomized spheres
	randomizeSpheres(viewportWidth, viewportHeight, &test, 5)

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

			imgMat[i][j] = getColor(r, test)
		}
	}

	// Generate the output image
	output := util.Img{Width: Width, Height: Height, Name: "test.png"}
	output.GenerateImage(imgMat)
}

// Get the color of the ray
func getColor(r util.Ray, world []worldObj.WorldObj) linAlgebra.Vec3 {
	// TEST multiple spheres
	for i := 0; i < len(world); i++ {
		if world[i].HitIntersect(r) {
			return world[i].GetColor()
		}
	}

	unitVec := r.Dir.Normalize()
	t := 0.5 * (unitVec.Y + 1.0)

	lerpWhite := linAlgebra.Vec3{X: 1, Y: 1, Z: 1}.Mult(1.0 - t)
	lerpBlue := linAlgebra.Vec3{X: 0.5, Y: 0.7, Z: 1}.Mult(t)

	return lerpWhite.Add(lerpBlue)
}

// Create an array of n amount of randomized spheres
func randomizeSpheres(width float64, height float64, world *[]worldObj.WorldObj, amount uint8) {
	for i := 0; i < int(amount); i++ {
		radius := rand.Float64() * 0.4 // radius of range [0.0, 0.4)

		color := linAlgebra.Vec3{X: rand.Float64(), Y: rand.Float64(), Z: rand.Float64()}
		pos := linAlgebra.Vec3{
			X: (rand.Float64() * (width)) - (width / 2),
			Y: (rand.Float64() * (height)) - (height / 2),
			Z: float64(-1),
		}

		randSp := worldObj.Sphere{Position: pos, Color: color, Radius: radius}
		*world = append(*world, randSp)
	}
}
