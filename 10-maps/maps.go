package main

import "fmt"

func main() {
	//nil map
	//var m map[string]string

	m := make(map[string]string, 2)

	m["a"] = "1"
	m["b"] = "2"
	fmt.Println(m)
	fmt.Println(m["a"])

	v, ok := m["a"]
	fmt.Println(v, ok)
	v, ok = m["c"]
	fmt.Println(v, ok)

	delete(m,"a")
	v, ok = m["a"]
	fmt.Println(v, ok)
}
