package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	const longForm = "Jun 16, 2022 at 14:04"
	//"01/02 03:04:05PM '06 -0700"
	parse, err := time.Parse(time.Layout, "06/16 02:08:00PM '22 +0200")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)

	t, err := time.Parse(time.Stamp, "Jun 16 14:15:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Clock())
}
