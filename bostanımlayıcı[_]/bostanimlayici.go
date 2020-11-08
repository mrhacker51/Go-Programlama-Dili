package main

import "fmt"

// _   = Bos Tanımlayıcı ...
// hata , veri   = method() --- _ ,veri = method()

func main() {

	var _, x, _ int = 0, 10, 0
	fmt.Println(x)

	bos()
}

func bos() {
	_, x, _ := 0, 20, 0
	fmt.Println(x)
}
