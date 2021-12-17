package main

import "fmt"

func main() {

	var i = 2
	if i%2 == 0 {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	var x = []int{1, 2, 3, 4, 5, 6}

	for pos, b := range x {
		if b%2 == 0 {
			fmt.Println("even", pos)
		} else {
			fmt.Println("odd", pos)
		}
	}
}
