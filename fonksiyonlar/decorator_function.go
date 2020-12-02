package main

// decorator function

import (
	"fmt"
)

func main() {
	birfunc(decoratorfunction)
	ikifunc()
}

func decoratorfunction() {
	fmt.Println("Connecting ....")
}

func birfunc(a func()) { // içeriye bir a adında degisken alsın .. bunuda fonksiyon olarak alsın
	fmt.Println("Failed ....")
	a() // icerisinde aldıgı degiskeni parametreyi burda calistiriyorum veya islemlerden öncede cagirabilirim tamamen programcıya kalmıs birsey
}

//

func decoratornew() {
	for a := 0; a <= 10; a++ {
		fmt.Printf("Value : %d\n", a)
	}
}

func ikifunc() {
	fmt.Println("***************")
	decoratornew()
}
