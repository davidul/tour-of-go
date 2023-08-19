package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"math"
)

func main() {
	figure.NewFigure("Arrays", "", true).Print()
	var x [2]int           //declare array of two elements
	var y = [2]int{3, 5}   //declare and initialize
	var z = [...]int{3, 5} //length is inferred
	var w = [...]string{"Hello", "World"}
	x[0] = 1
	var max [math.MaxInt32]int

	fmt.Printf("x[0] = %v x[1] = %v\n", x[0], x[1])
	//compile time error
	//fmt.Printf("%v", x[2])

	fmt.Printf("y[0] = %v y[1] = %v\n", y[0], y[1])

	fmt.Printf("z[0] = %v z[1] = %v\n", z[0], z[1])

	fmt.Printf("%v", w)

	x[0] = 100
	max[12345] = 100

	fmt.Printf("%d\n", max[0])

	show(x[:])
}

func show(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("\n")
}
