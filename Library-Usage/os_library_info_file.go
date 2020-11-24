package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileinfo *os.FileInfo
	err      error
)

func main() {
	fileinfo, err := os.Stat("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name : ", fileinfo.Name())
	fmt.Println("File Size : ", fileinfo.Size())
	fmt.Println("File Mode : ", fileinfo.Mode())
	fmt.Println("File ModTime : ", fileinfo.ModTime())
	fmt.Println("File Sys : ", fileinfo.Sys())
	fmt.Println("File IsDir : ", fileinfo.IsDir())

	if fileinfo.IsDir() == false {
		fmt.Println("This is not File")
	} else {
		fmt.Println("File Found ...")
	}
}
