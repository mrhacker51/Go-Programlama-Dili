package main

import "fmt"

// sınıftan sınıfa miras alma islemi inheritance
type Human struct {
	Firstname, Lastname string
}

type Ogrenci struct {
	Human
	Age int
}

//

// sınıfları methotlar ile kullanma

type Requests struct {
	HOST   string
	Port   int
	Method string
}

func NewFunc(host, method string, port int) *Requests {
	rr := new(Requests)
	rr.HOST = host
	rr.Port = port
	rr.Method = method

	return rr
}

//

type NewSocket struct {
	Port int
}

func (sock *NewSocket) socketfunction(port int) {
	if sock.Port == port {
		fmt.Println("Requests Success ....")
	} else {
		fmt.Println("Requests Failed .....")
	}
}

//

func main() {
	result := new(Ogrenci)
	fmt.Println(result)
	result.Firstname = "Hacker"
	result.Lastname = "Man"
	result.Age = 18
	fmt.Println(result.Firstname, "\n", result.Lastname)

	// Requests Class , NewFunc Function

	RequestsClassMethod := NewFunc("www.youtube.com", "HEAD", 80) // Function ...
	print := fmt.Println                                          // Print
	print(RequestsClassMethod.HOST)                               // Class HOST Parametr
	print(RequestsClassMethod.Method)                             // Class Method Parametr
	print(RequestsClassMethod.Port)                               // Class Port Parametr

	//

	NewSocks := NewSocket{80}
	NewSocks.socketfunction(80)

}
