package main

import "fmt"

type shape interface {
	draw()
}

type rect struct {
	name string
	height int
}

func (re rect) draw() {
	var height = re.height

	if height > 33 {
		height = 33
	}

	drawRightAngledTriangle(height)
}

func drawRightAngledTriangle(max int) {
	for i := 1; i <= max; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func main() {
	fmt.Println("Please choose a height for your lovely triangle : ")

	var inputSize int
	fmt.Scanln(&inputSize)

	var myShape = rect{"right-angled", inputSize}
	myShape.draw()
}
