package main

import "fmt"

func main() {
	variable := sum(1, 2, 3, 4)
	fmt.Println(variable)

	for i := 0; i <= 3; i++ {
		fmt.Println(variable[i])
	}

	ex := space("Hey", "Hey2", "Hey3", "Hey4")
	fmt.Printf("Variable : %s", ex)
}

func sum(sayi ...int) []int {
	result := []int{}

	for _, number := range sayi {
		result = append(result, number)
	}

	return result
}

//

func space(speak ...string) []string {
	result := []string{}

	for _, rr := range speak {
		result = append(result, rr)
	}
	return result
}
