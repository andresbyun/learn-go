package main

import (
	"fmt"
	"learn-go/util"
)

func main() {
	// Set the image's height and width
	fmt.Println("Initializing!")
	Height := 720
	Width := 1280

	// Generate the output image
	output := util.Img{Width: Width, Height: Height, Name: "test.png"}
	output.GenerateImage()
}
