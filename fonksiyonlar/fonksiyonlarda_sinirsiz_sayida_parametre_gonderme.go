package main

import "fmt"

func main() {
	a := sum(2, 3, 4, 5)
	fmt.Println(a)

	messages("Hello")

}

func sum(sayi ...int) int {
	result := 0

	for _, values := range sayi {
		result += values
	}

	return result
}

//

func messages(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}
