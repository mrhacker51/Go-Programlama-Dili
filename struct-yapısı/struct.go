package main

import (
	"fmt"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
	Country   string
	City      string
}

/*
type Humans struct {
	Username string
	Password int
	CountryName string
	CityName string
}
*/

func main() {
	structHuman := Human{FirstName: "Anna", LastName: "Debbie", Age: 21, Country: "Russia", City: "Moskov"}
	fmt.Println(structHuman.FirstName, structHuman.LastName, structHuman.Age, structHuman.Country, structHuman.City)

 // veya adlandırma yapmadanda kullanılabilir ... Class yapısı 
	structHumann := Human{"Anna", "Debbie", 21, "Russia", "Moskov"}
	fmt.Println(structHumann.FirstName, structHumann.LastName, structHumann.Age, structHumann.Country, structHumann.City)
}
