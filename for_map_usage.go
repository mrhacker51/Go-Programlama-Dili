package main

import "fmt"

func main() {

	baskentler := map[string]string{"Turkey": "Ankara", "France": "Paris", "Japan": "Tokyo"}

	for key, values := range baskentler {
		fmt.Println("Key : ", key, "Values : ", values)
	}

	lecture := map[string]string{"Day1": "PHP", "Day2": "Golang", "Day3": "Python", "Day4": "C"}

	for day, language := range lecture {
		fmt.Println("Programming Day : ", day, "Language :", language)
	}

	student_list := map[string]string{
		"Class1": "Ahmet",
		"Class2": "Mehmet",
		"Class3": "Buse",
		"Class4": "Hande",
		"Class5": "Mehtap",
	}

	for class, student := range student_list {
		fmt.Println("Class Name : ", class, "Student Name : ", student)
	}
}
