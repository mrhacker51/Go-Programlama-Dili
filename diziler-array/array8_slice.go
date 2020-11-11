package main

import "fmt"

func main() {

	myArray := [...]int{51, 34, 55, 42}
	fmt.Println(myArray[:])

	//

	myArray2 := [...]int{10, 20, 30}

	mySlice := myArray2[:]
	fmt.Println(mySlice)

	//

	mySlice[0] = 25
	fmt.Println(mySlice)
	//
	fmt.Println(myArray2)

	//

	myStringArray := [...]string{"PHP", "Golang", "Python", "C", "Assembly", "JavaScript"}

	fmt.Println(myStringArray[:])

	//

	result := myStringArray[:]

	fmt.Println(result)

	//

	result[3] = "NodeJs"

	fmt.Println(result)
}
