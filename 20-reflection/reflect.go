package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 1
	fmt.Printf("Type of = %s \n", reflect.TypeOf(i))
	fmt.Printf("Value of = %v \n", reflect.ValueOf(i))

	of := reflect.ArrayOf(5, reflect.TypeOf(i))
	valueOf := reflect.ValueOf(of)
	fmt.Println(valueOf)

	type Sample struct {
		name string
		age  int
	}

	sample := new(Sample)
	sample2 := Sample{
		name: "asd",
		age:  0,
	}

	fmt.Printf("Typeof sample = %s \n", reflect.TypeOf(sample))
	fmt.Printf("Name = %v \n", reflect.TypeOf(sample).Name())
	fmt.Printf("Elem = %v \n", reflect.TypeOf(sample).Elem())
	fmt.Printf("Kind = %v \n", reflect.TypeOf(sample).Kind())
	fmt.Printf("NumField = %v \n", reflect.TypeOf(*sample).NumField())
	for i := 0; i < reflect.TypeOf(*sample).NumField(); i++ {
		field := reflect.TypeOf(*sample).Field(i)
		fmt.Printf("Field Name %v \n", field.Name)
		fmt.Printf("Field Type %v \n", field.Type)
	}

	value := reflect.ValueOf(sample)
	fmt.Printf("Valueof sampe = %v \n", value)
	fmt.Printf("Valueof Kind = %v \n", value.Kind())
	if value.CanAddr() {
		fmt.Printf("Valueof sampe = %v \n", value.Addr())
	}

	//fmt.Printf("Valueof sampe = %v \n", value.Bytes())
	fmt.Println(reflect.TypeOf(sample2))
	v := reflect.ValueOf(sample2)
	p_v := &v
	fmt.Println(p_v.Kind())
	fmt.Println(v.Addr())

	//foo := Foo{
	//	name:    "David",
	//	address: "Address",
	//}

	//bar := FooBar{age: 2}

}

type Bar interface {
	Name()
}

type Foo struct {
	name    string
	address string
}

type FooBar struct {
	age int
}

func (F Foo) Name() {
	fmt.Println("Foo Name")
}

func (F FooBar) Name() {
	fmt.Println("FooBar Name")
}
