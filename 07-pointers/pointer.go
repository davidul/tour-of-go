package main

import "fmt"

func main() {

	var x = 1
	fmt.Printf("Value of x %d and pointer to x %p \n", x, &x)

	y := new(int)
	fmt.Printf("y is pointer %p with int value %d \n", y, *y)
	//map is a reference type, always passed as a pointer
	m := make(map[string]string, 2)
	key := "key"
	p(m, &key)
	println(m[key])
}

func p(m map[string]string, key *string) {
	m[*key] = "value"
}
