package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := []byte("GolangProgramming")                     // Cipher Text .. No Encoding
	encodeCipher := base64.StdEncoding.EncodeToString(text) // Encode ...
	fmt.Printf("%s", encodeCipher)                          // Print ...
}
