package util

import (
	"learn-go/util/linAlgebra"
)

type Ray struct {
	Orig, Dir linAlgebra.Vec3
}

// get the position of the ray at t
func (r Ray) At(t float64) linAlgebra.Vec3 {
	return r.Orig.Add((r.Dir.Mult(t)))
}
