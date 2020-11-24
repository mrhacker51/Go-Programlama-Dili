package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("proxy.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
