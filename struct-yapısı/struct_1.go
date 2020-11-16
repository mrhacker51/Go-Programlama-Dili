package main

import (
	"fmt"
)

type WebSocket struct {
	Host     string
	Port     int
	Requests string
}

func main() {
	//

	Web := WebSocket{Host: "www.google.com", Port: 80, Requests: "GET"}
	fmt.Println(Web.Host, Web.Port, Web.Requests)

	if Web.Host == "www.google.com" {
		fmt.Println("OK...")
	} else {
		fmt.Println("Error....")
	}

	Webs := WebSocket{"www.google.com", 80, "GET"}
	fmt.Println(Webs.Host, Webs.Port, Webs.Requests)

	if Webs.Host == "www.google.com" {
		fmt.Println("OK...")
	} else {
		fmt.Println("Error....")
	}

}
