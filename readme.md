# Go samples
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