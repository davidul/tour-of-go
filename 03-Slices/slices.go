package main

import "fmt"

func main() {

	var x []int //slice
	fmt.Printf("Initial length is zero len(x) = %v, capacity=%v\n", len(x), cap(x))
	x = append(x, 10)
	fmt.Printf("Append 1 element %v\n", x)
	x = append(x, 11, 22, 33)
	fmt.Printf("Append another 3 elements %v\n", x)

	fmt.Printf("Length %v  Capacity %v\n", len(x), cap(x))

	var y = make([]int, 3)
	fmt.Printf("Array via make %v Length=%v, Capacity=%v\n", y, len(y), cap(y))
	y[0] = 0
	y[1] = 1
	y[2] = 2

	var z = y[:2]
	fmt.Printf("slice first two elements %v\n", z)

	var v = make([]int, 10)
	for i := 0; i < 10; i++ {
		v[i] = i
	}
	var w = v[3:5]
	fmt.Printf("slice 3:5 %v\n", w)
	fmt.Println(w, v)

}
