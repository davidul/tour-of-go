package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type MyType struct {
		FirstName string `json:"first_name"`
		Surname string `json:"surname"`
		Age     int    `json:"age"`
	}

	myType := MyType{
		FirstName: "David",
		Surname:   "Ulicny",
		Age:       45,
	}

	marshal, err := json.Marshal(&myType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s",marshal)
	}
}
