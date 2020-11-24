package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("proxies.txt", os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytefile := []byte("Okey Golang ....")

	newFile, err := file.Write(bytefile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New %d Byte", newFile)

}
