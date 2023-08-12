package main

import "fmt"

func main() {
	//for {
	//	fmt.Println("Endless loop")
	//}

	for a := true; a; {
		fmt.Println("a")
		a = false
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 10 {
		fmt.Println("j=", j)
		j++
	}

	v := []int{1, 2, 3}
	for a, b := range v {
		fmt.Println(a, b)
	}
}
