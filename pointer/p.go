package main

// POINTER

import (
	f "fmt"
)

func main() {
	p := "Pointer..."

	f.Println(p)  // "Pointer..."
	f.Println(&p) // p nin addresi
	pp := &p
	f.Println(pp) // p'nin addresini aldı

	f.Println(*pp) // "Pointer..."  p'nin degerini aldı

}
