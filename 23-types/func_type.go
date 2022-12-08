package main

import "fmt"

type FuncType func(a int, b int) int

func funcMain() {
	funcType := FuncType(func(a int, b int) int {
		return a + b
	})

	fmt.Println(funcType.calc(1, 1))

	mult := FuncType(func(a int, b int) int {
		return a * b
	})

	fmt.Println(mult.calc(2, 2))
}
func (f FuncType) calc(a int, b int) int {
	return f(a, b)
}
