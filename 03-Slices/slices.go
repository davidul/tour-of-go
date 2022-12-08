package main

import "fmt"

func main() {

	slice()
	fromArray()
	resize()

	var x []int //slice, length avoided
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

	bytes := make([]byte, 4, 4)
	bytes[0] = 10
	bytes[1] = 4
	bytes[2] = 2
	bytes[3] = 3
	//bytes[4] = 4 //throws an error, index out of range
}

func slice() {
	bytes := make([]byte, 5, 5)
	fmt.Printf("Len %d Cap %d \n", len(bytes), cap(bytes))
	bytes[0] = 1

}

func fromArray() {
	a := []byte{1, 2, 3, 4, 5}
	s := a[1:3]
	for i := range s {
		fmt.Printf("%d\n", s[i])
	}

	a[1] = 4

	for i := range s {
		fmt.Printf("%d\n", s[i])
	}

}

func resize() {
	fmt.Printf("=====================\n")
	fmt.Printf("Resize slice and copy\n")
	a := []byte{1, 2}
	s := a[:]
	fmt.Printf("Len %d, Cap %d \n", len(s), cap(s))
	t := make([]byte, len(s), (cap(s)+1)*2)
	for i := range t {
		t[i] = s[i]
	}
	s = t
	fmt.Printf("Len %d, Cap %d \n", len(t), cap(t))
	fmt.Printf("Len %d, Cap %d \n", len(s), cap(s))
	s = s[0:6]
	s[5] = 1
	fmt.Printf("=====================\n")

	//copy. append functions
}
