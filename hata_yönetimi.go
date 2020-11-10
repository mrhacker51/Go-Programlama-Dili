package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	my, err := os.Create("test.txt")
	// my , err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(my)
	fmt.Println(my.Name())

	//

	my_errors := errors.New("Bu Bir Hatadir ....")

	fmt.Println(my_errors)

	//

	my_errors_ := errors.New("Program Error  .... !")

	fmt.Println(my_errors_)

	//

	var my_e = errors.New("Not Found .... !")

	fmt.Println(my_e)

	//

	f, _ := os.Open("test.txt")

	fmt.Println(f.Name())

	//

	ff := errors.New("Attack ERROR ... ")

	fmt.Println(ff.Error())

	//

	_, newerror := os.Open("test.txt")

	if newerror != nil {
		log.Fatal("Error ...", newerror)
	}

}
