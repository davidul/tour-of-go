# Go samples
[01-Variables](#01-Variables)  
[02-Arrays](#02-Arrays)  
[03-Slices](#03-Slices)  
[04-Loops](#04-Loops)

## 01-Variables
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

## 02-Arrays
Arrays have fixed length.
Declare array of ints of size 2. Arrays are indexed from 0.
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

## 02-Arrays
Arrays have fixed length. Array has single type.
Length is part of the type.
Declare array of ints of size 2. Arrays are indexed from 0.
```go
var x [2]int
```
Assign a value to first element of the array
```go
x[0] = 1
```

## 03-Slices
Slices are like arrays but they are dynamic. There is an underlying array behind a slice.
Slices are indexed from 0.
Declare a slice of ints, note slice does not have a length
```go
var x []int
x = make([]int, 3) // make a slice of ints of length 3
x[0] = 1
```

### From array

```go
a := []byte{1, 2, 3, 4, 5}
s := a[1:3]
```


## 04-Loops
In a loop you are executing a block of code multiple times.
Go has only one loop construct:
For loop

The for loop has three components:
- init statement: executed before the first iteration
- condition expression: evaluated before every iteration
- post statement: executed at the end of every iteration
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

### Endless loop
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

## While loop
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

