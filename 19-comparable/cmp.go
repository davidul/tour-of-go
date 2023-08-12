package main

import (
	"fmt"
	"sort"
)

type A struct {
	x int
	s string
}

type B struct {
	a A
	y int
}

type D struct {
	m map[string]string
}

type E struct {
	m *map[string]string
}

type C interface {
}

func main() {
	a1 := A{x: 1, s: "s"}
	a2 := A{x: 1, s: "s"}
	fmt.Println(a1 == a2) //true
	b1 := B{a: a1, y: 1}
	b2 := B{a: a1, y: 1}
	fmt.Println(b1 == b2) //true

	pa1 := &A{x: 8, s: "x"}
	pa2 := &A{x: 8, s: "x"}
	fmt.Println(pa1 == pa2)   //false
	fmt.Println(*pa1 == *pa2) //true
	pa1 = pa2
	fmt.Println(pa1 == pa2) //true

	//d1 := D{m: make(map[string]string)}
	//d2 := D{m: make(map[string]string)}
	//fmt.Println(d1 == d2) //compile error
	m := make(map[string]string)
	m1 := make(map[string]string)
	e1 := E{m: &m}
	e2 := E{m: &m1}
	fmt.Println(e1 == e2) //false
}

func gtlt() {
	a1 := A{x: 1, s: "s"}
	a2 := A{x: 1, s: "s"}
	if "asdaasdada" > "asdasas" {

	}
	sort.Sort()

}
