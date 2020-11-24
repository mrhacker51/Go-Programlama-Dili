package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filename := "proxies.txt"
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	print := fmt.Println

	print(string(read))

}
