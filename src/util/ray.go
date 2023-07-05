package util

import (
	linAlgebra "learn-go/util/lin-algebra"
)

type Ray struct {
	Orig, Dir linAlgebra.Vec3
}

// get the position of the ray at t
func (r Ray) At(t float64) linAlgebra.Vec3 {
	return r.Orig.Add((r.Dir.Mult(t)))
}

// Get the color of the ray
func (r Ray) getColor() linAlgebra.Vec3 {
	unitVec := r.Dir.Normalize()
	t := 0.5 * (unitVec.Y + 1.0)

	lerpWhite := linAlgebra.Vec3{X: 1, Y: 1, Z: 1}.Mult(t)
	lerpBlue := linAlgebra.Vec3{X: 0.5, Y: 0.7, Z: 1}.Mult(1.0 - t)

	return lerpWhite.Add(lerpBlue)
}
