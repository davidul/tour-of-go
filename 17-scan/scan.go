package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var s1 string

	fmt.Println("Reading one line, no spaces")
	scanln, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println(scanln)
	fmt.Println(s)

	fmt.Println("Reading one line, empty spaces ignored")
	scan, err := fmt.Scan(&s1)
	if err != nil {
		panic(err)
	}

	fmt.Println(scan)
	fmt.Println(s1)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
