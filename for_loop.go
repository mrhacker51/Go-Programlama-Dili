package main

import "fmt"

func main() {
	for i := 0; i < 50; i++ {
		fmt.Printf("Greater : \n%d\n", i)
	}

	for x := 100; x >= 0; x-- {
		fmt.Printf("%d\n", x)
	}

	numeric := 100

	for numeric <= 200 {
		numeric++
		if numeric%2 == 0 {
			fmt.Printf("Number : %d\n", numeric)
		}

	}
}
