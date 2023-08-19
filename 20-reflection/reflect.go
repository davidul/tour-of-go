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

	Reflect(reflect.TypeOf(sample))
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
	//fmt.Println(v.Addr())

	//foo := Foo{
	//	name:    "David",
	//	address: "Address",
	//}

	//bar := FooBar{age: 2}

}

func Reflect(p reflect.Type) {
	fmt.Println("type of")

	k := p.Kind()
	fmt.Println("Kind", p.Kind())
	fmt.Println("Name", p.Name())
	fmt.Println("Elem", p.Elem())
	if k == reflect.Struct {
		fmt.Println("NumField", p.NumField())
	}
	fmt.Println("PkgPath", p.PkgPath())
	fmt.Println("Size", p.Size())
	fmt.Println("String", p.String())
	fmt.Println("Align", p.Align())
	if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 ||
		k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 ||
		k == reflect.Uintptr || k == reflect.Float32 || k == reflect.Float64 || k == reflect.Complex64 ||
		k == reflect.Complex128 {
		fmt.Println("Bits", p.Bits())
	}
	if k == reflect.Chan {
		fmt.Println("Chandir", p.ChanDir())
	}
	//fmt.Println(p.Implements(of))
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
