package main

import "fmt"

type Person[K interface{}] struct {
	k    K
	name string
}

type Address struct {
	street string
	city   string
}

func main() {
	f[int32](18)
	addressablePerson := Person[Address]{
		name: "David",
		k: Address{
			street: "Street",
			city:   "City",
		}}

	p(addressablePerson)
}

func f[T int32 | int64](t T) {

}

func p(person Person[Address]) {
	address := person.k
	fmt.Println(address.street)

}
