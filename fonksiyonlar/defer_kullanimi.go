package main

import "fmt"

func main() {

  // defer = erteleme ... 
  
	// fmt.Println("Hello Golang")

	// defer fmt.Println("Are You Golang ??? ...")

	// fmt.Println("You ....?")

	/*
		result =

		Hello Golang
		You ....?
		Are You Golang ??? ...

	*/
	//

	i := 0
	defer fmt.Println("Program Closed ....")

	for i <= 20 {
		fmt.Println("Program Opened ...")
		i++
	}

	for j := 5; j >= 0; j-- {
		defer fmt.Println("Program Error ...")
	}

	defer (func() {
		fmt.Println("Function [[[]]]..... !!")
	})()

	(func() {
		defer fmt.Println("Function [[[]]]..... !!")
	})()

}
