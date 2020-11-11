package main

import (
	"fmt"
)

func main() {

	var numbers = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(numbers)
	fmt.Printf("Number of numbers : %d", len(numbers))
}


// [...] Otomatik Size ...
// []int{} 
