package main

import "fmt"

func main() {
	//nil map
	//var m map[string]string

	m := make(map[string]string, 2)

	m["a"] = "1"
	m["b"] = "2"
	fmt.Printf("Map ref: %p \n", m)
	fmt.Printf("Map value: %s \n", m["a"])

	v, ok := m["a"]
	fmt.Printf("Value exists in a map: %s %t \n", v, ok)
	v, ok = m["c"]
	fmt.Printf("Value does not exist in a map: %s %t \n", v, ok)

	delete(m, "a")
	v, ok = m["a"]
	fmt.Printf("Deleted value: %s %t \n", v, ok)

	duplicate()
}

func duplicate() {
	m := make(map[uint16]string, 1)
	m[1] = "1"
	m[2] = "2"
	//replace value
	m[1] = "1_1"

	//iterate over the map
	for u := range m {
		fmt.Printf("Key: %d Value: %s \n", u, m[u])
	}
}
