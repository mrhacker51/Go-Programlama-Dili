package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("proxies.txt", os.O_WRONLY|os.O_CREATE|os.O_RDONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte("Which Challenge ?")

	newFile, err := f.Write(bytes)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Byteee %d", newFile)

	readfile, err := ioutil.ReadFile("proxies.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", string(readfile))

}
