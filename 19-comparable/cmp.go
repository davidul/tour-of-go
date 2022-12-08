package main

import "fmt"

type A struct {
	x int
	s string
}

type B struct {
	a A
	y int
}

type C interface {
}

func main() {
	a1 := A{x: 1, s: "s"}
	a2 := A{x: 1, s: "s"}
	fmt.Println(a1 == a2)
	b1 := B{a: a1, y: 1}
	b2 := B{a: a1, y: 1}
	fmt.Println(b1 == b2)
}
