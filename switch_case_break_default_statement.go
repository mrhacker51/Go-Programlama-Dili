package main

import "fmt"

func main() {
	var a int

	fmt.Printf("İnput : ")
	fmt.Scanf("%v", &a)

	switch {
	case a == 10:
		fmt.Printf("Nice ... Girilen Deger : %v", a)
		break
	case a > 10:
		fmt.Printf("TheBest ... Girilen Deger : %v", a)
		break
	case a < 10:
		fmt.Printf("Vuu ... Girilen Deger : %v", a)
		break
	default:
		fmt.Printf("Error ... Girilen Deger : %v", a)
	}

	switch num := 6; num%2 == 0 {
	case true:
		fmt.Println("even value")

	case false:
		fmt.Println("odd value")
	default:
		fmt.Println("Error .....")
		break
	}

	switch numbers := 100; numbers%2 == 1 {
	case true:
		fmt.Printf("Yes")
		break
	case false:
		fmt.Printf("No")
		break
	default:
		fmt.Printf("Errors ....")
		break
	}
}
