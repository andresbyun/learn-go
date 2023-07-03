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

// magnitute/length of the vector
func (a Vec3) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// scalar multiplication
func (a Vec3) Mult(t float64) Vec3 {
	return Vec3{
		a.X * t,
		a.Y * t,
		a.Z * t,
	}
}

// scalar division
func (a Vec3) Div(t float64) Vec3 {
	s := 1 / t
	return a.Mult(s)
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

// cross product
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		(a.Y * b.Z) - (b.Y * a.Z),
		(a.Z * b.X) - (b.Z * a.X),
		(a.X * b.Y) - (b.X * a.Y),
	}
}
