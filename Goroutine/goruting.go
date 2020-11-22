// Thread gibi fakat tam thread olmayan işlem parçacıkları ..

package main

import (
	"fmt"
	"time"
)

func gorutin() {
	for i := 0; i <= 5; i++ {
		fmt.Println("GORUTİN FUNC 1 ")
	}
}

func gorutinF() {
	for i := 0; i <= 7; i++ {
		fmt.Println("GORUTİN FUNC 2 ")
	}
}

func gorutinFF() {
	for i := 0; i <= 9; i++ {
		fmt.Println("GORUTİN FUNC 3 ")
	}
}

func main() {
	for i := 0; i <= 1; i++ {
		go gorutin()
		go gorutinF()
		time.Sleep(100 * time.Millisecond)
		go gorutinFF()
		time.Sleep(100 * time.Millisecond)
	}
}
