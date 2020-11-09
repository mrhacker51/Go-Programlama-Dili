package main

import "fmt"

func main() {

	x := 26

	if x%2 == 1 {
		fmt.Println("Tek ...")
	} else {
		fmt.Println("Cift ...")
	}

	if y := 55; y%2 == 1 {
		fmt.Println("Y : Tek ...")
	} else {
		fmt.Println("Y : Cift")
	}

	var name string = "Go"
	var surname string = "Lang"

	if name == "Go" && surname == "Lang" {
		name = "NewGo"
		surname = "NewLang"
		fmt.Println("Success ....")
	} else {
		fmt.Println("Name and Surname Not Found ....")
		fmt.Println("Program closed ....")
	}
	fmt.Println("[+] New name and Surname variable")
	fmt.Println(name, surname)

}
