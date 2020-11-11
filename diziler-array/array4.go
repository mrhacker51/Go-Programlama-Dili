package main

import (
	"fmt"
)

func main() {

	var numbers = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(numbers)
	fmt.Printf("Number of numbers : \n%d\n", len(numbers))

	for _, values := range numbers {
		fmt.Println("Values  : ", values)
	}
}
