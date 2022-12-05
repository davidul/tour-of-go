package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	david, davidP, p := initialization()
	// p is zeroed, we need to initialize manually
	p.lastName = "last_name"
	p.firstName = "first_name"
	p.age = 20
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

func initialization() (person, *person, *person) {
	// Initialize as composite literal
	var david = person{
		firstName: "David",
		lastName:  "Ulicny",
		age:       45}

	// pointer to person
	var davidP = &person{
		firstName: "David",
		lastName:  "Ulicny",
		age:       45}

	// with new function
	// allocates memory a zeroes every field
	p := new(person)

	fmt.Printf("First Name = %s, Last Name = %s, Age = %d\n", p.firstName, p.lastName, p.age)

	return david, davidP, p
}

func f(p person) {
	p.firstName = "Daniel"
}

func f1(p *person) {
	p.firstName = "Daniel"
}
