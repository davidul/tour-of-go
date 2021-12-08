package main

import (
	"fmt"
	"time"
)

func main() {

	go mul(2,2)

	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)

	time.Sleep(time.Second)
	fmt.Println("Done!")

	c := make(chan int, 1)
	go func(c chan int) {
		writeToChannel(c, 9)
	}(c)

	fmt.Println("Read ", <- c)
}

func mul(a int, b int) int {
	return a * b
}

func writeToChannel(c chan int, x int)  {
	c <- x
	close(c)
}
