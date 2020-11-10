package main

import (
	"fmt"
)

func main() {

	a := [3]int{}

	a[0] = 10
	a[1] = 20
	a[2] = 30

	fmt.Println(a[0], a[1], a[2])

	//

	var colors [3]string

	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "DarkGreen"

	fmt.Println(colors)

	fmt.Println("\n", colors[0], "\n", colors[1], "\n", colors[2])

	//

	color := [3]string{}

	color[0] = "Yellow"
	color[1] = "White"
	color[2] = "Black"

	fmt.Println(color)
	fmt.Println("\n", color[0], "\n", color[1], "\n", color[2])
}
