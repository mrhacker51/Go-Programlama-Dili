package main

import (
	"fmt"
)

type WebSocket struct {
	Host     string
	Port     int
	Requests string
}

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func NewHuman(firstname, lastname string) *Human {
	humanobject := new(Human)
	humanobject.FirstName = firstname
	humanobject.LastName = lastname
	return humanobject
}

func NewsHuman(firstname, lastname string, age int) *Human {
	ObjectHumanNew := new(Human)
	ObjectHumanNew.FirstName = "DEFAULT VARİABLE"
	ObjectHumanNew.LastName = "DEFAULT VARİABLE"
	ObjectHumanNew.Age = 18

	return ObjectHumanNew
}

func main() {
	//

	Web := new(WebSocket)
	Web.Host = "www.google.com"
	Web.Port = 80
	Web.Requests = "POST"

	fmt.Println(Web.Host, ":", Web.Port, ":", Web.Requests)

	//

	resulthuman := NewHuman("Lena", "Temnikova")
	fmt.Println(resulthuman)

	fmt.Println(resulthuman.FirstName, resulthuman.LastName)

	//

	ResultNewhumans := NewsHuman("a", "a", 25)
	fmt.Println(ResultNewhumans.FirstName)
	fmt.Println(ResultNewhumans.LastName)
	fmt.Println(ResultNewhumans.Age)
}
