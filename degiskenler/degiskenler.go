package main

import "fmt"

/*

Degisken Yapilari :



*/

// 1. var x = 5 global oluyor var ....
// 2. var x int = 5  global oluyor var ...
// 3. x := 5
/* 4. var (
	s1 = "PHP"
	s2 = "Python"
	s3 = "Go"
	s4 = "BashScript"
)
// 5. x,y := 10,true
*/

// Örnek 1
var x int = 10

func main() {
	fmt.Println(x)

	fmt.Println("Function 1 start ....")
	test() // Asagidaki fonksiyonu ana main'de cagırdık ...

	fmt.Println("Function 2 start .....")
	test_2() // Asagidaki fonksiyonu ana main'de cagırdık ...
	fmt.Println("[+] Program closed ......")

	fmt.Println("Local Function Start ....")
	test_3() // Asagidaki fonksiyonu ana main'de cagırdık ...

	fmt.Println("Multi Variable Func ...")
	test_4() // Asagidaki fonksiyonu ana main'de cagırdık ...

	fmt.Println("Multi := Local Deger Func Start ...")
	test_5() // Asagidaki fonksiyonu ana main'de cagırdık ...

	fmt.Println("Sabit Deger ... Func ....")
	pii() // Asagidaki fonksiyonu ana main'de cagırdık ...
}

//

// Örnek 2

var y = 5

func test() {
	fmt.Println("Deger : ", y)
}

// Örnek 3

var z string = "HelloGo"

func test_2() {
	fmt.Println("String Deger : ", z)
}

// Örnek 4 // Local

func test_3() {
	x := 10
	fmt.Println("Local Deger : ", x)
}

// Örnek 5 // Multi Variable

var (
	s1 = "PHP"
	s2 = "Python"
	s3 = "Go"
	s4 = "BashScript"
)

func test_4() {
	fmt.Println("Lecture1 : ", s1, "\n", "Lecture2 : ", s2, "\n", "Lecture3 : ", s3, "\n", "Lecture4 : ", s4)
}

// Örnek 6

func test_5() {
	x, y := 10, true
	fmt.Println("Deger : ", x, "Deger2 : ", y)
}

// Sabit Degisken ...

// CONST

const pi float32 = 3.14

func pii() {
	fmt.Println("Sabit Deger : ", pi)
}
