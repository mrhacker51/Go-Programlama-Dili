package main

import "fmt"

func main() {

	result_male, result_female := swap("James", "Anna")
	print := fmt.Println
	print(result_male, result_female)
}

func swap(x, y string) (string, string) {
	return x, y
}

// (string,string)  
