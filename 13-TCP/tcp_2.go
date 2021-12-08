package main

import (
	"fmt"
	"net"
	"sort"
)

func main2() {

	ports := make(chan int, 100)
	result := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}


	go func() {
		for i := 0; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= 1024; i++ {
		port := <-result
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(result)
	sort.Ints(openports)

	for _, openport := range openports {
		fmt.Println(openport)
	}
}

func worker(ports, result chan int){
	for p := range ports{
		fmt.Println("Scanning ", p)
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		dial, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		dial.Close()
		result <- p
	}
}