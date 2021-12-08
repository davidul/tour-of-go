package main

import "fmt"

func main1() {
	var i int = 23
	var x int
	x = 43
	y := 3

	fmt.Println(i)
	fmt.Println(x)
	fmt.Println(y)

	//var b bool = true
	const name = 34
	const name1 = "ddd"

	//i2 := add(i, x)
	//s, s2 := swap("x", "y")
	fmt.Println(factorial(5))

	defer fmt.Println("cruel world")
	fmt.Println("goodbye")
	m := make(map[string]string)
	m["a"] = "b"
	fmt.Println(m["a"])
}

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string){
	return b,a
}

func factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * factorial(n - 1)
}
