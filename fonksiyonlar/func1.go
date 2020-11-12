package main

import "fmt"

func main() {
	result := Sum(5, 5)
	fmt.Println(result)
}

func Sum(a int, b int) int {
	return (a + b)
}
