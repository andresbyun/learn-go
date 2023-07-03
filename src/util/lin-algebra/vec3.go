package linAlgebra

import (
	"math"
)

// very simple vector class
type Vec3 struct {
	X, Y, Z float64
}

// add vector b to vector a (i.e. a + b)
func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

// subtract vector b to vector a (i.e. a - b)
func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

// dot product of vectors a and b
func (a Vec3) Dot(b Vec3) float64 {
	return ((a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z))
}

// return the magnitute/length of the vector
func (a Vec3) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// normalize vector
func (a Vec3) Normalize() Vec3 {
	len := a.Length()
	return Vec3{
		a.X / len,
		a.Y / len,
		a.Z / len,
	}
}
