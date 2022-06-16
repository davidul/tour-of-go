package main

import "fmt"

func main() {

	var x = 1
	fmt.Println(x, &x)

	//map is a reference type, always passed as a pointer
	m := make(map[string]string, 2)
	key := "key"
	p(m, &key)
	println(m[key])
}

func p(m map[string]string, key *string) {
	m[*key] = "value"
}
