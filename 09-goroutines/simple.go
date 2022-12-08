package main

import (
	"fmt"
	"time"
)

func simple() {
	ch := make(chan string)

	go func() {
		message := <-ch
		fmt.Println(message)
		ch <- "pong"
	}()

	ch <- "ping"
	fmt.Println(<-ch)
}

func pubSub() {
	fmt.Println("Pub-Sub")
	c := make(chan int)
	//send only channel
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		ch <- x * x //write 9 to c
	}(c, 3)

	d := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch // read 9 from c
		fmt.Println(n)
		time.Sleep(time.Second)
		d <- struct{}{}
	}(c)

	<-d //wait for value
	fmt.Println("End Pub-Sub")
}
