package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("hello golang")

	defer f.Close()

}

// FİLE CREATE AND FİLE WRİTİNG
