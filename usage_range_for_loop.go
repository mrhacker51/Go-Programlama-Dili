package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	for index, variable := range numbers {
		fmt.Println("İndex : ", index, "is", "Variable", variable)
	}

	strings_ := [...]string{"PHP", "Python", "C", "Golang"}

	for n, y := range strings_ {
		fmt.Println("İndex : ", n, "İs", "Variable : ", y)
	}

}
