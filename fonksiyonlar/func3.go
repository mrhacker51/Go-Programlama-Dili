package main

import "fmt"

func main() {

	var user_a, user_b int
	fmt.Printf("Userİnput : ")
	fmt.Scanf("%d", &user_a)

	fmt.Printf("Userİnput : ")
	fmt.Scanf("%d", &user_b)

	result := calc(user_a, user_b)

	if result == 75 {
		fmt.Println(result)
	} else if result == 50 {
		fmt.Println(result)
	} else {
		fmt.Println("[+] Good Number .... ")
	}
}

func calc(a, b int) int {
	return (a * b)
}
