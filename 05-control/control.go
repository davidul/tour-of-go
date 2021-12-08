package main

import "fmt"

func main() {

	var  i = 2
	if i % 2 == 0 {
		fmt.Println(i)
	}

	var x = [] int {1,2,3,4,5,6}

	for _,b := range x {
		if b % 2 == 0{
			fmt.Println("even")
		}else{
			fmt.Println("odd")
		}
	}
}
