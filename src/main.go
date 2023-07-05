package main

import (
	"fmt"
	"learn-go/util"
)

func main() {
	fmt.Println("Hello, World!")
	Height := 400
	Width := 400

	output := util.Img{Width: Width, Height: Height, Name: "test.png"}
	output.GenerateImage()
}
