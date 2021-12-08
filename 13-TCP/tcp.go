package main

import (
	"fmt"
	"net"
	"sync"
)

func main1() {

	var wg sync.WaitGroup

	for i := 1; i < 1025; i++{
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("localhost:%d", j)
			tcp, err := net.Dial("tcp", address)
			fmt.Println("Connecting to ", j)
			if err != nil{
				return
			}
			tcp.Close()
			fmt.Printf("Open %d \n", j)
		}(i)

		wg.Wait()
	}
}
