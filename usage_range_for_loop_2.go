package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range numbers {
		if i%2 == 0 {
			fmt.Println("Cift : ", i)
		} else {
			fmt.Println("Tek ", i)
		}

	}

	language := [...]string{"PHP", "Golang", "Python", "JavaScript", "NodeJs", "Assembly", "Csharp", "C", "C++", "Dart", "Kotlin", "Java"}

	for _, y := range language {
		fmt.Println("Language : ", y)
	}

}
