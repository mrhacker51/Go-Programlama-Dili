package main

import "fmt"

func main() {
	result := user_input("root", "root")
	fmt.Println(result)

	n_result := user_function(10, 20)

	fmt.Println(n_result)
}


//

func user_input(a, b string) string {
	user_input_get := a
	password_input_get := b

	return (user_input_get + password_input_get)
}

//

func user_function(a, b int) int {
	return ((a * 2) + (b * 2))
}
