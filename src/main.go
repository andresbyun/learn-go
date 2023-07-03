package main

import (
	"fmt"
	"learn-go/util"
	linAlgebra "learn-go/util/lin-algebra"
)

func main() {
	m := map[string]float64{
		"x": 0.5,
		"y": 0.5,
		"z": 0,
	}
	v := linAlgebra.Vec3{X: m["x"], Y: m["y"], Z: m["z"]}

	fmt.Printf("Testing length: %v\n", v.Length())
	fmt.Printf("Testing normal: %v\n", v.Normalize())

	testVar := util.Ray{
		Orig: linAlgebra.Vec3{X: 0, Y: 0, Z: 0},
		Dir:  v,
	}
	fmt.Printf("Testing values: %v\n", testVar)

}
