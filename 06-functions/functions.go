package main

import "fmt"

func main() {
	i := mul(2, 2)
	fmt.Println(i)

	fmt.Println(sum(1,2,3))

	x, y := plusMinus(1, 1)
	fmt.Println(x, y)
}

func mul(a int, b int) int  {
	return a*b
}

func sum(a ... int) int  {
	ret := 0
	for i := 0; i < len(a); i++{
		ret += a[i]
	}

	return ret
}

func plusMinus(a int, b int) (int, int){
	return a+b, a-b
}
