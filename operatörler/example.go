package main

import "fmt"

func main() {
	x := 5
	y := 10

	total := x + y
	fmt.Println("Sonuc : ", total)

	total = total - 5
	fmt.Println("Sonuc : ", total)

	total += total + 5
	fmt.Println("Sonuc : ", total)

	total -= total - 5

	fmt.Println("Sonuc : ", total)

	total *= total * 3
	fmt.Println("Sonuc : ", total)

	total /= total / 2
	fmt.Println("Sonuc : ", total)

	total++
	fmt.Println("Sonuc : ", total)
	total--
	fmt.Println("Sonuc : ", total)

	total = total % 2
	fmt.Println("Sonuc : ", total)
}
