package main

import "fmt"

func main() {
	var y int
	var x int = 10
	var a,b,c int
	var e,f,g int = 4,4,4
	z := 1

	h,i := 10, "Hello"

	fmt.Println("y1 = ", y)
	y = 2
	fmt.Println("y2 = ", y)
	fmt.Println(x)
	fmt.Println(a, b, c)
	fmt.Println(e, f, g)
	fmt.Println(z)
	fmt.Println(h,i)

	var(j int
	k int = 8
	l string)

	fmt.Println(j, k, l)

	k = 6
	fmt.Println(k)
	w := 1
	fmt.Println(w)
	w = 2
	fmt.Println(w)

	const name = 3
	//name = 4 -> Error
}
