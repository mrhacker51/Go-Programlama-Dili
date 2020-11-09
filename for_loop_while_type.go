package main

import (
	"fmt"
	"time"
)

func main() {

	x := 100

	for x <= 500 {
		x++
		fmt.Printf("Log : %d\n", x)
		time.Sleep(300 * time.Millisecond)
		if x == 300 {
			break
		}
	}

}
