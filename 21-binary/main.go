package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	a := 1
	b := 7
	c := -7

	fmt.Printf("%b\n", a)
	fmt.Printf("%b %d\n", ^b, ^b)
	fmt.Printf("%b\n", c)

	fmt.Printf(strconv.FormatInt(-7, 2))
	fmt.Printf("\n")

	buf1 := make([]byte, 8)
	buf2 := make([]byte, 8)
	memory := make([]byte, 16)
	binary.PutUvarint(buf1, 12)
	binary.BigEndian.PutUint64(buf2, 12)
	binary.BigEndian.PutUint64(buf2, 12+13<<8)
	fmt.Printf("%d\n", buf2[7])
	fmt.Printf("%d\n", buf2[6])
	copy(memory[2:], buf2)
	copy(memory, buf1)
	binary.LittleEndian.PutUint64(buf2, 12)
}
