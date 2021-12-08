package main

import "fmt"

func main() {

	var x [] int
	fmt.Println(len(x))
	x = append(x, 10)
	fmt.Println(x[0])
	x = append(x, 11,22,33)

	fmt.Println(x, len(x), cap(x))

	var y = make([]int, 3)
	fmt.Println(y, len(y), cap(y))

	var z = y[:2]
	fmt.Println(z)

	var v = make([]int, 10)
	var w = v[3:5]
	fmt.Println(w)
	w[0] = 1;
	fmt.Println(w, v)


}
