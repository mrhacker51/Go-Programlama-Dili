package main

import "fmt"

func main() {
	var username, password string

	fmt.Printf("Username : ")
	fmt.Scanf("%v", &username)

	fmt.Printf("Password : ")
	fmt.Scanf("%v", &password)

	switch {
	case username == "root" && password == "root":
		fmt.Printf("[+] Success Login .... Loading ....\nWelcome \n%v\n%v\n", username, password)
		break
	case username == "admin" && password == "admin":
		fmt.Printf("[:(] You Hacker ? ? ?  .... Hımmm ....")
		break
	default:
		fmt.Println("Wrong Username or Password ....")
		break
	}
}
