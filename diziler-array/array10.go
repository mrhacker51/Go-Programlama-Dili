package main

import (
	"fmt"
)

func main() {

	var my_array = []string{}

	my_array = append(my_array, "PHP")
	fmt.Println(my_array)

	my_array = append(my_array, "Golang")

	fmt.Println(my_array)

	//

	var arry []string

	fmt.Println(arry)

	arry = append(arry, "CSharp")

	fmt.Println(arry)

	//

	for i := 0; i <= len(arry)-1; i++ {
		fmt.Println(arry[i])
	}

	//

	for _, values := range arry {
		fmt.Println(values)
	}
	//

}
