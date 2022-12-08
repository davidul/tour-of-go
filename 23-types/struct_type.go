package main

import "fmt"

type Street struct {
	name        string
	houseNo     int
	apartmentNo int
}
type Address struct {
	street  Street //nested struct
	zipcode string
}

func addressMain() {
	fmt.Printf("=== Address Main ===\n")
	address := Address{
		street: Street{
			name:        "Main",
			houseNo:     25,
			apartmentNo: 34},
		zipcode: "25091",
	}

	fmt.Printf("%v", address)
}
