
[01-Variables](#01-Variables)  
[02-Arrays](#02-Arrays)  
[03-Slices](#03-Slices)  
[04-Loops](#04-Loops) 
[05-Functions](#05-Functions)


- Golang is statically typed language. 
- It is compiled language. 
- It is a strongly typed language.
- It is a garbage collected language. 
- It is a concurrent language.

Go syntax is similar to C.
Semicolons are optoinal.
Functions are first class citizens.
Go has pointers but no pointer arithmetic.

# 01-Variables
In Go variables are mutable.
Declare a variable
```go
var y int
```
Default value is 0, for all numeric types.
You can declare and initialize in one line

```go
var x int = 10
```
You can omit the type, Go will infer it
```go
var x = 10
```
You can also omit the `var` keyword and use `:=`
```go
x := 10
```

Default values:
```go
var y int       //0
	var f32 float32 //0
	var f64 float64 //0
	var bv bool     //false
	var s string    //empty string
	fmt.Println("=== Default values ===")
	fmt.Printf("%v %v %v %v %q\n", y, f32, f64, bv, s)
```

# 02-Arrays
- Arrays have fixed length.
- Arrays are indexed from 0.
- Array has single type.
- Length is part of the type.


Declare array of ints of size 2. 
```go
var x [2]int
```

Declare and initialize
```go
var y = [2]int{3, 5}
```

Infer the length of an array
```go
var z = [...]int{3, 5}
```

Assign a value to first element of the array
```go
x[0] = 1
```

Iterate an array:
```go
for i := 0; i < len(x); i++ {
    fmt.Println(x[i])
}
```


# 03-Slices
Slices are like arrays but they are dynamic. There is an underlying array behind a slice.
Slices are indexed from 0.
Declare a slice of ints, note slice does not have a length
```go
var x []int
x = make([]int, 3) // make a slice of ints of length 3
x[0] = 1
```

## From array

```go
a := []byte{1, 2, 3, 4, 5}
s := a[1:3]
```

If you updated the underlying array, the slice will be updated too.
```go 
a := []byte{1, 2, 3, 4, 5}
s := a[1:3]
a[1] = 100
fmt.Println(s) // [100 3]
```


# 04-Loops
In a loop you are executing a block of code multiple times.
Go has only one loop construct:
For loop

The for loop has three components:
- init statement: executed before the first iteration
- condition expression: evaluated before every iteration
- post statement: executed at the end of every iteration

Other statements:
- body: executed at the end of every iteration
- break: terminates the loop
- continue: skips the rest of the body and starts the next iteration
- label: used to break or continue to a specific loop
- goto: jumps to a label
- range: used to iterate over an array, slice, string, map or channel
- for: used to iterate over an array, slice, string, map or channel

```go
for init_statement; condition_expression; post_statement {
    body
}
```

All components are optional and you can omit them all and create
endless loop
```go
for {
    fmt.Println("Endless loop")
}
``` 

You can omit the init statement and use a variable declared outside the loop
```go
i := 0
for ; i < 5; i++ {
    fmt.Println(i)
}
```

You can use the init statement and condition expression only
```go
for a := true; a; {
    fmt.Println("a")
    a = false
}
```

# While loop
While loop style
```go
j := 0
for j < 10 {
    fmt.Println("j=", j)
    j++
}
```


```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

# 05-Functions

Functions are the building blocks of a program. They are used to break a program into smaller pieces.
Functions are declared with the keyword `func` followed by the function name, a list of parameters and the return type.
This is so-called function signature.

Example:
```go
func add(x int, y int) int {
    return x + y
}
```
If the function defines the result parameters it has to return them. This can be achieved with the `return` keyword.
Or terminating statement.

Function can be pass as a parameter to another function.
```go
func add(x int, y int) int {
    return x + y
}
```

```go
func add2(adder func(a interface{}, b interface{}) int, 
	a int, b int) int {
	return adder(a, b)
}

add2(add, 1, 2)
```


Return type can be omitted if all return parameters are named.
```go
func add(x int, y int) (result int) {
    result = x + y
    return
}
```

Multiple return values
```go
func add(x int, y int) (int, int) {
    return x + y, x - y
}
```



