
package main	

import (	
	"log"	
	"os"	
)	

func main() {	
	oldpath := "tester.txt"	
	newpath := "./news/test.txt"	
	err := os.Rename(oldpath, newpath)	

	if err != nil {	
		log.Fatal(err)	
	}	

}
