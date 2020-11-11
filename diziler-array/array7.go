package main

import "fmt"

func main() {

	var space [3]string

	space[0] = "Mars"
	space[1] = "Venüs"
	space[2] = "Neptune"

	i := 0
	for i <= cap(space)-1 {
		fmt.Println(space[i])
		i++
	}

	for j := 0; j <= cap(space)-1; j++ {
		fmt.Println(space[j])
	}

}
