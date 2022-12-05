package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	b := 7
	c := -7

	fmt.Printf("%b\n", a)
	fmt.Printf("%b %d\n", ^b, ^b)
	fmt.Printf("%b\n", c)

	fmt.Printf(strconv.FormatInt(-7, 2))

}
