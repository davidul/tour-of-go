package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	var david = person{
		firstName: "David",
		lastName:  "Ulicny",
		age:       45}
	var davidP = &person{
		firstName: "David",
		lastName:  "Ulicny",
		age:       45}
	p := new(person)
	p.lastName = "last_name"
	p.firstName = "first_name"
	p.age = 20
	fmt.Println(p)
	//Equal
	fmt.Println(david == *davidP)

	f(david)
	f1(davidP) //via pointer it mutates the struct
	fmt.Println(david)
	fmt.Println(davidP)

	//Not equal
	fmt.Println(david == *davidP)

	var bob = person{"Bob", "Cat", 12}
	fmt.Println(bob)

	alice := person{}
	alice.firstName = "Alice"
	alice.lastName = "Wonderland"
	alice.age = 18
	fmt.Println(alice)
}

func f(p person) {
	p.firstName = "Daniel"
}

func f1(p *person) {
	p.firstName = "Daniel"
}
