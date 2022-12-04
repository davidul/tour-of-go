package main

import "fmt"

type FuncType func(a int, b int) int

type Calculator interface {
	Calc(a int, b int) int
}

func (f FuncType) calc(a int, b int) int {
	return f(a, b)
}

func main() {
	funcType := FuncType(func(a int, b int) int {
		return a + b
	})

	fmt.Println(funcType.calc(1, 1))

	mult := FuncType(func(a int, b int) int {
		return a * b
	})

	fmt.Println(mult.calc(2, 2))
}
