package main

import (
	"fmt"
	"time"
)

func main() {
	var username, password string
	fmt.Printf("Username : ")
	fmt.Scanf("%s", &username)
	fmt.Printf("Password : ")
	fmt.Scanf("%s", &password)
	print := fmt.Println
	if username == "root" && password == "root" {
		x, _ := time.ParseDuration("1h30m")
		print(x)
		print("Success...")
		now := time.Now()
		print(now)
	} else {
		x, _ := time.ParseDuration("10m")
		print(x)
		print("Failed...")
	}
}
