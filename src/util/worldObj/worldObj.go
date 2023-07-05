package worldObj

import (
	"learn-go/util"
	"learn-go/util/linAlgebra"
)

type WorldObj interface {
	GetPosition() linAlgebra.Vec3
	GetColor() linAlgebra.Vec3
	HitIntersect(r util.Ray) bool
}

// Spheres
type Sphere struct {
	Position linAlgebra.Vec3
	Color    linAlgebra.Vec3
	Radius   float64
}

func (s Sphere) GetPosition() linAlgebra.Vec3 {
	return s.Position
}

func (s Sphere) GetColor() linAlgebra.Vec3 {
	return s.Color
}

func (s Sphere) HitIntersect(r util.Ray) bool {
	oc := r.Orig.Sub(s.Position)
	a := r.Dir.Dot(r.Dir)
	b := 2.0 * oc.Dot(r.Dir)
	c := oc.Dot(oc) - (float64(s.Radius) * float64(s.Radius))

	discriminant := (b * b) - 4*(a*c)

	return (discriminant > 0)
}
