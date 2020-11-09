package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("Deger : %d", i)
		if i == 5 {
			break
		}
	}

	for z := 0; z <= 100; z++ {
		if z%2 == 0 {
			fmt.Println("Cift : ", z)
		} else {
			fmt.Println("Tek  : ", z)
		}
		fmt.Println(z)
	}

	d := 100

	for d <= 200 {
		fmt.Printf("Success....[+] %d", d)
		d++
	}

	for t := 0; t <= 10; t++ {
		if t == 7 {
			fmt.Printf("Found ... \n%d", t)
			break
		} else if t == 8 {
			fmt.Printf("FF ... \n%d", t)
			continue
		}

		fmt.Printf("<3<3<3<3")
	}
}
