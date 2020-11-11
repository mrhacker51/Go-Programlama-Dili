package main

import "fmt"

func main() {

	var space [3]string

	space[0] = "ElonMusk"
	space[1] = "Nasa"
	space[2] = "Tesla"

	fmt.Println("Count : ", len(space))
	fmt.Println("Capacity : ", cap(space))
	fmt.Println(space)

	fmt.Println(space[0], space[1], space[2])
}
