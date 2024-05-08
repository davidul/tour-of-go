package main

import (
	"fmt"
	"math"
)

func main() {
	//declare a variable
	//default values
	var y int       //0
	var f32 float32 //0
	var f64 float64 //0
	var bv bool     //false
	var s string    //empty string
	fmt.Println("=== Default values ===")
	fmt.Printf("int: %v, float32:%v, float64:%v, bool:%v string:%q\n", y, f32, f64, bv, s)
	fmt.Println("======================")
	var x int = 10 //declare and initialize
	x++

	var e, f, g int = 4, 4, 4
	z := 1 //declare, initialize and infer type
	fmt.Printf("x:%d e:%d f:%d g:%d z:%d \n", x, e, f, g, z)
	h, i := 10, "Hello" //infer different types
	fmt.Printf("%d %s \n", h, i)
	y = 2
	fmt.Printf("y = %d \n", y)

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

	integers()
}

func integers() {
	var a int8 = 1
	var b int16 = 2
	var c int32 = 3
	var d int64 = 4
	fmt.Printf(" %d %d %d %d \n", a, b, c, d)
	fmt.Printf("max int = %d , min int = %d \n", math.MaxInt, math.MinInt)
	fmt.Printf("max int8 = %d , min int8 = %d \n", math.MaxInt8, math.MinInt8)
	fmt.Printf("max int16 = %d , min int16 = %d \n", math.MaxInt16, math.MinInt16)
	fmt.Printf("max int32 = %d , min int32 = %d \n", math.MaxInt32, math.MinInt32)
	fmt.Printf("max int64 = %d , min int64 = %d \n", math.MaxInt64, math.MinInt64)

	fmt.Printf("max unit8 = %d, min uint8 = %d\n", math.MaxUint8, 0)
	fmt.Printf("max unit16 = %d, min uint16 = %d\n", math.MaxUint16, 0)
	fmt.Printf("max unit32 = %d, min uint32 = %d\n", math.MaxUint32, 0)
	fmt.Printf("max unit64 = %d, min uint64 = %d\n", uint64(math.MaxUint64), 0)
	//fmt.Printf("%d \n", math.MaxUint)
	fmt.Printf("%T", a)

	//var f

}
