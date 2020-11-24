package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("proxy.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// flag kısmı bu sekilde 2 degeri aynı anda ekleme yapılarakta kullanılabilir // os.O_CREATE | os.O_RDWR
	// os.O_CREATE == dosya yoksa oluşturma
	// os.O_RDWR == dosyaya hem okuma hemde yazma islemi yapar
	// os.O_APPEND == dosyanın sonuna ekleme yapar
	// os.O_RDONLY == sadece okuma islemi yapar
	// os.O_WRONLY == sadece yazma islemi yapar
	// os.O_TRUNC == dosya acılırken dosyayı keser
	for i := 0; i <= 10; i++ {
		f.WriteString("PROXİES ADDED ....\n")
	}

}
