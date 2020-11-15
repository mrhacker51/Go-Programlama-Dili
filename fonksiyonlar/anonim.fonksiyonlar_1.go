package main

import "fmt"

func main() {

	var (
		anonyfunc = func(x, y int) {
			fmt.Println(x*y, x+y, x/y, x-y)
		}
	)

	anonyfunc(5, 5)

	//

	funcnewanony := func() {
		fmt.Println("Program Loading ....")
	}

	funcnewanony()

	//

	// PHP'de oldugu gibi anonim fonksiyonu return etmis iken ekrana yazdırma ....
	fmt.Println(func(x, y int) int {
		return x + y
	}(5, 5))

	//

	// PHP 'de oldugu gibi Anonim fonksiyonu otomatik calistirma :))

	(func(x, y int) {
		fmt.Println("Wuw .....")
	})(5, 5)
}
