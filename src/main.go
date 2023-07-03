package main

import (
	"fmt"
	linAlgebra "learn-go/util/lin-algebra"
)

func main() {
	m := map[string]float64{
		"x": 0.5,
		"y": 0.5,
		"z": 0,
	}
	v := linAlgebra.Vec3{X: m["x"], Y: m["y"], Z: m["z"]}

	fmt.Printf("Testing values: %v\n", v)
	fmt.Printf("Testing length: %v\n", v.Length())
	fmt.Printf("Testing normal: %v\n", v.Normalize())
}
