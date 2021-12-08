package main

import "fmt"

func main() {
	var x[2] int
	var y = [2]int{3, 5}
	var z = [...]int{3, 5}

	fmt.Println(x[0])
	fmt.Println(y[0])
	fmt.Println(z[1])
}
