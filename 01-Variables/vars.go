package main

import "fmt"

func main() {
	//declare a variable
	//default values
	var y int       //0
	var f32 float32 //0
	var f64 float64 //0
	var bv bool     //false
	var s string    //empty string
	fmt.Println("=== Default values ===")
	fmt.Printf("%v %v %v %v %q\n", y, f32, f64, bv, s)
	fmt.Println("======================")
	var x int = 10 //declare and initialize
	var e, f, g int = 4, 4, 4
	z := 1 //declare, initialize and infer type

	h, i := 10, "Hello" //infer different types

	y = 2
	fmt.Println("y2 = ", y)
	fmt.Println(x)
	fmt.Println(e, f, g)
	fmt.Println(z)
	fmt.Println(h, i)

	//declaration in block
	var (
		j int
		k int = 8
		l string
	)

	fmt.Println(j, k, l)

	k = 6
	fmt.Println(k)
	w := 1
	fmt.Println(w)
	w = 2
	fmt.Println(w)

	const name = 3
	//name = 4 -> Error
}
