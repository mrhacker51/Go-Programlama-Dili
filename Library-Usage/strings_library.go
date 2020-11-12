package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// import ( s "strings") şeklindede kısa tanım s.ToUpper vs ...

func main() {

	// color library ....
	color.Red("hello world")
	color.Blue("Hacker")
	color.WhiteString("hacker")
	color.Yellow("hacker")
	color.Green("hacker")
	colors := "HelloWorld"

	// strings library

	//

	fmt.Println(strings.Count(colors, "e"))
	fmt.Println(strings.Index(colors, "e"))
	fmt.Println(strings.Replace(colors, "e", "a", 1))
	fmt.Println(strings.ReplaceAll(colors, "e", "wwww"))
	fmt.Println(colors)
	fmt.Println(strings.Split(colors, "o"))

	for _, values := range colors {
		fmt.Println("Value : ", values) // char type .... return ...
	}
	fmt.Println(strings.ToTitle(colors))
	fmt.Println(strings.ToUpper(colors))
	fmt.Println(strings.Title(colors))
	fmt.Println(strings.ToLower(colors))
	fmt.Println(strings.HasPrefix(colors, "He"))
	fmt.Println(strings.HasSuffix(colors, "ld"))


	// import "s" sekli ...
	// var prints = fmt.Println // print kısa tanım ... 


	fmt.Println(s.Count(colors, "e"))
	fmt.Println(s.Index(colors, "e"))
	fmt.Println(s.Replace(colors, "e", "a", 1))
	fmt.Println(s.ReplaceAll(colors, "e", "wwww"))
	fmt.Println(colors)
	fmt.Println(s.Split(colors, "o"))

	for _, values := range colors {
		fmt.Println("Value : ", values) // char type .... return ...
	}
	fmt.Println(s.ToTitle(colors))
	fmt.Println(s.ToUpper(colors))
	fmt.Println(s.Title(colors))
	fmt.Println(s.ToLower(colors))
	fmt.Println(s.HasPrefix(colors, "He"))
	fmt.Println(s.HasSuffix(colors, "ld"))
}
