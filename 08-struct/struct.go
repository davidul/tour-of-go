package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName string
		age int
	}

	var david = person{firstName: "David", lastName: "Ulicny", age: 45}
	fmt.Println(david)
	
	var bob = person{"Bob", "Cat", 12}
	fmt.Println(bob)

	alice := person{}
	alice.firstName = "Alice"
	alice.lastName = "Wonderland"
	alice.age = 18
	fmt.Println(alice)
}
