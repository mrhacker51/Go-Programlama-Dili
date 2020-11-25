package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	text := []byte("GolangProgramming")                     // Cipher Text .. No Encoding
	encodeCipher := base64.StdEncoding.EncodeToString(text) // Encode ...
	fmt.Printf("%s\n", encodeCipher)                        // Print ...

	// DECODED

	decodeCipher, err := base64.StdEncoding.DecodeString(encodeCipher)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", decodeCipher)

	fmt.Println(string(decodeCipher))
}
