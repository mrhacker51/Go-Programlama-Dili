package main

import "fmt"

func main() {
	var a string
	fmt.Printf("Username : ")
	fmt.Scanf("%v", &a)

	switch a {
	case "ahmet":
		fmt.Printf("Merhaba %v", a)
		break
	case "mehmet":
		fmt.Printf("Merhaba %v", a)
		break
	default:
		fmt.Printf("Giris basarisiz .... %v", a)
		break
	}
}
