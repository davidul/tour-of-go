Variable in Go is defined with "var" keyword, which can be 
in certain situations omitted.
We can declare variable "y" of type int without assigning a value to it.
```go
var y int
```
By default, it gets value 0.

We can declare and assign at the same time
```go
var x int = 10
z := 1
```
Variable "x" is type int and is assigned value 10.
Variable "z" is also an int and is assigned a value 1. Type of "z" inferred by the compiler.

Variables in Go are mutable, unless specified explicitly the opposite.
Imutable variable is prefixed with const
```go
const name = 3
//name = 4 compile time error
```