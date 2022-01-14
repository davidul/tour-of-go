package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var b1 byte = 42
	var b2 int8 = 55
	var b3 uint8 = 65

	fmt.Println(b1)
	fmt.Printf("b1 = %c\n", b1)
	fmt.Println(b2)
	fmt.Printf("b2 = %c\n", b2)
	fmt.Println(b3)
	fmt.Printf("b3 = %c\n", b3)

	var big1 byte = 200
	var big2 byte = 201
	fmt.Println(big1 + big2)

	//string to bytes
	fmt.Print("David -> ")
	fmt.Println([]byte("David"))
	fmt.Print("kancelář -> ")
	fmt.Println([]byte("kancelář"))
	fmt.Print("rune     -> ")
	fmt.Println([]rune("kancelář"))
	fmt.Print("ř -> ")
	fmt.Println([]byte("ř"))

	//bytes to string
	fmt.Println(string([]byte{42, 43, 55, 89}))

	var hx = 0x30
	var oct = 020
	fmt.Println(hx)
	fmt.Println(oct)

	file, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		panic(err)
	}

	fmt.Println(file)
	fmt.Println(string(file))

}
